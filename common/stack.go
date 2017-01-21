package common

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"io"
	"strings"
	"time"
)

// CreateStackName will create a name for a stack
func CreateStackName(stackType StackType, name string) string {
	return fmt.Sprintf("mu-%s-%s", stackType, name)
}

// StackWaiter for waiting on stack status to be final
type StackWaiter interface {
	AwaitFinalStatus(stackName string) string
}

// StackUpserter for applying changes to a stack
type StackUpserter interface {
	UpsertStack(stackName string, templateBodyReader io.Reader, parameters map[string]string, tags map[string]string) error
}

// StackLister for listing stacks
type StackLister interface {
	ListStacks(stackType StackType) ([]*Stack, error)
}

// StackGetter for getting stacks
type StackGetter interface {
	GetStack(stackName string) (*Stack, error)
}

// StackDeleter for deleting stacks
type StackDeleter interface {
	DeleteStack(stackName string) error
}

// ImageFinder for finding latest image
type ImageFinder interface {
	FindLatestImageID(namePattern string) (string, error)
}

// StackManager composite of all stack capabilities
type StackManager interface {
	StackUpserter
	StackWaiter
	StackLister
	StackGetter
	StackDeleter
	ImageFinder
}

type cloudformationStackManager struct {
	cfnAPI cloudformationiface.CloudFormationAPI
	ec2API ec2iface.EC2API
}

// TODO: support "dry-run" and write the template to a file
// fmt.Sprintf("%s/%s.yml", os.TempDir(), name),

// NewStackManager creates a new StackManager backed by cloudformation
func newStackManager(region string) (StackManager, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	log.Debugf("Connecting to CloudFormation service in region:%s", region)
	cfnAPI := cloudformation.New(sess, &aws.Config{Region: aws.String(region)})

	log.Debugf("Connecting to EC2 service in region:%s", region)
	ec2API := ec2.New(sess, &aws.Config{Region: aws.String(region)})

	return &cloudformationStackManager{
		cfnAPI: cfnAPI,
		ec2API: ec2API,
	}, nil

}

func buildStackParameters(parameters map[string]string) []*cloudformation.Parameter {
	stackParameters := make([]*cloudformation.Parameter, 0, len(parameters))
	for key, value := range parameters {
		stackParameters = append(stackParameters,
			&cloudformation.Parameter{
				ParameterKey:   aws.String(key),
				ParameterValue: aws.String(value),
			})
	}
	return stackParameters
}

func buildStackTags(tags map[string]string) []*cloudformation.Tag {
	stackTags := make([]*cloudformation.Tag, 0, len(tags)+2)

	stackTags = append(stackTags,
		&cloudformation.Tag{
			Key:   aws.String("mu:version"),
			Value: aws.String(GetVersion()),
		})

	for key, value := range tags {
		stackTags = append(stackTags,
			&cloudformation.Tag{
				Key:   aws.String(fmt.Sprintf("mu:%s", key)),
				Value: aws.String(value),
			})
	}
	return stackTags
}

// UpsertStack will create/update the cloudformation stack
func (cfnMgr *cloudformationStackManager) UpsertStack(stackName string, templateBodyReader io.Reader, parameters map[string]string, tags map[string]string) error {
	stackStatus := cfnMgr.AwaitFinalStatus(stackName)

	// load the template
	templateBodyBytes := new(bytes.Buffer)
	templateBodyBytes.ReadFrom(templateBodyReader)
	templateBody := aws.String(templateBodyBytes.String())

	// stack parameters
	stackParameters := buildStackParameters(parameters)

	// stack tags
	stackTags := buildStackTags(tags)

	cfnAPI := cfnMgr.cfnAPI
	if stackStatus == "" {

		log.Debugf("  Creating stack named '%s'", stackName)
		log.Debugf("  Stack parameters:\n\t%s", stackParameters)
		log.Debugf("  Stack tags:\n\t%s", stackTags)
		params := &cloudformation.CreateStackInput{
			StackName: aws.String(stackName),
			Capabilities: []*string{
				aws.String(cloudformation.CapabilityCapabilityIam),
			},
			Parameters:   stackParameters,
			TemplateBody: templateBody,
			Tags:         stackTags,
		}
		_, err := cfnAPI.CreateStack(params)
		log.Debug("  Create stack complete err=%s", err)
		if err != nil {
			return err
		}

		waitParams := &cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		}
		log.Debug("  Waiting for stack to exist...")
		cfnAPI.WaitUntilStackExists(waitParams)
		log.Debug("  Stack exists.")

	} else {
		log.Debugf("  Updating stack named '%s'", stackName)
		log.Debugf("  Prior state: %s", stackStatus)
		log.Debugf("  Stack parameters:\n\t%s", stackParameters)
		log.Debugf("  Stack tags:\n\t%s", stackTags)
		params := &cloudformation.UpdateStackInput{
			StackName: aws.String(stackName),
			Capabilities: []*string{
				aws.String(cloudformation.CapabilityCapabilityIam),
			},
			Parameters:   stackParameters,
			TemplateBody: templateBody,
			Tags:         stackTags,
		}

		_, err := cfnAPI.UpdateStack(params)
		log.Debug("  Update stack complete err=%s", err)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == "ValidationError" && awsErr.Message() == "No updates are to be performed." {
					log.Infof("  No changes for stack '%s'", stackName)
					return nil
				}
			}
			return err
		}

	}
	return nil
}

// AwaitFinalStatus waits for the stack to arrive in a final status
//  returns: final status, or empty string if stack doesn't exist
func (cfnMgr *cloudformationStackManager) AwaitFinalStatus(stackName string) string {
	cfnAPI := cfnMgr.cfnAPI
	params := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}
	resp, err := cfnAPI.DescribeStacks(params)

	if err == nil && resp != nil && len(resp.Stacks) == 1 {
		switch *resp.Stacks[0].StackStatus {
		case cloudformation.StackStatusReviewInProgress,
			cloudformation.StackStatusCreateInProgress,
			cloudformation.StackStatusRollbackInProgress:
			// wait for create
			log.Debugf("  Waiting for stack:%s to complete...current status=%s", stackName, *resp.Stacks[0].StackStatus)
			cfnAPI.WaitUntilStackCreateComplete(params)
			resp, err = cfnAPI.DescribeStacks(params)
		case cloudformation.StackStatusDeleteInProgress:
			// wait for delete
			log.Debugf("  Waiting for stack:%s to delete...current status=%s", stackName, *resp.Stacks[0].StackStatus)
			cfnAPI.WaitUntilStackDeleteComplete(params)
			resp, err = cfnAPI.DescribeStacks(params)
		case cloudformation.StackStatusUpdateInProgress,
			cloudformation.StackStatusUpdateRollbackInProgress,
			cloudformation.StackStatusUpdateCompleteCleanupInProgress,
			cloudformation.StackStatusUpdateRollbackCompleteCleanupInProgress:
			// wait for update
			log.Debugf("  Waiting for stack:%s to update...current status=%s", stackName, *resp.Stacks[0].StackStatus)
			cfnAPI.WaitUntilStackUpdateComplete(params)
			resp, err = cfnAPI.DescribeStacks(params)
		case cloudformation.StackStatusCreateFailed,
			cloudformation.StackStatusCreateComplete,
			cloudformation.StackStatusRollbackFailed,
			cloudformation.StackStatusRollbackComplete,
			cloudformation.StackStatusDeleteFailed,
			cloudformation.StackStatusDeleteComplete,
			cloudformation.StackStatusUpdateComplete,
			cloudformation.StackStatusUpdateRollbackFailed,
			cloudformation.StackStatusUpdateRollbackComplete:
			// no op

		}

		if len(resp.Stacks) > 0 {
			log.Debugf("  Returning final status for stack:%s ... status=%s", stackName, *resp.Stacks[0].StackStatus)
			return *resp.Stacks[0].StackStatus
		}
	}

	log.Debugf("  Stack doesn't exist ... stack=%s", stackName)
	return ""
}

func buildStack(stackDetails *cloudformation.Stack) *Stack {
	stack := new(Stack)
	stack.ID = aws.StringValue(stackDetails.StackId)
	stack.Name = aws.StringValue(stackDetails.StackName)
	stack.Status = aws.StringValue(stackDetails.StackStatus)
	stack.StatusReason = aws.StringValue(stackDetails.StackStatusReason)
	stack.LastUpdateTime = aws.TimeValue(stackDetails.LastUpdatedTime)
	stack.Tags = make(map[string]string)
	stack.Outputs = make(map[string]string)

	for _, tag := range stackDetails.Tags {
		key := aws.StringValue(tag.Key)
		if strings.HasPrefix(key, "mu:") {
			stack.Tags[key[3:]] = aws.StringValue(tag.Value)
		}
	}

	for _, output := range stackDetails.Outputs {
		stack.Outputs[aws.StringValue(output.OutputKey)] = aws.StringValue(output.OutputValue)
	}

	return stack
}

// ListStacks will find mu stacks
func (cfnMgr *cloudformationStackManager) ListStacks(stackType StackType) ([]*Stack, error) {
	cfnAPI := cfnMgr.cfnAPI

	params := &cloudformation.DescribeStacksInput{}

	var stacks []*Stack

	log.Debugf("Searching for stacks of type '%s'", stackType)

	err := cfnAPI.DescribeStacksPages(params,
		func(page *cloudformation.DescribeStacksOutput, lastPage bool) bool {
			for _, stackDetails := range page.Stacks {
				if cloudformation.StackStatusDeleteComplete == aws.StringValue(stackDetails.StackStatus) {
					continue
				}

				stack := buildStack(stackDetails)

				if stack.Tags["type"] == string(stackType) {
					stacks = append(stacks, stack)
				}
			}

			return true
		})

	if err != nil {
		return nil, err
	}
	return stacks, nil
}

// GetStack get a specific stack
func (cfnMgr *cloudformationStackManager) GetStack(stackName string) (*Stack, error) {
	cfnAPI := cfnMgr.cfnAPI

	params := &cloudformation.DescribeStacksInput{StackName: aws.String(stackName)}

	log.Debugf("Searching for stack named '%s'", stackName)

	resp, err := cfnAPI.DescribeStacks(params)
	if err != nil {
		return nil, err
	}
	stack := buildStack(resp.Stacks[0])

	return stack, nil
}

// FindLatestImageID for a given
func (cfnMgr *cloudformationStackManager) FindLatestImageID(namePattern string) (string, error) {
	ec2Api := cfnMgr.ec2API
	resp, err := ec2Api.DescribeImages(&ec2.DescribeImagesInput{
		Owners: []*string{aws.String("amazon")},
		Filters: []*ec2.Filter{
			{
				Name: aws.String("name"),
				Values: []*string{
					aws.String(namePattern),
				},
			},
		},
	})

	if err != nil {
		return "", err
	}

	var imageID string
	var imageCreateDate time.Time
	for _, image := range resp.Images {
		createDate, err := time.Parse(time.RFC3339, aws.StringValue(image.CreationDate))
		if err != nil {
			return "", err
		}
		if imageCreateDate.Before(createDate) {
			imageCreateDate = createDate
			imageID = aws.StringValue(image.ImageId)
		}
	}

	if imageID == "" {
		return "", errors.New("Unable to find image")
	}
	log.Debugf("Found latest imageId %s for pattern %s", imageID, namePattern)
	return imageID, nil
}

// DeleteStack delete a specific stack
func (cfnMgr *cloudformationStackManager) DeleteStack(stackName string) error {
	cfnAPI := cfnMgr.cfnAPI

	params := &cloudformation.DeleteStackInput{StackName: aws.String(stackName)}

	log.Debugf("Deleting stack named '%s'", stackName)

	_, err := cfnAPI.DeleteStack(params)
	return err
}
