// Code generated by go-bindata DO NOT EDIT.
// sources:
// completion.sh
package completion

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

var _completionSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x7b\x73\x1b\xb7\x11\xff\x9f\x9f\x62\x73\xe4\xc4\x92\xac\xb3\xac\xb4\x7f\xc9\xa1\x26\x89\x2d\xb7\x9e\x26\x51\xc6\x6a\xda\x69\x55\x0d\x07\xbc\xdb\xe3\xa1\x02\x81\x1b\x3c\x48\x33\xaa\xbe\x7b\x67\x81\x7b\xf2\x8e\x8a\xd5\x2a\x56\xd2\x8a\x33\x1e\x93\x78\x2c\x16\xfb\xf8\xed\x02\x5a\x8c\x3f\x3b\x9a\x73\x79\x34\x67\x26\x1f\x8d\xc6\x30\x9b\xa5\x89\x32\xb3\x2f\x73\x14\x05\x6a\xc8\x9c\x4c\x4e\xa9\x39\xb4\x26\x82\x83\x71\xf3\x44\x2d\x97\x4c\xa6\xa7\xa3\x51\x39\x3c\xc5\xb9\x5b\xec\xed\xc3\xcd\x08\x00\x80\x67\x70\x79\x09\xb1\x84\xc9\xcd\x9b\xd7\xe7\x17\xb3\xd7\xe7\xdf\xfd\x30\x7b\x73\xf6\xcd\x8f\x7f\x98\xbd\x7d\xf7\xed\xd9\x2d\x5c\x5d\xbd\x02\x9b\xa3\xf4\xa3\xe9\x83\x49\xae\x20\x9a\xdc\xbc\xfd\xf1\xfb\xd7\xdf\x7f\xfd\xdd\xd9\xe5\xf1\xd5\xed\x09\x4c\x0e\x22\x38\x3d\xa5\xf6\x41\x32\x91\x9f\x9e\xf1\xd1\x2d\x31\x7e\xc1\x65\x82\x30\x5b\xa0\x9d\x25\x6a\x59\xcc\xd6\x4a\xa7\x66\x36\xdf\xcc\x34\x66\xc0\x0d\x48\x65\x81\xad\x18\x17\x6c\x2e\x10\xb8\x84\x9f\x4c\x0e\x6b\x04\x89\x98\x02\x83\x42\x69\x8b\x29\xac\x50\x1b\xae\x24\xd8\x9c\x59\x48\x98\x84\x39\x82\x51\x4e\x27\x98\x8e\xc6\xb0\xce\x51\x82\x33\x5c\x2e\x68\xf6\x8b\xd1\x88\x67\xf0\x99\xdf\xeb\x4f\x30\xb9\xf9\xfb\xc5\x1f\x67\x7f\x39\x7b\x7f\xf1\xee\xfc\xfb\x93\xe7\x06\x2d\xed\x13\x3e\xff\x1c\x3e\x03\xbb\x29\x76\xb2\x76\x7a\x94\xe2\xea\x48\x3a\x21\xe0\x8b\xd3\xcf\x8f\x4b\xc1\xec\x18\x5c\x8b\x98\x3e\xeb\x9c\x0b\x84\x4b\x98\x8c\x21\x5e\x58\x78\x09\x57\xaf\x20\x55\x75\x77\xf5\x49\x98\x41\x88\x26\xc7\x11\x70\xd9\xef\x74\x7a\xbf\xd7\xd8\xea\x9c\x4e\x6e\xbc\xd8\xff\x7a\xfe\xfe\xcd\xc5\xa5\xff\xfa\x9a\xbe\x5f\xdd\xee\x9c\xf5\xea\x55\xaf\xab\xd0\xb8\xda\xbd\x0c\xf5\xee\x5a\x27\x3e\xbe\xdf\x4a\x5e\x5a\xbb\x97\xf2\xdd\xd3\xbd\xa8\xb3\xda\x57\x57\xb7\xd1\xee\x29\x03\x8b\x24\x44\xe6\x0e\xb1\x51\xf7\x74\xd2\x6c\xe2\x3e\xb4\x63\xb9\x9b\xf0\x18\x98\x31\x6e\x89\x50\x33\xff\xcd\xfb\xb3\xaf\xff\x74\x41\xe6\xcd\x84\x46\x96\x6e\xc0\xa0\x05\xc3\x24\x8a\xcd\x4e\x32\x26\xe7\x99\xbd\x0f\x4f\x68\x58\xd2\x6b\xec\x52\x49\x95\xc4\xd1\xed\x28\xe3\x23\xf2\xc5\x37\x98\x31\x27\x2c\xcc\x31\x67\x2b\xae\x34\x58\x05\x0b\xb4\x64\xda\x20\xf1\x83\x85\x12\x40\x20\xd3\x6a\xe9\x5b\x13\xa7\x35\xca\xba\xe3\x05\xbc\x93\xa1\x9d\x8c\x57\x65\x10\xc7\x04\x48\x40\x73\x96\x6c\x33\x47\x50\x36\x47\x0d\x86\x5b\xc7\x2c\x57\xd2\x8c\xc6\xe4\xd3\x34\x25\x73\xd6\x69\x0c\x1e\x6c\x72\xe5\x44\x0a\x28\x53\xa2\x5c\x08\xa4\xb1\x87\x60\x73\x6e\x40\xa3\x75\x5a\x1a\x38\x3e\x0c\xc4\xd6\xdc\x20\xbc\x1c\x8d\x47\x63\xf8\xda\x8b\xd9\x00\x83\x49\xc5\xe9\x8a\x69\xee\x91\x03\x3f\x70\x63\x4d\xb5\x58\xc2\x84\x20\x38\x48\x94\xb4\xb4\x31\xbf\xea\x9a\x0b\xe1\x5b\x58\xc5\x92\x72\x32\x6d\x01\xe7\xd6\x12\xf7\x22\xce\x4d\x5b\x60\x04\x4a\x4a\xa7\x30\x47\x1a\x88\x2b\x26\x1c\x23\x1c\xab\x05\xcb\x65\xe1\x2c\x18\xab\xa9\x7f\xcf\x6f\x9c\x1b\x48\x31\xe3\x12\x53\x5a\xa8\x44\xf0\x25\xf1\x6a\x54\x29\x31\xf9\x8c\x94\x47\xb0\xa8\xd5\x5c\xe0\x72\xbf\x01\x7a\xaf\xd9\x59\xb9\x91\x59\xc1\xb4\xc1\x1a\x95\x2a\x44\x8a\x26\x49\x04\xb1\xb0\xf4\x85\xd8\x8b\xb6\xa0\x89\x4f\xa3\xc9\x8d\xf7\xc5\xcb\xe4\xaa\xc4\x71\x68\xa0\x8a\xf7\xa0\x2a\xa8\xff\x5f\x71\xde\x77\x8e\x31\x85\x1b\x6f\x1c\xe4\x05\x90\x09\xb6\x38\x84\x44\xcd\x35\x03\x63\x55\x61\x80\x58\xa4\xcd\xf3\xe5\x12\x53\xce\x2c\x8a\x0d\x6d\x74\x8d\xa5\x05\x40\x8e\x1a\x7b\x64\xcb\xbe\xe3\x5e\xc7\x96\x7f\xc4\x07\xfb\xdb\x4d\x07\x7d\x26\x4b\x71\x4d\x69\x6f\x7d\xff\x1a\x93\x0d\x70\x03\x98\x65\x98\x58\xbe\x22\x0e\x17\x42\xcd\x99\x20\x46\x83\x49\xb1\x0d\xf0\x10\x91\x12\xcd\xd6\x02\x6c\xae\x95\x5b\xe4\x5e\xc7\x4c\x2f\xdc\x12\xa5\x35\xc0\x82\x71\x14\x5a\x2d\x34\x5b\x0e\x2c\x24\xd9\x8a\x2f\x98\x45\x53\xba\x8a\x4c\xac\x0f\x77\x1a\xd1\x3b\x97\x51\xc4\x8b\x37\x60\x26\xd6\x6c\x63\xc8\x0c\x6a\xb7\x75\x92\x40\xc6\x5b\xdc\x00\x6d\x96\x59\xca\x18\xb8\x4c\x49\xdc\xdb\xbe\x7e\x48\x12\xe7\x32\xd1\x48\xac\xd2\x2a\x73\xcc\x94\x46\x98\x6b\x64\xd7\x34\x43\x39\x4b\xae\x4e\x13\x85\x52\x45\x6f\x85\xbd\xbd\xe4\xf9\xf3\xfd\xbe\x70\x3d\x81\xbb\xf4\xd4\xc1\xaf\x36\x15\x8f\x5a\xd0\x68\xfb\x25\xa5\x12\xa5\xa5\xe7\x4c\xa6\x02\x7d\x0c\xd6\x58\x88\x4d\x6d\xe4\xed\x94\x07\xa2\x06\x56\x8c\x4f\x58\x46\x7e\x8c\x50\x09\x13\xc0\xa7\x2f\xc3\xcf\x4c\xe9\x16\xfe\x90\xd7\x45\x93\xaf\xa2\x8e\x4b\x8c\x21\xe3\x82\xe4\x47\x52\x68\x11\x25\x9f\x4a\x72\xe2\xf4\x99\x85\x25\xb3\x49\xde\xc1\xcb\x8e\x2a\x42\xde\x35\x69\xad\x34\x9d\x92\x0b\x3a\x1d\x1d\xf4\x73\xae\xb0\x6a\x92\x63\x72\x4d\x33\x3d\xd5\x66\xa6\xca\x32\xd4\x84\x10\x95\x53\x05\x3b\x64\x49\x82\x05\x19\x9a\xac\xad\xae\xe4\x90\x1b\x58\x32\x7d\x8d\x29\xcc\x37\xd4\x3d\xdd\x5a\x88\x67\x60\xd4\x21\x30\x30\x05\x4b\xb0\xca\xc8\x96\x4c\x3a\x26\xc4\x06\x58\x9a\x62\x0a\xc6\x27\x71\xc4\x8a\x33\xa8\xbd\x4f\x7c\x28\x30\x21\x44\xb3\x8a\xc6\x40\xe6\xb4\x47\x7e\x8f\x6b\x9d\x25\x4a\xf0\x68\xf6\x30\x98\xf0\xc4\xf1\xc1\x74\xc0\x41\xe9\x43\x51\xf5\xfd\xd9\x0f\xdf\xfe\xed\x92\x3f\x7f\x7e\x35\xed\x90\x1a\x9c\x30\x10\x2a\xef\x4f\x1a\x3e\x8a\x76\xc7\x84\x33\xde\x58\xaf\xcf\x7d\x5f\x33\x21\x4c\xf0\xb6\x26\xc8\x90\x1e\x94\xd3\xed\x28\xb2\x42\xe9\x8d\xea\xc5\x68\x0c\x7f\x3e\x7f\x73\x7e\xd2\x28\xd1\x9b\xbb\x1f\xc6\xa4\x47\x7d\x36\x17\x1b\x72\x7e\x5a\x05\x96\xe4\xa8\xf8\xa1\x10\x3c\xe1\x56\x6c\x68\x3a\x45\x12\x56\x06\xb1\x10\xe3\x84\x50\x6b\xa2\x50\x45\x33\x13\xc2\xd9\x76\x34\x33\x89\x2a\x02\xd8\x30\x4d\x26\xa7\x35\x26\xf6\x64\x34\xae\x80\xc2\x10\x57\x9a\x6d\x08\x0b\x9a\xdd\x98\x10\xb4\x6b\xc4\xca\x95\x48\x4d\x33\xe9\xa4\xb2\x5f\xbf\x73\xab\xfc\x6a\xc0\xa5\x55\x5b\xfe\xdc\x50\xf4\x0e\x3d\xaa\x8d\x93\xd5\xb3\xd7\xcc\xc0\x82\xaf\x50\x1e\x96\xde\x11\xe2\xae\x0f\xbd\x64\xf7\x89\x75\x4c\xd4\xa3\xe9\x9f\x5f\xcc\x23\xb1\x31\x2a\xe1\x21\x02\x97\x9c\x36\x00\xb0\x4c\x83\xe7\xdf\x54\x5b\xf2\x99\x67\x37\x30\x96\x1e\xbc\x4c\x4b\xd7\x0d\x23\xa3\x61\xef\xed\x22\xd1\x76\x92\x71\xe2\xe9\x94\x80\x54\x7d\x02\x30\xb5\x8c\x84\x8c\xa6\x8a\xe6\x5b\x1e\x6b\xd6\xac\x80\x98\x78\xae\x75\x20\x19\xa9\x9b\x76\x33\xeb\x0c\x6e\x08\x4e\xeb\xed\x1d\x1d\xc5\x47\xb3\xdb\xd1\x1d\x1c\x57\x16\x41\x6a\x6a\x34\x3b\xb9\x11\xcc\xd4\x3c\xdd\xce\x26\x37\x0d\xf5\xdb\xae\xb3\xb4\x99\x9f\xde\x39\xaf\x2b\x84\xd6\xb8\xe9\xa4\x23\x81\xce\x38\x9e\x41\x8a\x89\x20\x33\x8d\x33\x88\x3a\x23\x23\x38\x85\xfa\xe4\x36\xa0\x1b\xfa\x4c\x76\x0a\x17\x85\xe9\xa7\x1c\x5d\xe1\x4c\x6e\xda\xb3\x6f\x21\x55\x18\x20\xd3\x3b\x55\x57\x0e\x25\x20\x54\x9f\x10\xcd\x76\xc1\x45\xe5\x10\xe1\x58\x3b\x2b\x84\x5b\x70\x39\x6b\x85\x9c\x3a\xd2\x91\x9e\x53\xae\x77\xc4\x2b\x73\xcd\x0b\x12\x51\x20\xe0\xf9\xa3\x18\x95\xb3\x15\x06\x67\xaa\x00\x2e\xe5\xe4\xe1\x4a\x6f\xb6\xcc\x3c\x4e\x61\x42\xe4\x07\x6d\x9b\xd6\xce\x78\x38\xa9\x4f\xf6\x28\xa9\x80\x88\x46\xfb\x8b\x8a\x08\x62\x7f\xa4\xce\x20\xe6\x64\x92\x10\x1d\xbc\x30\x79\xb4\x3f\x78\xfe\xed\x8a\xd5\xef\x3a\xc0\x61\x13\x63\x7d\xb6\x3c\xa1\xe5\xfa\x60\xfc\x22\x74\x74\xda\xeb\xbc\x61\x50\xbc\x7e\x39\xe6\x6c\x5e\x0b\xb2\x4c\x07\xa0\x63\x65\xfe\xee\xe0\xae\xb4\x7a\x4b\x2a\x2d\xad\xd2\x31\xab\xa1\x5b\x79\xe7\x74\x2f\x12\xdc\xd8\xb8\xd0\x6a\xc5\x53\xd4\x26\x82\x48\xa8\x05\x97\xe1\x7f\xe5\x6c\x79\xcc\x0d\xd3\x28\xa8\xd3\x9c\x90\x5b\x47\xfb\x35\x57\xfe\x3a\xa3\x8d\x3d\x5b\x8c\x54\x81\xd6\xe9\x81\x3c\x7d\x20\x00\xee\x48\xa8\xc8\xc4\x3d\x0f\x1e\x06\x7f\x2e\xcf\xbe\x1f\xdd\x0e\xc2\x7e\x74\x6a\x38\x20\xe1\x5d\xc1\xc3\x27\x8a\xb5\xa2\x67\x5e\xcc\x9f\x48\xdd\xa5\xde\xa0\x52\x9c\xef\x8a\xe2\xb8\x60\xc6\xf8\x9b\x87\x5e\x4b\x4c\x06\xdc\x6a\xd6\x7c\xc5\x2c\xc6\xd7\xb8\x69\x37\x06\xa3\x69\x5a\x28\x03\x23\xe7\x2a\x5b\x1e\xca\x3e\x86\xb2\xaf\x71\x99\xdf\x1a\x57\x14\x4a\xdb\x90\x6f\xd6\xa9\x49\x0b\x4a\x36\xd8\xbf\xb8\xd8\x42\xbb\x1d\xf6\xf3\x4b\x5a\xe6\xc7\x1a\xd4\x96\xcd\x70\x63\x67\xb5\xaf\x7e\x5a\xe3\xa9\x6c\x87\xf4\xfc\x4f\xa3\xe4\x83\x79\xff\xff\xaf\x76\x7b\x80\xa0\x9c\x7d\x1c\xa5\x3e\xe9\xf2\xc1\x74\x99\x08\x67\x2c\xea\x4f\x1f\xc8\x99\xb5\x2c\xc9\x23\x88\x4a\x37\x25\xb0\x88\x20\xd2\xb8\x54\x2b\xf4\x5f\x08\x9a\x23\x88\x0c\x5a\x57\x3c\xc5\xf5\x6d\xd2\xff\x7d\x5c\x2f\x55\x3f\x0b\x9a\x78\xf2\xe4\xce\xe7\x37\xeb\xc9\x3e\xea\x3e\x5e\xac\x0d\xc6\x84\xe9\x53\xe0\xfd\x05\x55\x1c\x40\xf2\x11\x95\x2c\x84\xff\xdf\xc9\xfa\xef\xcd\x4f\x6a\xfe\x25\xd4\x4c\x21\xf0\xb7\x0d\xcc\xbf\x3a\xa1\xfa\x74\xe2\x71\x64\xea\x7b\xa2\x38\x4e\x58\x9c\xa0\xb6\xa6\x39\x89\x72\x69\x30\x71\x1a\xeb\x86\xd6\xb1\x94\x7e\xa9\xd8\xdf\xd3\xb6\x1b\xc2\xa5\xd4\xd3\xd9\xf8\xb7\xe7\xdd\x4a\x66\x7c\xf1\xe9\x13\x6e\x83\x94\x5f\x9b\x5c\xad\x23\x88\x9c\xa4\x9f\x4f\x49\x35\x3c\x78\x52\xed\xb5\x4b\x28\xf3\x84\xdb\x83\x7a\xfb\xcf\xdc\x65\x46\x76\xfb\x24\xd1\x87\x94\xa8\x87\x80\x27\x91\x3e\x88\x48\x43\x30\x7e\x84\x4b\x94\x34\xed\x5d\x9d\x3c\x81\x3a\x3c\x34\xa8\x97\x7f\x41\x64\x69\xfa\x78\x67\x2e\x57\xa4\xcc\x3e\xdc\x31\xeb\x57\xe6\x3a\x8f\x7c\x6b\xf1\xa0\x17\x15\xbf\x36\xd1\x3e\xe6\x6d\xc1\xff\x9a\x4c\xfb\x52\x9c\x1e\x7f\x02\x98\x77\x36\x8f\x20\x2a\x8f\xb0\xf4\xcd\x87\xf0\xe6\xf2\xbc\x3c\x0c\xde\x01\xfd\x64\xe5\x65\xe9\x7f\xa5\x94\x30\xae\xa9\x8e\x08\x05\x30\xcd\x79\x6b\x96\x72\x6d\xc2\xd0\x50\xf6\xfa\xee\xed\xc5\x74\xf2\xec\x1f\xf2\x19\xf8\x22\xc9\x58\x83\xe0\x12\x5f\x41\xaa\xb6\xa9\x3c\x9f\xee\x45\x13\xea\x0c\x15\x0b\x12\xe1\x4b\xf8\x72\x8f\x64\x52\x15\x53\x90\xbf\x43\x1c\xd7\xcb\x7e\xd1\x2a\x36\x09\xdb\xe8\xd0\xba\xd9\x5a\xa0\x55\xd2\xde\xad\x7e\x78\xeb\x6b\x84\xca\x45\xaa\xd1\x11\xec\xa0\xf0\x91\xbb\xdb\x92\xc9\x7d\x76\x57\xce\x8b\x69\xde\xc0\x26\xbb\xcc\xff\x50\xb3\xbd\x5d\x59\xc2\x31\x6c\x62\x8b\x93\xc6\xc8\x7f\xae\xe4\x65\xe7\x64\x3f\x7b\x0c\x52\xb5\x6b\xa6\xaa\xe2\xb0\x30\x01\x01\xb9\x2f\x49\xf4\x36\x05\x4a\xb7\xab\xd6\x1e\xc6\xbf\x87\xef\x05\x42\x05\xdf\xc0\xed\x40\x4b\x42\x9a\x2f\x72\x0b\x52\xad\x07\xe6\xfb\x82\x3a\x5f\x56\x24\x90\xad\xd0\x57\x9e\xfa\xea\x61\x65\x51\x5a\x4e\x5e\xa2\x34\xa4\x68\x31\xb1\x5c\x2e\xca\xfd\xf9\x62\x38\xcb\xae\x11\x56\x4c\x38\x34\x30\x77\xd6\x57\xf2\x19\x2c\x98\xf6\xa5\x6f\x82\x5f\xf7\xab\x9b\xc6\x10\xc7\x9e\x3b\x3f\x0d\xb8\x34\x96\x0c\xc9\xd7\xf6\x53\xfb\xd4\xb7\x0f\x4c\x5b\xe3\x33\x8d\xbe\xf2\x69\xad\xb4\xde\x10\x27\x6c\x4e\xbc\x56\x37\x21\x5b\x97\x20\x60\x73\x5f\x19\x6d\x14\x70\xfb\xcc\x80\x61\x19\x82\x55\xc0\x17\x52\x95\xef\x02\x7a\xab\x7c\xdc\x3d\x09\x41\x45\x5f\x11\x8f\x7f\xc7\x32\x20\x34\xa9\xfa\xe5\x8c\x24\x11\xef\x7a\x24\xbc\x42\x19\xc3\xe7\x02\x7b\xc6\xda\xa5\xc3\x25\xf9\xbb\x00\x67\xd8\x02\x0f\x9b\xe7\x10\x65\x85\xb8\x51\xfe\x6d\x85\x2b\xca\x07\x08\xed\x6a\xf4\x72\x51\xab\xda\xa5\x99\x87\x5e\x91\xc6\x69\x1c\xbc\xaf\x1a\x43\xae\xd6\xb0\x46\x58\x33\x69\x69\x6a\x10\xcd\xb0\xd2\x1e\x3d\xfb\x6d\x5e\x4e\xec\x28\x0e\x9f\x7c\x55\x23\x88\xc9\x51\x88\x50\x4e\x9a\x72\xc3\xe6\x02\xa7\x17\xaf\xbf\x78\xf9\xbb\xdf\x97\xfd\x65\x23\x38\xe9\x0c\xb6\x9e\x9b\xac\x99\x96\x5c\x2e\x0e\xfd\x2b\x29\xe0\xa6\x1a\x50\x3f\x5a\x28\xc3\xa1\xd3\x61\x84\x7f\x55\x11\x9e\x20\x75\xea\x1c\xcb\x65\x48\x27\x9d\x76\xaf\x1c\xad\x94\xa5\xe9\x19\xff\x10\x68\xb6\x0b\x24\x23\xbf\xa5\x72\x23\x75\x21\xf3\x74\x6f\xbf\xda\x5a\xb7\x30\xb7\xfb\x84\x65\xce\x4c\xde\x02\x7a\x28\x58\x72\xcd\x16\x58\xd6\x8c\xbf\x83\x39\x0a\x8e\x2b\x84\xa5\x33\xb6\x24\x37\x0f\xe0\xc0\x84\xc0\xb4\x06\x15\xb1\x09\x45\x7a\xfe\x99\xa2\x17\xf4\xf0\x03\xb9\x58\x42\x34\x3d\x89\x06\xe5\x31\x1a\xd0\xd0\x85\x65\xda\x76\xcb\x01\x41\x49\x20\x20\x2e\xab\x02\xcb\x77\x2a\x07\x75\x34\xf0\xf3\x49\xff\x35\xfe\xc7\xca\xf3\x55\xe6\x33\xf4\xb3\xf5\x55\xaa\x50\xf0\x1e\xbf\xed\x3c\xb5\xf1\x44\xfe\x1d\x00\x00\xff\xff\x80\x5d\xb4\x3d\x7c\x39\x00\x00")

func completionShBytes() ([]byte, error) {
	return bindataRead(
		_completionSh,
		"completion.sh",
	)
}

func completionSh() (*asset, error) {
	bytes, err := completionShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "completion.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"completion.sh": completionSh,
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
	"completion.sh": &bintree{completionSh, map[string]*bintree{}},
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

