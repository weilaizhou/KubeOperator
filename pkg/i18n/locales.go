// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xcd\x6e\xe3\x36\x10\xbe\xeb\x29\x26\xf2\xb5\x58\xf4\xec\x1b\x23\x31\x89\xba\xb2\x64\xe8\x67\xdb\xf4\x22\xd0\xd2\x58\x66\x23\x93\x02\x49\x35\xcd\xde\xfa\x5e\x7d\xa7\xbe\x42\x41\x89\x96\x95\x38\x41\xbc\xe8\x29\x30\xc2\x6f\xe6\x9b\x99\x6f\xbe\xd1\xaa\x96\xc7\xa3\x14\x5e\x42\x36\xb4\xa2\xbf\x45\x79\x91\xaf\xc1\x4f\xd8\x11\x81\x75\x0a\x59\xf3\x02\xf8\x17\xd7\x46\xfb\x5e\xb4\xad\x92\xb4\x38\x3f\xda\x76\xc8\x34\xc2\x9e\x77\x1d\x70\x01\xe6\x80\xd0\xc9\x9a\x75\x10\x6d\x41\x4e\xbf\xf5\x8b\x36\x78\x04\x8d\xc6\x70\xd1\x42\xcf\x5a\xf4\x3d\x6f\x55\x77\x83\x36\xa8\xbc\x20\x2e\xf3\x82\x66\x55\x48\x63\x5a\xd0\xea\x8e\x44\x31\x0d\xd7\xe0\xd7\x4c\x80\x90\x06\x1a\xec\xd0\x20\xb8\xe7\x36\x49\x3d\x28\x85\xc2\x80\x36\xcc\xa0\x3f\x07\x88\xf2\x91\x5a\x56\x26\x49\x94\xdc\xaf\xc1\x2f\x0e\x0b\x98\x1e\x83\xa9\x41\x08\x2e\xda\x0b\x50\x9c\x06\x24\x5e\x83\x1f\x1d\x7b\xa9\xcc\x8c\xaa\x99\xb0\xa8\x1d\xc2\xd0\xb7\x8a\x35\xd8\x8c\xcc\x15\x36\x28\x0c\x67\x9d\xf7\x8a\x74\x95\xd1\x3c\x2d\xb3\x80\xae\xc1\xbf\x63\xbc\xc3\x06\x8c\x74\xfc\x6f\xa0\x38\xa0\x42\xcb\x83\x09\x60\x5a\xcb\x9a\x33\x83\x0d\x1c\xa4\x36\x30\x88\x06\x15\x98\x03\xd7\xf0\x84\x2f\xfe\x07\x61\xab\xdf\xd3\xe4\x87\x62\x7f\x97\x02\xdf\x89\x7d\x47\xca\xb8\xa8\x82\x8c\x86\x34\x29\x22\x12\x57\x01\x49\xc6\x2e\x4c\x69\xd7\xe0\x87\xb8\x67\x43\x67\xe0\x5c\xe9\xa2\x15\x53\xd2\xb1\x13\x96\xbc\xf7\x90\xe6\x45\x45\xe2\x8c\x92\xf0\xf1\x2c\x8c\x07\x5b\xd7\x5b\xf5\xb8\xba\x46\xc4\x3c\xe8\x77\xcb\x99\xfa\x62\x2b\x72\x21\x16\x65\x3d\x73\x73\x18\x95\xe5\xc6\xf4\x5e\xdc\xea\xf6\xb1\xda\x66\xe9\x2f\x34\x28\xfe\x57\x8a\x5e\xc9\x3f\xb0\x36\xbe\x97\x3f\xe6\x05\xdd\x54\x4e\xff\x77\x69\x99\x84\x9f\xc8\x7f\xcf\x95\x36\xff\xfe\xf3\xb7\xef\x25\xa9\xc5\x91\x6f\x24\x8a\xc9\x6d\x6c\x1b\x9c\x48\x88\x7a\x60\x7f\x32\xde\xb1\x5d\x37\xee\xc3\xa0\x51\x79\x5b\x92\xe7\xbf\xa6\x59\x38\x26\xd9\x90\x22\x78\x70\x3a\xee\x99\xd6\xcf\x52\x35\x96\x2f\x17\xb5\x54\x6a\x64\x95\x66\xd1\x7d\x94\x90\xf8\xe2\xbd\x54\xbc\xe5\x82\x75\x1f\x01\xcb\xfc\x2c\x7e\x12\x14\xd1\x37\xea\x80\x96\xc6\x69\x5b\x50\x58\x72\xcd\x0d\xb8\x32\x6b\x29\x0c\xab\xcd\x58\x26\x6b\x8e\x5c\x70\x6d\x14\x33\x52\xdd\xb8\x80\xcb\xd6\x24\x12\xf4\x50\x1f\xc6\x80\x63\x17\x48\xb8\x89\x92\x4b\xad\xd9\xa4\x8d\xd3\xdb\x18\x74\xa2\x70\xa1\xb7\x9b\xd7\xa4\x33\x1a\x93\x82\x86\x8b\x21\x97\x16\x76\x60\x96\xfa\x72\x94\x6e\x82\x23\x85\x38\x24\xdb\x99\x41\xb9\x0d\xc9\xcc\xa0\x6b\x58\xff\x36\x31\x36\x7c\xca\xeb\xad\x14\xb6\x5c\x8a\x93\xce\x32\x7a\x1f\xa5\xc9\xb5\x5b\x0f\x13\xf8\x33\xa5\xd9\x65\xb5\xa9\xec\xdf\x53\x22\xbb\xf0\x57\xa7\x19\xb7\xfd\x33\x39\x77\x4c\xbc\x36\xbf\x32\xb6\xbe\x17\x4c\x45\xb7\x68\x96\x8b\x75\x1e\xfc\x01\xeb\xa7\xe9\x3f\x52\xec\x79\x3b\x28\x66\xb8\x14\x63\x4f\xa3\x0d\xb9\xa7\x1f\x87\xe2\x47\xd6\xe2\x55\x81\x3c\x6f\x25\x7b\x14\xda\xb0\xfa\xc9\xbb\xa7\xc5\xa9\xcd\x34\xcb\xd2\x6c\x12\x94\xeb\xe4\x5e\x0e\x62\xf4\x1f\x37\xdb\x0d\x1e\x77\xa8\x66\x79\x90\x30\x5c\xca\x61\x87\x28\x80\x35\xce\xbc\x1d\x64\xb6\x41\xa7\x9f\x8f\x3d\xd0\x01\xde\x33\xc0\x13\xf6\x81\xe4\x95\xeb\xa9\xf5\x04\x07\x58\xf4\x7f\xee\x67\xf0\x8e\xaa\xbd\x95\x90\x0d\x7a\x49\x1a\xd2\xd9\x46\xdd\x11\xab\x0a\x92\x7f\x5d\x83\x4f\x9a\x06\xec\x23\x90\xea\x74\x0f\xc7\x9f\xa7\x51\xbb\xb3\xf6\x53\x3f\x75\xf9\x99\x71\x03\xdc\x40\x23\x05\x7e\xb1\x09\x76\xac\x7e\x1a\x7a\x52\xd7\x72\x10\xc6\xdb\x92\x8c\x6c\x2a\xba\xd9\x16\x8f\xf6\xe4\x09\x3d\xec\xf7\xbc\xe6\xf6\xa4\xf6\x4c\xb1\x23\x1a\x54\xda\xba\x56\x51\xe5\xe5\x76\x9b\x66\xe3\x76\x09\x3d\xf4\xf6\x38\x5a\xe1\xbd\xf4\xf6\xee\x3e\xd0\xe0\xeb\xd9\xc6\xbf\xa1\xe2\x7b\x5e\x8f\xf3\x84\xfd\xa8\xd0\xd9\x98\x6f\x49\xf0\xb5\xdc\x56\x24\x08\xd2\x32\xf9\x11\x8b\x7e\x45\xfc\x6a\xaf\xf6\x56\x56\xe7\x6f\xae\xe8\x15\xd9\x2c\xea\x07\x92\xb8\xa9\xde\x8e\x1c\x3d\x57\xe3\x5d\x14\xd3\xc9\xa1\xdc\x3e\x38\xe5\xbb\xb1\x99\xb9\x28\x7b\x34\x10\x76\xb8\x97\x0a\x41\x3f\x73\x53\x1f\xec\xe7\xd1\xe2\x01\x9b\xca\x7e\xb5\xaf\x53\x96\xcb\x8f\x9b\x1d\x5a\xb0\x05\x62\x03\x43\x3f\x2e\xd4\x02\x96\xd1\xbc\x48\x33\x7a\x89\x53\xa8\x8d\x54\x5c\xb4\xd3\x0a\x9e\xd6\x23\x43\x2d\x07\x55\xe3\x65\x0f\x17\x65\x7e\x5a\xdc\xf9\x02\xbe\x7f\xa0\xcf\x3b\x33\x9f\xe3\x53\xeb\x77\xd8\x49\xd1\x6a\x3b\xa1\x57\x86\x54\xd8\xef\x18\xd9\xe3\xe4\x1b\x8b\xa5\xec\x51\xed\xa5\x3a\xba\x8d\xb2\x76\x3e\xb9\xfd\x7c\x33\xa6\x69\xe4\x2f\xa2\x3e\x28\x29\xf8\x77\xdb\x27\x8d\x4a\x03\x53\x08\x3f\xbb\xdb\x10\xa7\xf7\x51\xf2\x16\x53\x2e\x4f\xa2\xfd\x92\xb9\x71\xaf\xc3\x28\x77\xf7\xbc\x38\x7f\xe7\xf6\x4a\x1e\xf8\x8e\x1b\x0d\xf6\x8d\xcb\xb1\x57\xf2\x08\x9d\x6c\x5b\x3b\x25\x2e\xbe\x5c\x73\x51\xbd\x55\xcd\xb5\x17\x44\xf9\xe8\x02\x6f\xad\xc1\x7e\x21\x73\x0d\x86\xe9\xa7\xb7\x36\x60\xad\xd4\xfb\x2f\x00\x00\xff\xff\xee\xf2\x8c\xa8\xd4\x0b\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 3028, mode: os.FileMode(420), modTime: time.Unix(1601366924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4f\x57\xda\xd8\x1b\xde\xe7\x53\x70\x60\xfb\x3b\xbf\x33\xeb\xee\xae\xc9\x55\x33\x0d\x49\x4e\x12\x3a\xe3\x6c\x72\x2c\xe6\xcc\x38\x55\xf0\xa0\x6e\x66\x55\x6a\x55\xb0\xa4\x30\x2d\x76\xb4\x65\xaa\x58\xa8\x9c\xb6\x80\xad\xad\x50\x22\xf6\xcb\xe4\xde\x84\x95\x5f\x61\xce\xcd\x0d\x31\x40\x99\xd1\x65\x7c\xff\x3c\xf7\x7d\x9e\xf7\x79\x89\x25\xd3\xab\xab\xe9\x14\x23\x82\x38\xd4\xe1\xcf\xbc\xaa\xa9\x77\x22\x51\x54\x32\x9d\xd3\x33\xd4\xf9\x84\x9a\x07\xa8\xd2\x88\x32\xbc\xac\x8b\x92\x76\x13\xe0\xb6\x3b\xa8\xd2\x70\xce\x2d\xc7\x3a\x72\x5b\x57\x4e\xbf\x35\xa8\x7e\x19\xfc\x7d\x82\xaa\xef\xd1\xce\x21\xfd\xce\xcb\x51\x86\x89\x25\x57\x36\xd7\x37\x8c\x0c\xc3\x0a\x09\x55\x83\x8a\xce\x41\x01\x6a\x50\x9f\x05\xbc\x00\xb9\x3b\x91\x28\xfe\xeb\x18\x9f\xef\xa3\xdc\xf1\xe0\xb0\x86\xfa\xcf\x51\xde\x74\xf6\x2e\xf0\xc3\xac\xf3\xf2\xf1\xe0\xd5\x8e\x73\x55\x8b\x06\xa9\xbc\xea\x81\x50\x12\xa2\xc8\x8b\x73\x77\x22\x51\x1a\x60\x77\x4d\x54\x69\xb8\xdf\x4a\x6e\xb5\x60\x77\x9b\xd7\x97\xd9\x89\x14\x41\x62\x81\x40\xde\xd5\xbe\x44\xdb\x75\x9a\xe6\x37\x36\x77\x9d\xde\xa9\x07\x34\x63\x2c\x19\xa9\x8d\xe5\xc5\x15\x66\x04\xa3\xae\x40\x55\x4a\x28\x2c\x24\xf9\x14\x66\xed\xa3\xfb\xb9\x7e\x7d\x99\x75\xdb\x75\xe7\xf4\x60\xf0\xac\x6e\x77\x9f\xe0\x4a\x1e\x6d\x9f\xbb\xd9\xb2\xdd\xb5\x70\xa5\x17\x9d\x52\x44\xff\x45\x12\x6f\x5b\x09\x15\xdb\x4e\xb9\x81\x0a\x5e\xb1\x59\x90\x10\x34\x9d\x55\x20\x07\x45\x8d\x07\x82\xce\x02\xd1\x7b\x1b\xed\x43\xa6\x61\x1d\xb8\xad\x1a\xda\x6d\x62\xb3\x65\x77\x4d\x77\xab\x4f\x9b\x78\x03\x61\x62\xbf\xa5\xd7\x37\x98\x79\x49\xd5\x74\x20\x28\x10\x70\x0b\x37\x6c\x52\xc8\x21\xba\x7d\xec\x5e\x74\xc0\xd4\x24\xe4\x20\xcf\xb1\x8a\x14\xf2\x90\xb2\xc9\x02\xfa\xcc\x82\x2e\x2b\xd2\x8f\x90\xd5\x6e\x5b\xab\xfa\xd5\x79\xd5\xf2\xd0\xab\x0b\xaa\x06\xe3\xba\xaf\xc2\x59\x29\x21\x72\xbe\x08\xb7\x73\x54\x72\xb8\xf2\x01\x57\x7a\xbc\xec\x85\x8b\x12\x09\x05\xf7\x00\x2f\x80\x19\x81\xcc\x86\x97\x23\xee\x97\xc7\xb8\x57\x22\x83\xb9\x38\x27\xf3\xd8\x5c\x37\x32\x8c\x0c\x54\xf5\x27\x49\xe1\xbc\xba\x71\xa0\xb1\xf3\x9e\x4a\x76\x9c\xe3\xec\xa0\x7c\xe8\xb6\xdb\x51\x46\x52\xf8\x39\x5e\x04\xc2\x68\xc8\xd3\xa3\xd1\xa8\x84\x7a\xa3\x36\xc0\x6a\xfc\x3d\xd2\xd5\x29\x37\x70\xae\x83\x2b\xef\x50\x89\x30\xe9\x3d\xb4\xe3\x66\xcb\x64\x3f\x5a\x55\xa7\xb4\x83\xfe\x3c\xf0\x00\x7b\xd9\xe1\xa7\x11\x71\x36\x6b\x34\xdf\x8b\x00\x5c\x9c\x17\xa7\x71\x1e\x59\x5c\x5a\x5d\x4e\x45\x68\x38\xa5\xde\x3d\x79\x1f\x62\x3f\x8c\x4e\x81\x02\xd0\x20\x17\x22\xc3\x87\xf9\xa9\x1a\x28\x8f\x8e\x3e\xca\x08\x1c\x90\x83\xa6\x09\x99\x03\x5e\x53\xf2\x75\xa4\x99\xfd\xad\x85\xcb\x5f\x7d\x9d\x65\x8c\x5f\x97\xd3\xa9\xa1\x02\x14\x38\xc7\x4b\xe2\xad\x56\x09\x15\x7a\xe8\xe8\x28\xac\x80\xd0\x02\x30\xb1\x3f\xd2\x29\x63\x58\x95\x2c\xd1\xed\x6a\x0e\x2b\x8c\x08\x6b\xab\xe1\xf4\x3f\xb9\xad\x2a\xca\x3d\x1b\xb5\x8a\x84\x40\x5c\xc2\x7d\xda\x41\xc5\x17\xc4\x56\x72\x67\x54\xd0\x94\x39\xfc\xe6\x21\x3e\xaa\x0f\xb6\x4d\xa7\x4f\x65\xc9\xc7\xc1\x1c\x9c\x96\xb8\x5f\x41\x5b\xc5\x29\x89\x4c\x2c\xbd\x66\xa4\xd6\x37\x16\x93\x0f\x98\x39\xa8\x0d\xa7\x04\x15\x45\x52\x08\xfb\xf9\x2b\x5a\x84\x8e\x84\xc4\xaf\x65\xd2\xbf\x1b\xc9\x8d\xb8\xb1\x7a\xdf\xc8\x04\x7c\x02\x8e\x0b\xf8\xa3\x2f\xc4\x1d\x0b\xed\x1d\x87\x32\x02\xeb\xf0\xf9\x9e\xa6\x21\x4a\xf9\x84\x6f\x0c\xb3\xe6\x81\xaa\xfb\x93\x22\x29\x5e\x70\x78\xe5\xaf\x2f\xb3\xdf\xf1\x9c\x54\x7a\xc9\x60\x44\x89\x83\x81\xe7\xf8\xe6\xad\x6b\x40\xbd\x4b\x18\xdb\xbe\xb0\xad\x17\xee\xde\x23\xe7\xd1\x57\xbc\x7f\x36\xd8\x2d\xe2\xe7\xa6\xdd\xaf\xe0\xe6\x1b\x54\x69\xe0\xfc\xa9\x5b\x2d\xfc\x2f\xe2\xb6\x3b\x4e\x33\x8f\xae\xb6\x51\x6b\xcb\xb6\x3e\xd0\xcf\xa8\x55\xc0\xed\xfd\xff\x93\x36\xf7\x17\x93\x0f\x36\xd7\x40\x32\x99\xde\x4c\x6d\x30\x32\x50\x40\x5c\x87\x71\x59\x5b\x20\x1d\x8a\x8f\xf0\xfe\xd9\x70\xed\xc9\xc3\xd5\x84\x2c\x4b\x8a\xe6\xf9\x9e\x89\xcb\x6d\x5c\x20\xc7\xc6\xf9\x68\xa1\xd7\x4f\xa2\x0c\x3b\x0f\xd9\xbb\xa1\xf3\x74\x5c\x1d\xbc\x2b\x04\x92\x0a\x9c\x6d\x06\xb0\x77\x13\xb2\x0e\x58\x56\x4a\x88\xb7\xf5\x38\x54\xdb\xb5\xad\xbe\xfb\xf9\x2d\x2a\x76\xa6\x38\x1d\x13\x5b\x5b\x59\x4c\x8d\x1d\x8f\xff\x28\x1b\xd6\xf3\x64\xd9\xd0\x11\x9e\xf1\xe6\xc4\xf8\xd8\x67\x79\x01\x52\x4b\x18\xca\x97\xfa\xa9\x5f\x9f\x40\xc5\x2f\x76\x6d\xeb\x02\xed\x98\x28\xb7\x8b\xcd\x93\x30\xfe\x91\xc5\xa1\x15\x83\x8b\x4c\xd9\xa3\xd1\xdf\xb9\xc8\x0a\x54\x35\x49\x81\x63\xe1\x38\x7b\x82\x6a\xe6\x30\x3c\xd0\xaf\x62\xac\xa7\x37\x33\x49\x63\x72\x24\xa1\x67\xfc\x0b\xf8\x30\x69\x63\xe7\xe8\x46\xcf\x23\xc7\xe7\xe3\x6b\xbb\xf7\x74\x4c\xd5\xee\xb7\x57\x6e\xb5\x80\x9b\x35\x2a\x4f\x1f\xe2\xca\xd2\xe2\x1a\xf5\xc8\xc0\x60\xe9\x24\x51\xa9\x80\x9b\x75\x94\x3b\x23\xc2\xf2\x8d\xb2\xf7\x83\xef\xa7\x82\x34\xc7\x8b\xe3\x19\x81\x9d\xd2\x13\xec\x35\xf0\xa2\x39\x5e\xf5\x6f\x18\xfd\x55\xe5\xbc\xcd\xe2\xe6\x09\xf9\x17\x4d\x71\x0e\xad\xc1\xe1\x4e\x64\xca\x61\x61\x62\xc9\xe5\x75\x86\xe5\x55\x6f\xe1\xc6\xb7\x90\x20\x1d\x1e\x7d\x9c\x7f\x8f\x8b\x45\xbb\xdb\x74\x5e\x3e\xb6\x2d\x0b\xed\x55\xa3\x0c\xf9\xfb\x27\x00\x00\xff\xff\x0d\x33\x19\x76\x21\x0a\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 2593, mode: os.FileMode(420), modTime: time.Unix(1601366898, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
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
