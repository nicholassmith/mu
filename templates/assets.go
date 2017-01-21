// Code generated by go-bindata.
// sources:
// assets/cluster.yml
// assets/repo.yml
// assets/vpc.yml
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsClusterYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1a\xfb\x6f\x13\x39\xfa\xf7\xfe\x15\x1f\x39\x4e\xd9\x5d\x31\x49\x9a\x2e\x2b\x18\x5d\x59\xa5\x69\x80\x88\x14\xa2\xa6\x80\xc4\x43\xc8\x99\xf9\x92\xf8\x98\xb1\xe7\x6c\x4f\xdb\x2c\xd7\xff\xfd\x64\x7b\x1e\x9e\x47\x42\xcb\x72\xb7\xb7\xa7\x0b\xaa\xc8\x7c\xfe\xfc\xbd\x5f\xf6\xc4\xf3\xbc\x83\xd1\xdb\xc5\x05\xc6\x49\x44\x14\x3e\xe5\x22\x26\xea\x0d\x0a\x49\x39\xf3\xa1\x3b\x1c\x1c\x0e\xbc\xc1\x63\x6f\xf0\xb8\x7b\x30\x27\x82\xc4\xa8\x50\x48\xff\x00\x60\xca\xa4\x22\x2c\xc0\x8b\x6d\x82\xfa\x19\xc0\x7c\x83\x85\x12\x94\xad\x0d\xe0\x14\x65\x20\x68\xa2\x0c\xa9\x1c\x1f\xd4\x36\x41\x50\x1c\x52\x89\xbd\x0c\x6d\x45\xd2\x48\xf9\xa0\x86\xbd\x98\x06\x82\x1f\x98\xad\x54\x60\x38\x26\x09\x09\xa8\xda\xba\x0c\x5e\xa6\xf1\x12\x45\x75\x67\xf7\xb0\xdb\xe4\x68\x11\x81\xaf\x80\x66\xbc\xa5\xe6\x1b\x91\x94\x05\x1b\xa0\x0c\xb6\x3c\x15\x30\x19\x2f\x20\x88\x52\xa9\x0c\xcd\x33\x72\xbd\xa0\xbf\xe1\x57\xf9\x0d\x5b\xf8\x9d\x91\x6b\x1a\xa7\x31\xb0\x36\xbe\x1b\xa2\x20\x20\x0c\x96\x98\x09\x80\xe1\x0e\x11\x5e\xe0\xf6\x25\x89\x6f\x65\xd3\x0c\x55\x6b\x45\xa4\xe4\x01\x25\x0a\xe1\x8a\xaa\x0d\x5c\x71\xf1\x19\x45\x29\x40\x0f\x60\x86\xe4\x12\x61\x19\x11\xf6\x59\x6f\x08\xa9\x24\xcb\x08\x61\xb1\x78\x0e\x24\x08\x50\xca\x9a\x37\xba\x5a\xc5\x85\xdc\x8c\xa2\x88\x5f\xf9\x4d\xe6\x8b\x74\xc9\x50\xc1\x4a\xf0\x18\xae\x36\x34\xd8\x18\x31\x34\x72\x83\x66\x43\x8b\x33\xca\x66\xc8\xd6\x6a\xe3\x43\xf7\xb1\x35\xe5\x19\xb9\x2e\x40\x87\x8f\xba\x55\x59\x06\x3d\xf3\xaf\x3f\x30\x60\x23\x11\x86\x73\xa2\x14\x0a\xe6\x43\xe7\x87\x0f\x1f\xc2\x2f\x87\x0f\x8e\x6e\x7e\xfc\xf0\xa1\x77\x9b\x87\x7e\xf6\x75\x78\xf3\x63\xc7\x90\x1c\x73\x26\x95\x20\x94\xa9\x8a\x8e\xdd\x38\x95\x4a\xfb\x8c\xc0\x25\x89\x68\x08\xe3\xe9\xe9\x39\x2c\x23\x1e\x7c\xf6\xe1\xba\x67\xfe\xf5\xaf\x7b\xc6\x52\x01\x89\xf0\x55\xaa\x2e\x36\x02\xe5\x86\x47\x61\x8b\xc9\x8a\x35\x20\x2a\xb3\x19\x01\xa9\x37\x02\x4f\x15\xe0\x25\x32\x05\x57\x34\x8a\x34\x4b\xca\xa8\xd2\x0e\x0d\xbf\x16\x8b\x8f\x06\x05\xff\x29\xfb\x46\xf6\x94\x7d\x2b\xf7\x23\xc3\x7d\x1a\x93\x35\x4e\xdb\x98\xea\xe0\x1e\x9d\x4d\xcb\xcc\xdb\x15\xd7\x6e\xd8\xbd\x49\x82\x9c\xd8\xfe\x0c\x30\xe1\xcf\x57\xa0\x36\xa8\x3d\x94\x9a\x5c\xa0\x71\xc2\x85\x82\x15\x17\x06\x6e\x88\x1d\x00\xcc\xd3\x65\x44\x03\x1b\xb6\xa3\x77\x87\xdf\x8f\xc1\xe8\xdd\x21\x48\x9b\x0d\xb4\xc9\x68\xf8\x3d\x19\x0d\x2b\x8c\xea\x66\xab\x32\x3e\xfa\x9e\x8c\x8f\xf6\x30\x3e\x43\x45\x42\xa2\x88\xe6\x36\x7a\xbb\xf0\xfd\x71\xc4\xd3\xd0\xb6\x12\xcd\xc2\x9f\x32\x85\x62\x45\x82\xac\xa8\x15\x8d\xe4\x99\xe0\x69\x22\x2d\x10\xc0\x83\x19\x59\x62\x94\x3f\xea\x4f\x98\x73\xe9\x14\xed\x63\xcc\xd9\x8a\xae\x53\x61\x48\x77\x0a\xdc\x6a\x73\xca\x3f\x5e\xa5\x4d\x55\x16\xb2\xda\x59\x81\xe5\xd5\xee\x36\x02\x8d\x52\xc5\x4d\xd2\x51\xb6\xbe\xab\x50\xb5\xee\x56\x59\xcb\x3a\x50\xd5\x50\x46\x8e\x82\x48\xb3\xf5\xee\xb0\x55\xde\x6a\x6d\xe2\xfd\x9a\x0b\x56\xe9\x30\xd5\xad\x2f\x70\xab\x37\xac\x05\x61\xca\x29\xe3\xf0\x83\xed\x1b\x3a\x1e\x18\x67\xf8\x63\x41\xab\xda\x20\xaa\xc4\xca\x62\xd9\x46\xb3\x20\xd1\xda\xeb\xab\x94\x32\x14\xb7\x53\x16\xbd\x0d\x02\x9e\x32\x55\x50\xab\x74\xf0\x2a\x95\xbc\x41\xef\xa5\x32\xe6\x2c\xa4\xda\x8d\xc6\xdc\xcf\x89\xac\x58\xab\xf3\x94\xf9\xfe\x4b\xae\x3a\x65\xd0\x1a\xd0\xe4\x1f\x29\x89\x64\xc7\x87\xf7\xf7\xce\x71\x95\x5b\xf8\x01\x74\xbb\x1f\x2d\x95\x5a\xf1\xb9\x13\xb5\x46\xe1\xda\x49\x77\xf8\x3b\xe8\x0e\xf7\xd0\x3d\xfa\x1d\x74\x8f\x72\xba\xe7\x28\x79\x2a\x02\x34\x86\x9d\x04\x72\x6c\x5d\xe0\xd6\x28\x53\x3d\x26\x63\x53\x42\xf2\x79\x68\x32\x5e\xe8\x5c\xcb\x52\xcd\x94\x8c\xc6\x16\x07\xa1\xf2\x60\xb0\xb3\x7a\x95\x20\x0b\xe5\x2b\xe6\xc3\xfb\x8f\x36\xb9\x04\x4f\x50\x28\x8a\x45\x5e\xbd\x99\x8f\xdf\x71\x86\xd3\x10\x99\xa2\x2b\x9a\x8b\xa6\xd5\xd4\x5a\x4e\x57\x65\x50\x79\x2d\x3e\x75\x16\x0d\xba\x29\xa1\x6f\x74\x45\xf5\xe1\xde\x22\x5d\xc2\xfd\x2f\x0d\x4f\xde\x38\x9b\x8c\xed\x8c\x3a\x2f\xb9\xd9\x76\x17\xee\xc3\x3b\x73\x1f\x7e\x47\xee\x47\x77\xe6\x7e\x74\x3b\xee\x33\x53\xb9\x2a\xe5\xd5\x24\xa3\xdd\x30\xe6\x4c\x11\xca\x50\xe4\x15\x4f\xe6\x45\x80\x32\x53\x04\x8a\x83\x41\x59\x17\xec\x4e\xb7\xca\x36\x2b\x90\xc5\x69\xaf\xd2\x17\x64\xed\x74\xac\x17\xb8\xb5\x1d\xb4\xd0\x25\x57\xba\x50\x68\xa1\x48\xf0\xb9\x82\xa2\x23\x8f\xac\x89\xc2\x91\xb2\xfa\xf9\xa0\x44\xa6\xf2\x58\xa0\xd1\x72\xce\x23\x1a\x14\xb5\x30\x4f\x9d\x05\x5d\x33\xe2\x34\xa4\x0b\x1a\x23\x4f\x95\x0f\xf3\x8b\xc3\x87\x67\x06\xfc\x3a\x09\x89\xc2\xea\x76\x27\x23\xce\x79\xa4\xff\xb3\x58\x25\xa1\x33\xca\x0a\x1b\x4e\xd9\x02\xc5\x25\x0d\x2a\xe6\x33\x06\x3c\x21\x2a\xd8\xd4\x0d\xab\xdb\x54\x2a\x51\x8b\xe2\xca\xa1\x3f\x6f\x09\x55\xaf\x58\x55\x78\xe9\x43\x57\x6b\xeb\xce\xcc\xae\xb4\xbb\xb2\x3a\xfb\x62\x51\x77\x24\xf0\x28\xfc\x7b\x2a\x55\x8c\x4c\x59\x2a\xe3\x0d\x61\x6b\x9c\xb2\x9a\x0b\xeb\x05\xc2\x89\xa8\x96\x62\x93\x6d\x1a\x73\x1e\x85\xfc\x8a\xf9\x70\x34\x18\xe4\xcd\xcf\xa2\x95\x6c\x7d\x38\x2c\x47\xf1\xff\x21\xad\x3c\xad\xd6\x19\xc6\x5c\x6c\x47\x11\x11\xf1\x73\xba\xde\x34\x14\x33\x53\xdf\x5b\x1d\x22\xbe\x6f\xb0\x76\xe9\xa3\xd7\x2a\x23\xa8\xa9\x10\xc6\x6a\x9e\x3e\x07\xd1\x55\xc6\x0b\x9e\xc0\xfd\x2f\x8d\x83\xd5\xcd\x5f\xcd\x1c\xf2\x10\x62\xca\x52\x55\xe6\x3c\x2a\x41\x03\xab\xb5\xdd\x7e\x8e\x12\xc5\xa5\x49\xa7\x0c\x47\xaf\xca\x44\x0f\xa2\x5a\xe4\xfe\x64\xbc\xc8\x75\x56\x44\x51\xa9\x68\xe0\xc3\xe8\x12\x05\x59\xe7\xd9\x3a\x47\x41\x79\xe8\x9a\x67\xa2\x87\x64\x9b\xa3\x66\x4d\x5a\xa7\x1b\x63\x14\x87\x2f\x6b\xf7\x86\xe8\xae\x05\x46\x41\x31\x65\x38\x35\xb0\x9a\x12\x79\x7d\xa2\x31\x32\x59\xc5\xb6\x8a\x66\x6d\x72\x67\x05\x2a\xdb\x6c\xe1\xf0\x38\x21\x82\x4a\xce\x5e\x25\x28\x88\xe2\xc2\x87\x67\xba\xe6\xa0\xb8\xd8\x10\xe6\x4a\xea\xf8\x7b\x96\x0f\x78\xdf\xdb\xdd\x94\x39\xde\xfe\x5b\xee\x6d\xe7\x18\xfb\x27\x73\xf6\x94\xdd\xd5\xd7\x79\x9d\xf8\xcf\xb8\x7a\x86\x52\xd6\xfd\xdc\xec\xa2\xfb\x6b\x56\x4b\x4f\x3e\xc8\x7c\x52\x1c\x03\x8d\xe6\xed\x47\x41\xaa\xca\xb6\x13\x18\x22\xee\xe1\x28\xe0\x71\x4c\x58\x58\x39\x30\x01\x0c\x0e\x3f\x91\x30\xfc\x94\x0f\xeb\x9f\x14\xff\x14\xb8\xd3\x63\x63\x7f\x16\x64\xff\xac\xad\x02\xfc\xe5\x5e\x7f\x49\x59\x7f\x49\xe4\xa6\xb1\x86\xc1\x86\xeb\x5a\xf9\x69\x3c\x7b\xbd\xb8\x98\x9c\x1f\xdf\xff\x52\x1a\xf5\x06\xe0\xc9\x13\xe8\xa3\x0a\xfa\x18\x48\xfd\xd7\xb3\xd2\x3b\x64\x56\x34\xc2\x9a\xe4\x1d\xb3\x23\x58\x31\xfd\xe7\x6d\xd2\xc4\xec\xea\x34\xc5\x66\xca\x54\xda\x1d\x62\xbf\x8f\x09\x65\x1f\x1b\x60\xa9\x67\x8b\xe3\xfb\x5f\xca\x41\xc3\x1d\xab\xf2\x8f\xc0\x35\xe5\x2c\x47\x3b\x37\x4f\x75\xac\x98\x87\xba\xab\x0f\x06\x83\x9f\x07\x83\x6e\x6d\x91\x5f\x31\x14\x3e\x08\xce\x55\x6d\x65\x6d\xc6\xf1\xe6\x4a\xa9\xf6\x86\xf3\xcf\xb2\x17\x1a\xf5\x49\xaa\xb8\x27\x30\xe2\x24\x44\xf1\x8d\x86\x68\xd0\xf1\x34\x87\xa6\x69\x94\xa0\xeb\x35\x0a\x79\x9c\x70\xa9\x7a\xa9\x99\x77\x1a\x48\x09\x51\x9b\xe3\xe2\x5c\xd2\x6b\x66\x42\x2f\x0f\xea\xde\xce\x68\x6e\x10\x25\x26\xd9\x8f\xfb\x3c\x51\x7d\x72\x25\x4d\xbc\x69\xa9\x29\xa3\x0a\xbc\x4b\xf0\x3c\xe3\x36\x70\xdd\xa6\xb3\xfa\x06\x3c\x4f\x64\xb2\xb4\x24\xa5\x59\xd5\xae\x83\xbd\x8e\x04\x10\x29\x23\xf2\xb8\xe6\x12\x69\x47\xba\x5a\x74\xca\xad\xbc\xa4\x95\x8c\xcc\xbc\x60\x63\xb5\x0e\x06\x40\x46\x96\x11\x86\xce\x0c\x57\x5f\x97\xa9\xc0\xf3\x94\x31\x5d\x2a\x76\x61\xb5\xe4\x09\xd8\xd3\x64\x7b\xb6\xec\xc5\xfc\x4a\x80\xed\xe8\x4a\xf9\x85\xa5\xad\x9d\xd9\x53\xde\x1f\x30\x48\x05\x55\xdb\xec\x7a\x0a\xde\x5b\xa4\xe7\x5c\xaa\xc5\x33\xc8\x43\xad\x72\x0b\x93\x91\x69\x5e\x36\x4d\x49\x9c\x43\xe7\x82\x6b\xc5\x8b\x99\x6c\x58\x5b\xc8\x76\xe4\xf7\x0d\x70\x6f\xba\x82\xf7\xce\x0d\xc4\x03\xa8\xde\x2d\x98\xa7\x8e\x7b\x68\xea\xe4\xb2\xbd\x96\x28\x4e\x9d\x52\x0c\xe6\x54\x76\x42\x24\xfe\xf2\xb3\x6b\xf7\x96\x24\x73\x0a\x24\x78\xd7\xd5\x94\xd9\xa6\xb1\xbd\x30\x89\x22\xf0\xb6\x40\xae\xa4\xa7\xad\xbe\xe4\x5c\x49\x25\x48\x52\x41\xfe\x43\xe2\xbf\xc1\x54\x9a\x43\x07\x78\x08\xf7\x7f\xbd\x1d\xe7\x96\x61\x79\x0f\xeb\xa6\x1b\x1b\xcd\x73\x3a\x3a\xd3\x95\xa2\xe9\xeb\x66\x54\xce\x89\xda\xf8\xd0\xe9\xe7\x11\x7f\xce\x9d\x44\xf1\x8a\xc0\xd1\x60\xcb\x5b\x7f\x6b\x67\x98\xe1\xb4\x4e\x64\x52\xa6\x31\x6a\x04\x3b\x7a\x9c\xf2\x20\x35\x73\x7e\x61\x4a\x3d\x21\x61\x15\xe4\xc1\x64\xb5\xc2\x40\xf9\xe0\x5e\x91\x5a\x06\x94\x05\x34\x21\x51\x35\xa3\xf3\x43\xe4\x41\x35\x71\x31\x18\xf6\x48\x4c\x7e\xe3\x8c\x5c\xe9\x16\x1a\x3b\xeb\x76\x4c\xaa\xde\x95\x4a\x25\xfd\x52\xe0\x1d\x76\x32\x7a\x50\xd7\x54\x56\x33\x9b\x48\x18\x48\x2f\xab\x7f\xe5\x99\x75\x87\xe6\xad\xba\xef\xd3\xbe\x4d\x6a\xab\xa7\x34\x61\xa2\x67\x87\x46\x30\xb7\xe0\x9e\xa2\xb8\x0b\x36\x95\x01\xbf\x44\x31\xe7\x51\x34\x61\x61\xc2\x29\x53\x2d\x68\x8b\x74\x19\x53\xf5\x53\x63\x45\xf8\x4d\x98\xf4\x35\xb1\x0a\x38\xef\x8c\x3e\x74\x7e\xd2\xa6\xb6\x15\xb0\xe5\xc2\x6e\xe8\xfb\x95\xa2\xb9\x23\xf0\x9c\xd7\x3a\x90\xd5\xa4\xb6\x9b\x22\x83\x96\xe7\xb5\xa1\xd7\x78\xa3\xa4\x25\x29\xea\x34\xb8\x27\xda\x8a\x1c\x53\xb6\x16\x28\x9d\xb0\x98\x26\x73\xc1\x15\x0f\x78\xe4\x83\x0a\xca\x82\xf5\x54\xf0\x78\xce\x85\x79\xb5\x3b\x2c\x1b\xd6\x05\x6f\x01\x8e\x69\x28\xa6\x49\x3e\xf5\x97\xaf\x0c\x26\xd1\xf2\xbf\xc0\x38\xb3\x93\x7f\x93\x5d\x1e\x0d\x5a\xec\xe2\x02\x73\xbb\xb8\x2f\x6b\x27\xb3\x93\xa1\xf6\xd5\x79\xda\x52\xa7\x9a\xa6\xc9\xe4\xda\xd5\xb3\x5b\x85\x74\x44\x2c\x84\x29\xe4\xfb\xe5\xe1\xc3\xa3\x87\x39\x74\x61\xaf\xa1\x2a\x0c\xf5\x04\xf0\x0c\xd5\x48\x29\xeb\xbf\x5e\x06\x76\x0d\xec\x22\xd9\x14\x70\xb0\x34\x60\x38\x99\x9d\xfc\x19\x34\x6c\x08\xdf\xaa\x62\xdd\x0e\x93\x40\x4e\xa2\x65\x53\xb7\x88\xe8\x33\xf4\x8c\x93\xf0\x84\x44\x84\x05\x94\xad\xdf\x0c\x7d\xbf\x04\x64\x27\xd1\xa6\x9a\xf6\x12\x58\xfe\xff\x7a\xfd\x8f\xbe\x5e\xaf\x4d\xba\xb5\x41\x43\xc7\x41\xe1\xff\x99\xee\x4d\xac\xed\x75\xcd\xae\x38\xc8\x36\xec\x88\x01\x37\x4c\x46\x82\x95\x97\x18\x93\x68\x99\xa1\x64\xef\x98\x1b\x97\x27\x96\xf9\x8a\x8b\x2b\x22\xc2\xb2\x26\x11\xb1\x46\x65\x34\xa9\xd3\xcb\x08\x39\x18\xc5\xdc\x50\xab\x62\x65\xfa\x3d\xbf\xb8\x98\x17\xca\x37\x09\xdc\xda\x0c\x75\xa6\x2d\x43\x5f\x2e\xc4\x1e\x31\xe0\xce\x0d\xe2\x55\xaa\x92\xd4\xe6\x98\x9e\xfb\x5f\x8b\x6c\x3c\x73\x91\x37\x4a\x25\x7e\xbf\x6f\xae\x39\x26\xd1\xb2\x77\xfa\x72\x61\xc6\xe1\xfe\x01\x34\x7e\xc6\x31\x3b\x81\xd7\xe7\xb3\x46\x38\x68\x53\x57\xe8\x96\x56\xaf\x04\x40\x85\xd8\x48\xb0\xfc\x17\x07\x9a\x6e\x8e\x68\x7f\x2c\x34\xb9\xd6\x3a\xe5\x7a\x66\x47\x21\xab\x5a\x6d\x6a\xf7\x1a\xa2\xb4\xbc\x58\xdc\x73\x45\xb6\xf3\x67\x10\xce\x0b\xe2\x6f\x91\x29\xe7\x71\xf0\xaf\x00\x00\x00\xff\xff\xb6\x7e\x29\x97\x5d\x27\x00\x00")

func assetsClusterYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsClusterYml,
		"assets/cluster.yml",
	)
}

func assetsClusterYml() (*asset, error) {
	bytes, err := assetsClusterYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/cluster.yml", size: 10077, mode: os.FileMode(420), modTime: time.Unix(1484868722, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsRepoYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x92\x4f\x6f\xda\x40\x10\xc5\xef\x7c\x8a\x49\x14\x29\x52\x24\x13\xd3\x4b\x9b\xbd\xb9\x84\x44\x48\x51\x6b\x61\x48\xce\x93\x65\x80\x15\xfb\xc7\x1a\xcf\x96\xd2\x88\xef\x5e\x79\x43\xa8\xa9\x5b\x9f\xac\xf7\x7e\xf3\xbc\xfb\x3c\x59\x96\x0d\x8a\x97\x6a\x4e\xae\xb6\x28\xf4\x10\xd8\xa1\x3c\x13\x37\x26\x78\x05\xd7\x9f\xf2\x51\x9e\xe5\x77\x59\x7e\x77\x3d\x28\x91\xd1\x91\x10\x37\x6a\x00\x30\xa3\x3a\x7c\x43\x47\xed\x3b\xc0\x7c\x5f\x93\x82\x4a\xd8\xf8\x75\x12\xee\xa9\xd1\x6c\x6a\x49\x31\x2d\x0b\x1e\x1d\x0d\x66\xd4\x84\xc8\x9a\x52\xc4\x44\x37\xad\xd3\x4d\x28\x5e\x2a\xa5\x26\xe3\x99\x52\xad\xd3\x18\x09\xbc\x4f\x76\xc9\xa1\x26\x16\xf3\x3e\xd9\x3e\x7f\x80\x74\x0c\xb8\x98\xd1\xea\x74\xaa\x1e\x53\x06\x6b\xf4\x7e\x4e\x3f\xe5\x63\x1e\xa0\x7b\xcd\xfc\x4b\x36\xca\xb3\xd1\xe7\xeb\x93\x5b\x09\x0a\x39\xf2\x9d\x81\x0c\x2a\xb3\x54\x50\x58\x1b\x76\x65\x6c\x36\x65\xb4\xf6\x64\x02\x4c\x56\x2b\xd2\x72\xf4\x3b\x7a\xc9\xc6\x6b\x53\xa3\x55\x1d\x11\xd2\x5d\xe1\xf2\xe6\xb2\x23\x8e\x83\x5f\x9a\xd4\xd9\x39\xc9\xfe\xc9\x6c\xe9\x5c\x04\xc0\x5d\xa3\xaa\x54\x67\xc1\x5e\xc1\x45\x15\x5f\x01\xd9\xab\x56\x37\xe8\xd4\xd5\x5b\xaa\xb3\xd0\x3a\x44\x2f\xd3\xe5\x41\xc5\x86\xf8\xf6\xa6\x13\x53\xe8\xbf\xbf\x96\x01\x69\x56\x8f\x24\xf7\x61\xe7\x6d\xc0\xe5\x82\xed\x43\xe0\x27\xdc\x13\xf7\xb0\xaf\x28\x7a\xf3\x48\x32\x75\xb8\xa6\x7f\xbb\xe3\x0d\xe9\x6d\x9a\x2e\x7e\xa0\xb1\xf8\x6a\xac\x91\x7d\x8f\x2d\xe3\x7f\x42\xa6\xde\x88\x41\xa1\x14\xb1\xa8\xdb\x13\xf5\x98\x77\x39\x11\x25\xb2\xf4\xfc\x71\x70\xb5\xa5\xf3\x8c\xef\x51\xea\x28\xa7\x55\x5e\xf0\xf1\xef\x9c\x2d\xee\x82\x2d\x84\x15\xc8\x86\x80\xa9\x0e\x09\x78\x46\x1b\xe9\xd8\x76\xaf\xe1\xe1\x72\xcb\x43\xd2\x3c\x3c\x3a\x33\x5a\x9b\xe0\x0f\x43\x74\xf8\x2b\x78\xdc\x35\x43\x1d\xdc\xed\xd5\xdb\xc7\x9e\x1e\x06\xbf\x03\x00\x00\xff\xff\xa7\xe7\x4b\xa3\x81\x03\x00\x00")

func assetsRepoYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsRepoYml,
		"assets/repo.yml",
	)
}

func assetsRepoYml() (*asset, error) {
	bytes, err := assetsRepoYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/repo.yml", size: 897, mode: os.FileMode(420), modTime: time.Unix(1484868582, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsVpcYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x99\x4b\x6f\xdb\x38\x10\xc7\xef\xf9\x14\xd3\x60\x01\x6d\x81\x3a\xb5\xa5\xa4\x68\x78\xf3\x3a\x4e\x63\x6c\x9b\x18\x96\xe1\x02\x69\xf6\x40\x4b\xe3\x58\x88\x44\x6a\x49\x2a\x0f\x14\xf9\xee\x0b\x51\x92\x1f\xb4\xe4\xfa\x15\xaf\x8b\x5c\x6a\x72\xc8\x99\xff\x4f\x43\xce\x48\xad\xd5\x6a\x47\xcd\xef\x6e\x1f\xa3\x38\xa4\x0a\x2f\xb9\x88\xa8\x1a\xa0\x90\x01\x67\x04\x2c\xbb\xde\xa8\xd7\xea\xe7\xb5\xfa\xb9\x75\xd4\xa5\x82\x46\xa8\x50\x48\x72\x04\xd0\x61\x52\x51\xe6\x61\x1f\x19\x65\xde\x4b\x3a\x04\x70\x81\xd2\x13\x41\xac\xf4\xe2\xc2\x02\x54\x66\x02\x8a\x43\x22\x11\x46\x5c\xc0\xa0\xdb\xd2\x0b\xfa\x2f\x31\x12\x70\x95\x08\xd8\xbd\x1e\x68\x86\x21\x7f\x42\x7f\x40\xc3\x04\x65\xb6\x69\x0d\x7c\x1c\xd1\x24\x54\x93\x5f\x7e\xe0\x51\x85\x7e\xee\x52\xcf\x91\x19\x23\x57\x8e\xf5\x36\x25\x31\xb9\xc9\x90\xa1\x82\x91\xe0\x11\x3c\x8d\x03\x6f\x9c\x06\x45\x53\x63\x70\xdd\x2b\xa0\x9e\x87\x52\x9e\x94\x87\xf6\x2d\x60\x5f\x91\xdd\xab\x31\x01\xeb\xdc\xca\x86\xe8\xf3\x64\xa8\xf1\xd9\x9a\x0f\xa8\x7e\xa2\xff\x3e\xd6\x67\x85\x75\xa9\x52\x28\x18\x81\xe3\x3f\xef\xee\xfc\x9f\x8d\x0f\xce\xeb\xfb\xbb\xbb\x93\x55\x7e\x7c\xcc\xff\x69\xbf\xbe\x3f\xd6\x5b\xb6\x38\x93\x4a\xd0\x80\xa9\x39\x8d\x56\x94\x48\x05\x43\x04\x0a\x8f\x34\x0c\x7c\x68\x75\x2e\x7a\x30\x0c\xb9\xf7\x40\xe0\xf9\x44\xff\x7d\x7c\x3e\x49\xa3\x1d\xc4\x5e\x2b\xf0\xc5\x5f\x7a\xae\x92\x96\x5e\xba\xfc\xb1\xad\xcb\xa6\x51\xc0\x69\x7c\x3a\x5c\x3a\xdd\x64\x18\x06\x5e\x06\xa1\x79\xdb\x58\x8b\x54\xf3\xb6\xb1\x63\x52\xf6\xe9\xef\x42\xca\x5e\x93\x94\xbd\x43\x52\x8d\xdf\x8a\x94\xb3\x26\x29\x67\x87\xa4\xec\x83\x26\xd5\xe2\xcc\x0f\xd2\x35\xba\x08\x5c\x51\x69\x1c\xc6\x8c\xd7\xf1\x25\x23\xe4\x9a\xab\xe3\xec\x67\x5a\x1d\xf4\x50\xfb\xdf\x84\x86\xf2\x98\xc0\x8f\x77\x3d\x1c\x55\x1e\xe4\x0f\x60\x59\xff\x94\x6d\x6f\x6f\xb1\xbd\xfd\xeb\xed\x9d\x2d\xb6\x77\x8c\xed\x7b\x28\x79\x22\xbc\xac\x58\x0e\xba\x2d\x32\x93\x22\xcd\xef\x2e\x21\xed\x96\x4d\x48\x71\x71\x77\x05\x8f\x51\xa8\xa0\xa8\xad\x00\xd3\x0c\x04\xed\x6d\xb6\x24\xe4\x26\x6d\x46\x87\x21\x5e\x30\xe9\x26\x71\xcc\x85\x22\x60\x29\x91\xa0\x65\x4e\x5f\x71\xa9\x18\x8d\x50\x1a\x06\x66\xab\x90\x39\x32\x46\x73\xdb\x3e\xbd\x97\x53\x1c\x7f\xe3\x0b\x81\x6b\x1a\x61\x3e\x02\xa0\x1b\x03\x02\xef\xdc\x64\x08\x7f\xfc\xd4\x02\x5d\x45\xbd\x87\xd4\xe8\x75\xf1\xce\x2e\xa7\x91\x4d\x17\x79\x9a\xe5\x19\x29\x49\xb2\x0a\x64\x83\xd8\xeb\xf8\x05\xae\x1c\xec\x22\xc8\xaa\xa4\xcb\xcd\xbf\xd1\x38\xb3\xe8\xc4\x37\xec\x2b\x4d\x98\x37\x26\x90\x52\xcb\xe7\x9b\x8f\x34\x08\xe9\x30\x08\x03\xf5\x72\xcb\x99\xd6\x8c\x21\x7a\x0a\x7e\x40\xfd\x03\xbc\xfb\x92\xee\x2a\xf3\x0c\xab\x22\x87\xea\x89\x8b\x07\x13\x5e\xe6\x77\x53\xc8\xb5\x58\x2f\xaf\x35\x16\xef\xfd\xad\x68\xdb\x3b\xa4\x6d\xef\x92\x76\xe3\x10\x68\xdb\x8b\xb5\x63\x2b\xda\xce\x0e\x69\x3b\xbb\xa4\x6d\x1f\x02\x6d\x47\xbf\xe0\xa4\xc5\x10\xd5\x17\xaa\xf0\x89\xbe\x94\xd3\x36\x8c\x2a\xa0\xee\x2b\xfc\xac\x02\xac\x14\xf8\xa0\xdb\xca\xe7\x9b\x4a\x51\x6f\x1c\x21\x53\x6b\xa5\x84\xe1\x65\x62\xb1\x48\x24\xd3\xd4\xe3\x89\xc2\x7e\x5a\x29\xca\x03\x9a\xce\xaf\x15\xc6\x9e\x33\x63\x5e\xce\x12\x25\x79\xc3\x15\x23\xf3\xe5\x0d\x23\x25\xcf\xa5\x42\xe6\x94\xc3\x44\xad\x09\x30\xb7\xbc\x40\xa9\x02\x46\xd3\x03\x3e\x73\x3e\xe7\xdf\x3a\x01\x56\x7d\x3e\x93\x42\x35\xf5\xd3\x94\x92\x7b\x81\x76\xb0\xec\xae\x29\x5d\xb0\x71\x71\xcd\xe6\x0d\xed\xf3\x8b\x56\x85\x64\xd4\x84\x3d\x09\xab\xaa\x63\x4b\x85\xd9\x5b\x08\x73\xf6\x24\xac\xaa\x64\x2c\x15\xe6\x6c\x20\x2c\x3f\xc0\x4d\x2f\x2c\x17\x31\x9d\x3f\xf0\xab\xa2\xc3\x86\x3c\x61\x7e\x3b\x1e\x63\x84\x82\x86\x5d\x2e\x94\x29\xb1\xcd\x94\xa8\xb8\xa3\x0d\xa3\x0a\xb1\x53\x2b\x83\xac\x81\x09\xa0\x97\x84\x78\x9d\x44\x43\x14\xe9\x4b\x61\xdd\x29\xfa\xf3\xae\xe0\x8a\x7b\x3c\x24\x60\x7d\xb2\x66\x6c\x9b\x5e\x96\x09\xfa\xfb\x58\xd1\xec\xdf\x0b\x94\x69\x83\x3f\xa2\xa1\x9c\x74\xf8\x4b\xee\x9f\x54\x73\x8f\xb2\x7b\x24\x13\x7a\x97\x82\x47\x3a\x02\xfb\xd4\x9a\x0c\xf6\x79\xea\xfe\xec\xcc\x39\xb3\xa6\xe4\x5c\xf7\xea\x70\x78\x9d\xbe\x09\x2f\x1d\x40\xf1\xc5\xf2\x97\xcc\x6c\xdb\x20\x96\x0d\xe4\xb8\xae\x94\x8a\x0f\x87\xd7\xd9\xff\x9c\x5f\x9f\xeb\x06\xab\x6c\xe0\x26\x51\x1a\xd6\xe1\x80\xaa\x6f\x05\x6a\xf6\x4d\x7b\x23\x4e\x26\xa6\xc9\x21\x34\x6a\xaf\x29\x66\xc5\x62\x53\xba\xe0\x6d\xdb\x83\xd5\x9e\x84\x51\x81\xf7\x2a\x6f\xab\x26\x61\x13\x79\xce\x5e\xe5\x6d\xd5\x2a\xac\x22\xef\x26\x51\x71\xa2\xb2\x6f\x5e\xba\xd8\xe7\xfd\xf6\xcc\xa7\xc6\xfe\x18\x21\xf0\x81\x8f\x40\x8d\x11\x1e\xe3\xac\x9e\x17\x95\x7b\xb6\x35\x68\x3f\xeb\xaf\x5a\x85\x7b\x1a\x55\x97\x76\xed\x6c\xf1\x68\x54\x06\x90\xb5\x02\x20\xb5\x61\xe0\xcf\xfd\xd7\x40\x16\x4a\xee\xf5\x92\x11\xd2\x19\x4d\xdb\x93\x8a\x03\x91\x4e\x2d\x49\xfc\x7c\x52\x47\x7d\xcd\xb5\x83\x75\x15\x2e\x28\x5b\x3c\x29\x6b\xaa\xb5\x37\x50\x6b\x2f\x53\x6b\xbf\x95\x5a\xbb\x44\xad\xb3\xa6\x5a\x67\x03\xb5\xce\x32\xb5\xce\x5b\xa9\x75\x3a\xfe\xd1\x7f\x01\x00\x00\xff\xff\x7b\x24\x36\xe9\xff\x1d\x00\x00")

func assetsVpcYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsVpcYml,
		"assets/vpc.yml",
	)
}

func assetsVpcYml() (*asset, error) {
	bytes, err := assetsVpcYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/vpc.yml", size: 7679, mode: os.FileMode(420), modTime: time.Unix(1484356037, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/cluster.yml": assetsClusterYml,
	"assets/repo.yml": assetsRepoYml,
	"assets/vpc.yml": assetsVpcYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"cluster.yml": &bintree{assetsClusterYml, map[string]*bintree{}},
		"repo.yml": &bintree{assetsRepoYml, map[string]*bintree{}},
		"vpc.yml": &bintree{assetsVpcYml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

