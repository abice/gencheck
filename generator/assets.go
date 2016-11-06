// Code generated by go-bindata.
// sources:
// template/bcp47.tmpl
// template/dive.tmpl
// template/hex.tmpl
// template/len.tmpl
// template/main.tmpl
// template/max.tmpl
// template/min.tmpl
// template/notnil.tmpl
// template/required.tmpl
// template/uuid.tmpl
// DO NOT EDIT!

package generator

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

var _templateBcp47Tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\x41\x6b\xdc\x30\x14\x84\xef\xfe\x15\x53\x83\x9a\xf5\xb2\xb1\x2f\x81\x40\x4a\x2e\x2d\x2d\xf4\x52\x72\x28\xbd\x2c\x0b\x51\xac\x27\xef\xa3\x5a\xc9\x95\x44\xda\x20\xf4\xdf\x8b\xe4\x6c\xbb\x29\x29\x7b\xb2\x79\x48\x33\xdf\xe8\x4d\x4a\xc3\xba\x01\xde\x7f\xb8\xbb\xba\xfe\x26\x0d\x2b\x19\x9d\xc7\x44\x96\xbc\x8c\x14\x30\x3a\x45\x88\x7b\x19\xf1\x93\x8d\xc1\x23\x79\xd6\x4f\x60\x0d\x09\xcd\x64\x14\x38\x40\x2e\xd7\x31\xba\xc3\x2c\x23\x3f\x18\x42\x88\x9e\xed\xd4\x00\xfb\x18\xe7\x70\x33\x0c\xd1\x39\x13\x7a\xa6\xa8\x7b\xe7\xa7\x61\x1f\x0f\x66\x78\x18\xe7\xab\xeb\x66\x3d\xe4\xdc\xa4\xa4\x48\xb3\x25\xb4\x75\xd8\xd6\x51\xb1\xa1\x1f\xe8\x3f\x15\xa3\xaf\x4f\x33\xa1\x5d\x74\x5b\xb4\xeb\xe3\x5f\xce\x4d\x39\xe6\x3d\x6e\x6e\x0b\xf7\xb8\xa7\xf1\x7b\xff\x39\x54\xa4\x55\x4a\xac\x61\x5d\xc4\x8a\xc3\x5d\xf4\xe8\xd1\xe5\xfc\x36\x25\xb2\x2a\xe7\xd0\xa7\xb4\x88\x7f\x91\x07\xca\xb9\x7b\x57\x75\xde\xdc\xc2\xb2\x41\x6a\x80\x94\x20\x95\xfa\xe8\xbd\x2b\x57\x5b\xf2\xbe\x1a\x56\x38\x32\x81\x5e\x21\xdc\xee\xfe\x30\x6e\x77\x47\xca\x9c\x1b\xed\x3c\xd8\x2a\xfa\xb5\x41\x60\x3b\x19\x7a\xe9\x5d\xe8\xbd\xb4\x13\xe1\x5f\xaa\x0a\x72\x2e\xe2\xff\x20\x4e\xc2\xbe\x62\x7a\x2e\xf0\x6a\xf6\x6c\xa3\xc6\xbd\x3e\xc4\xbe\x0e\xf5\xaa\x15\x22\x6c\x85\x50\x3b\x5c\x42\x88\xd0\x6e\x5e\xa6\x12\x61\x53\x34\xbb\x7b\xfc\x35\xea\xca\x9b\x01\xa7\xef\xb6\xec\xf7\xb9\x66\xec\xec\xe2\x78\xf4\x5b\x4a\x50\xaa\x55\x56\xf7\x58\x6a\x09\x67\x9f\x0b\x77\x21\x02\x44\xb8\x68\x4f\x0c\x4e\xb2\x77\x8b\x74\x8d\xbc\x7c\x71\x99\x73\xf3\x3b\x00\x00\xff\xff\x2e\xc7\xf3\xb6\xe7\x02\x00\x00")

func templateBcp47TmplBytes() ([]byte, error) {
	return bindataRead(
		_templateBcp47Tmpl,
		"template/bcp47.tmpl",
	)
}

func templateBcp47Tmpl() (*asset, error) {
	bytes, err := templateBcp47TmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/bcp47.tmpl", size: 743, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateDiveTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x4f\x4b\x03\x41\x0c\xc5\xef\xf3\x29\x9e\x0b\xa5\xdd\x42\xb7\x77\xa5\x37\xf5\x28\x82\xe2\x7d\xd8\xc9\xd4\xe0\x38\x5b\x32\x63\x45\x42\xbe\xbb\xec\x6e\xff\x80\xde\x92\xf0\xf2\xde\x8f\xa7\xba\x5d\x3b\x20\xf0\x91\xf0\xcd\x29\xa1\xf7\x29\x61\x4f\xb9\x7f\xa7\xfe\x03\x47\x9f\x38\xf8\x4a\x18\x32\x3c\x4a\x95\xaf\xbe\x3a\x60\xbd\x35\x73\xaa\x81\x22\x67\x42\x33\x7e\x37\xd3\x05\x1c\x31\x08\x56\x5c\x5e\x26\x2d\xba\xf6\xba\x3c\x57\x41\xd7\x5e\x74\x5c\xa6\x03\xcc\x38\xa2\x74\xaa\xdd\x23\x53\x0a\x4f\xfe\x93\xcc\x70\xb3\x43\xe6\x04\x85\xea\x06\x94\x83\x99\xe3\x08\x12\xc1\xed\xee\xc2\xd7\xbd\x9d\xf8\x56\x7f\xff\xdb\xbb\x49\x7b\x76\x71\x80\x2a\x7c\x08\x0f\x22\xc3\x18\xda\x90\x48\x03\x33\xf7\x9f\xc6\x54\xe7\x3c\x55\x50\x2a\x84\x79\xdc\x53\x26\xf1\x95\x87\x3c\x7b\xac\x0e\xc2\xb9\x46\x34\xf7\x63\x77\x5c\x90\x87\x3a\xf7\x35\x96\x15\x47\x14\x2c\x17\x05\x8b\xb2\x6c\x70\x45\x3b\x8d\xaf\x3f\x07\x6a\x67\xe7\x73\x18\xe5\x80\x8d\x99\xfb\x0d\x00\x00\xff\xff\xde\x64\x88\x1d\x93\x01\x00\x00")

func templateDiveTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDiveTmpl,
		"template/dive.tmpl",
	)
}

func templateDiveTmpl() (*asset, error) {
	bytes, err := templateDiveTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/dive.tmpl", size: 403, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateHexTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\xc1\x4e\xeb\x30\x10\x45\xf7\xf9\x8a\xfb\x2c\xf9\x35\x41\x6d\x3e\x00\xd4\x25\x08\x36\x88\x05\xbb\xaa\x52\xa3\x7a\x9c\x5a\xa4\x4e\xb1\x23\x54\x34\x9a\x7f\x47\x76\x54\x48\x51\x81\x55\x22\xcb\xbe\xf7\x9c\x99\x82\xd9\x90\x75\x9e\xa0\x76\x74\x54\x22\x05\x33\x9c\x05\xbd\xa2\xbe\x73\xd4\x99\xe7\xf7\x03\x41\xc5\x21\x38\xdf\x2a\xa8\xab\xd3\x9f\x48\x91\xae\x85\x80\xeb\x25\x5a\xf2\xdb\x1d\x6d\x5f\xea\x87\x78\x4f\xc7\x92\xd9\x59\xf8\x7e\x40\xe9\xe2\xd3\x10\x50\xa3\x12\xf9\xcf\x4c\xde\x88\xc4\x9a\x79\x8c\x7e\x6c\xf6\x24\x52\xdd\xe4\x94\x7f\x4b\x78\xd7\x81\x0b\x80\x19\x8d\x31\xb7\x21\xf4\xe9\xa9\xa2\x10\x72\x5d\x46\xa3\x2e\xd2\x05\xbe\xd5\xfa\x93\x70\xb5\x3e\x31\x8a\x14\xb6\x0f\x70\xde\xd0\x71\x8e\xe8\x7c\xdb\xd1\x79\x77\x62\x0f\x8d\x6f\x09\xdf\xa9\x32\xc8\xef\x82\x3f\x21\x4c\x54\x2f\x54\xfe\xa5\x5b\x1e\x82\xf3\x83\xc5\xc6\xee\x87\x3a\x1f\xda\x52\x69\x1d\x57\x5a\x9b\x35\x16\xd0\x3a\xaa\xf9\xb9\x93\x8e\xf3\x94\x59\x6d\xf0\x55\x54\xa5\x89\x01\xd3\xa9\x8d\xbb\x6d\xc9\x53\x68\x06\xd7\xfb\xb1\xf1\xd4\x97\xd6\x0f\x17\xf3\xda\xde\x9a\xce\x19\xf4\x1e\x36\xc5\x61\xa6\x23\x74\x9c\xa9\x49\xfc\xc4\xbc\x1a\x83\xb3\xf0\xf8\xc5\x42\xa4\xf8\x08\x00\x00\xff\xff\x1b\x7e\x02\xe1\x59\x02\x00\x00")

func templateHexTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateHexTmpl,
		"template/hex.tmpl",
	)
}

func templateHexTmpl() (*asset, error) {
	bytes, err := templateHexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/hex.tmpl", size: 601, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateLenTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x53\xc1\x6a\xdc\x30\x10\xbd\xef\x57\xbc\x0a\x42\xbc\x81\x38\x24\x2d\x21\x97\x1c\x72\x68\x6f\x0d\x39\xb4\xb7\x5c\xb4\xd6\xd8\x16\xd5\x8e\x5c\x69\x94\xc4\x08\xff\x7b\x91\xbc\x5b\x96\x50\x7a\xe8\xa1\x17\xcf\x13\x7a\x3c\xbf\x79\x33\xda\xe4\x6c\xa8\xb7\x4c\x50\x8e\x58\x2d\xcb\x26\xe7\xab\x0b\x7c\x8f\x04\x19\x09\xbb\x64\x9d\xc0\x32\x1c\x31\xfa\xc4\x9d\x58\xcf\xe8\x7d\xa8\xb7\x32\x4f\x14\x21\xa3\x16\xc4\x34\x4d\x3e\x08\xac\xb4\xb8\xb8\xaa\x32\xb0\x3d\x7c\x40\x03\x1b\xbf\xea\x09\x2d\xb6\x15\x3f\x84\xa0\xe7\xf5\x44\x3f\xd1\x7e\xb1\xe4\xcc\xb7\x79\x22\xa8\x28\xc1\xf2\xa0\xb6\x68\x46\x1d\x9f\x02\xf5\xf6\x0d\xaa\x1b\x35\xab\x13\xda\x76\x59\x36\x28\xd2\x1f\x1a\x47\xdc\xc4\x36\xe7\xf5\xf2\x51\xef\x69\x59\xb6\xb8\xbf\x47\xce\xed\x93\x0e\x7a\x5f\x8e\x79\x03\x00\x39\x43\x1b\xf3\x39\x04\x1f\xd0\x42\x51\x01\xb1\x7d\xa4\xd7\xe6\xb9\xf4\x3d\xc8\x88\xbd\x8d\x7b\x2d\xdd\xf8\xac\xb6\x0a\xf5\x27\x87\x2c\x1e\x8c\xb1\x3c\xa0\x1b\xa9\xfb\x11\x6b\xef\x96\x85\x06\x0a\x78\xd1\x2e\x95\x00\x3c\x3a\x6f\xb9\xb3\x86\xf0\x6a\x65\xac\xd9\x4c\x81\x3a\x32\xc4\x82\x48\x82\xdd\x8c\xc1\x5f\x4e\x4e\xcf\x43\xf0\x89\xcd\xd5\x8b\x76\xd6\x68\xf1\xe1\x77\x5a\xe4\x22\x95\xbe\xde\x85\x62\x59\x54\xfd\xde\xad\xe5\xfa\x76\xad\x1f\x6f\xd6\x7a\xfb\x49\x41\xa5\x95\x95\x0e\xb4\x74\xe4\xa5\x23\x31\x1d\x99\xbb\x59\x48\x41\x85\xc4\x54\x86\x5d\x73\x7c\x9f\xe1\x9f\x22\xfc\x87\x00\xff\xd6\x55\xef\xbc\x2e\x8e\x6b\xad\x0e\x2b\xaa\x16\x3b\xbf\x9f\x1c\xbd\x9d\xe2\xeb\x9b\xbb\x2a\xf9\x7f\xec\xae\xf3\x18\x88\x29\xe8\xb2\xef\xab\x4c\x33\x05\xcb\xd2\xd7\x77\x02\x1b\xc1\x5e\x50\xa7\x88\xf2\x22\x8a\x1d\x9c\x9f\x45\x9c\xc5\xf3\xe3\xba\x16\x7b\xa7\x9b\x7b\x10\x26\x36\x2b\x2a\xe0\x72\x59\x36\xbf\x02\x00\x00\xff\xff\x4e\x8b\xbc\x81\x84\x03\x00\x00")

func templateLenTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateLenTmpl,
		"template/len.tmpl",
	)
}

func templateLenTmpl() (*asset, error) {
	bytes, err := templateLenTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/len.tmpl", size: 900, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\x4d\x8b\xdb\x30\x10\xbd\xfb\x57\x3c\x42\xa0\x76\x58\xdb\x39\x2f\x6c\x0f\x6d\xb3\x25\x87\x4d\xa1\x94\xde\x15\x7b\x6c\x8b\xc8\xb2\x91\xe5\xb4\x45\xe8\xbf\x17\x7d\x38\x0e\x61\xf7\x12\x86\xbc\x99\x37\x4f\x6f\x9e\x8d\xc9\x51\x53\xc3\x25\x61\xd3\x11\xab\x49\x6d\xac\x4d\x46\x56\x5d\x58\x4b\x30\xa6\x88\xa5\xb5\x49\xc2\xfb\x71\x50\x1a\x69\xb2\x69\xb9\xee\xe6\x73\x51\x0d\x7d\xc9\xce\xbc\xa2\xb2\x25\x59\x75\x54\x5d\x36\x49\x96\x18\x43\xb2\x46\xee\x26\xee\xd9\x27\xad\xe6\x4a\x3b\x76\x63\xb6\x6a\x16\xf4\xc6\x46\x3c\xbf\xa0\x70\xf5\xe4\xfb\x9b\x59\x56\x48\x27\x18\xc3\x1b\x14\xa3\x56\x6f\xa4\xbb\xa1\x86\xb5\x3b\x4f\x6a\xad\x31\x85\x64\x3d\x59\x9b\xe1\x37\x13\xbc\x66\x9a\x90\x66\x48\x49\xa9\x41\x65\x30\x09\x60\x4c\xb9\xc3\x0f\x29\xfe\xa1\x67\x17\x82\xee\x08\x1e\x9d\xd0\xb3\x11\xbc\xc1\x1f\xfa\xa4\x08\xed\xc0\x65\x0b\x3d\x60\x9e\x08\x5c\x17\xd8\x95\xd6\xfa\xf1\xdc\x35\xb5\x1a\xa9\x20\x19\xd5\x65\xd8\x7b\x81\x88\xd2\x5a\x31\x9c\x99\x78\x65\x5c\xbc\xb2\x49\x7b\xe0\xca\x14\xae\x87\xb0\x68\x71\xa3\x88\x1a\xf9\x20\x03\xe2\x09\x48\x4c\x14\x46\x62\xfb\xf3\x8b\x97\x9a\x7e\x38\xf6\x84\xfd\x13\x8c\x59\xf5\x58\x9b\x05\x2a\xe7\x89\xaf\xe0\x3c\x8f\x75\xb9\xc3\x51\x93\x72\xde\xe8\x4e\x0d\x73\xdb\x79\x17\xdc\xa4\xf7\x40\x0f\xa8\x14\x05\x38\xfc\x3d\xad\xaf\x57\x4c\xb6\x84\x2d\x97\x35\xfd\x7d\xc2\xb6\xe1\x24\xea\x87\x33\x01\x7e\xe3\xd6\x5d\xc2\x41\xa1\xa9\x38\xf9\xc3\xdc\xd0\xd0\xbf\xc2\x3f\x1f\xe7\xef\x5c\xde\x3e\xb8\x0c\x94\x25\xbe\x1c\xbe\x1f\x4f\xeb\x22\x6b\xb1\xfa\x32\x45\x92\x1c\x51\xaf\x3a\x46\xc1\xfe\x95\x6e\x6b\xd8\x1f\xe9\x3c\xe1\xa2\xca\x2b\xbd\x6d\x0a\x34\x5f\x99\x10\xbf\xa8\x1f\x85\xb3\x25\x90\x84\xdf\x9b\xde\x1c\x8b\xdb\x9e\xec\x70\xfa\xf6\xa1\xb6\xe5\x85\x77\x37\x59\xa7\xd7\x7c\xfa\x63\xbf\x1b\x50\x74\xec\x1a\x2f\x73\x17\xcc\xf7\x73\xe9\x41\xde\x40\x90\x4c\x63\xa2\x32\x7c\xde\xfb\x6f\x01\x50\xa4\x67\x25\x97\xa8\x25\xc0\xa3\x9c\xd8\x20\xb9\x48\x6c\xb2\x24\xea\x7f\x00\x00\x00\xff\xff\xd7\x72\x9c\x7a\x16\x04\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 1046, mode: os.FileMode(420), modTime: time.Unix(1478389953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMaxTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x52\xc1\x6a\xdc\x30\x10\xbd\xfb\x2b\x1e\x82\x10\x3b\x50\x87\xa4\x25\xe4\x54\x28\xa5\xbd\x35\xe4\xd0\xde\x72\x51\xd6\xe3\xb5\xa8\x3c\x76\x35\x52\xb3\x8b\xd0\xbf\x17\x49\xbb\xe9\x92\xd2\xd0\x43\x2f\x9e\x67\xcf\xe3\xf9\xcd\x9b\x69\x62\x1c\x68\x34\x4c\x50\xb3\xde\xa9\x94\x9a\x18\x2f\x2f\xf0\x4d\x08\x7e\x22\x3c\x06\x63\x3d\x0c\xc3\x12\x63\x0c\xbc\xf1\x66\x61\x8c\x8b\x2b\x5d\xbf\x5f\x49\xe0\x27\xed\x21\x61\x5d\x17\xe7\x61\x7c\x8f\x8b\xcb\x22\x03\x33\x62\x71\x68\x61\xe4\x8b\x5e\xd1\xa3\x2b\xf8\x83\x73\x7a\x5f\xdf\xe8\x07\xfa\xcf\x86\xec\xf0\x75\xbf\x12\x94\x78\x67\x78\xab\x3a\xb4\x93\x96\x7b\x47\xa3\xd9\x41\x6d\x26\xcd\xea\x84\xd6\xa5\xd4\x20\x4b\x5b\xe2\x56\xfa\x18\x6b\xeb\x4e\xcf\x94\x52\x87\xf7\x88\xb1\xbf\xd7\x4e\xcf\x29\x21\x36\x00\x10\x23\xf4\x30\x7c\x72\x6e\x71\xe8\xd1\xae\xce\xb0\x1f\xa1\x28\x7f\x90\xfe\x8e\x9e\xda\x07\x65\x89\xb7\x7e\xc2\x93\x16\xcc\x8b\xcb\xb3\x6b\xc6\x99\x3c\xa8\x4e\xa1\xca\x75\x28\x3f\x3e\xe4\xf3\x71\xa2\xcd\x77\x29\x41\x18\xf6\xb4\x25\x87\x9f\xda\x06\x92\xe7\xe1\xc9\x0a\x65\x9b\x2f\x66\x34\xec\x55\x79\xde\xd6\x72\x75\x53\xeb\xdb\xeb\x5a\x6f\xde\x29\xa8\x50\x59\xe1\x40\x0b\x47\x5e\x38\x12\xc3\x91\xf9\xb8\xf7\xa4\xa0\x5c\x60\xca\xbb\x33\x23\x5e\x46\xf2\x67\x22\xff\x92\x47\x0e\xc2\x92\xc8\xdf\x82\x78\x6d\xc4\xd1\x2e\x3a\xdb\x2f\xb5\xd8\x2d\xa8\xf8\xdd\x2c\xf3\x6a\x69\x77\x8a\xaf\xae\x6f\x15\xfe\xb7\xf7\xd7\x96\xf8\xdb\x7b\xdd\xd4\x96\x98\x9c\xce\x87\x5d\x75\x9f\x55\x67\xc3\x30\x02\x5e\x7c\x5e\xae\x19\x90\x4f\x3f\xbb\xc3\xf9\x99\xe0\x4c\xce\x8f\x77\x99\xdd\x9e\x9e\xe8\x41\x98\x78\xa8\x28\x83\x37\x29\x35\xbf\x02\x00\x00\xff\xff\xfc\xf1\x82\xfe\x6d\x03\x00\x00")

func templateMaxTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMaxTmpl,
		"template/max.tmpl",
	)
}

func templateMaxTmpl() (*asset, error) {
	bytes, err := templateMaxTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/max.tmpl", size: 877, mode: os.FileMode(420), modTime: time.Unix(1478386586, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMinTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x93\x4f\x6b\xdc\x30\x10\xc5\xef\xfb\x29\x1e\x82\x90\x75\x20\x0e\x49\x4b\xc8\xa1\x97\x1c\xda\x5b\x43\x0e\xed\x2d\x17\xc5\x1a\xdb\x43\xb5\x23\x57\x7f\xb2\x31\xc2\xdf\xbd\x48\xde\x0d\x4b\x0a\x6d\x6f\xbd\x78\xc6\xd6\xe3\xf9\xe9\xa7\xd1\x26\x67\x43\x3d\x0b\x41\xed\x58\xd4\xb2\x6c\x72\xbe\xba\xc0\xf7\x40\x88\x23\xe1\x39\xb1\x8d\x60\x81\x25\x41\x9f\xa4\x8b\xec\x04\xbd\xf3\x75\x35\xce\x13\x05\xc4\x51\x47\x84\x34\x4d\xce\x47\x70\x6c\x71\x71\x55\x6d\xc0\x3d\x9c\xc7\x16\x1c\xbe\xea\x09\x2d\x9a\xda\xdf\x7b\xaf\xe7\xf5\x8d\x7e\xa2\xfd\xc2\x64\xcd\xb7\x79\x22\xa8\x10\x3d\xcb\xa0\x1a\x6c\x47\x1d\x1e\x3d\xf5\xfc\x0a\xd5\x8d\x5a\xd4\x89\xac\x59\x96\x0d\x8a\xb5\x25\xd9\x86\x36\xe7\x75\xe9\x41\xef\x68\x59\x1a\x7c\x42\xce\xed\xa3\xf6\x7a\xb7\x2c\xc8\x1b\x00\xc8\x19\xda\x98\xcf\xde\x3b\x8f\x16\xdb\xc9\xb3\xc4\x1e\x8a\xca\x87\xd0\x3e\xd0\x7e\xfb\xa4\x2c\xc9\x10\x47\xec\x75\x80\xa5\x50\xf7\x24\x38\x0b\x4f\xaa\x51\x58\xed\x1a\xd4\x1f\x1f\xf8\xdc\x1b\xc3\x32\xa0\x1b\xa9\xfb\x11\x2a\x0f\x96\x48\x03\x79\xbc\x68\x9b\x0a\x14\x87\xce\xb1\x74\x6c\x08\x7b\x8e\x63\xe5\x35\x79\xea\xc8\x90\x44\x04\x8a\x78\x9e\x31\xb8\xcb\xc9\xea\x79\xf0\x2e\x89\xb9\x7a\xd1\x96\x8d\x8e\xce\xbf\x11\x24\x1b\xa8\xec\xf5\x1d\x28\x96\xa8\xea\xf3\x6e\x2d\xd7\xb7\x6b\xfd\x70\xb3\xd6\xdb\x8f\x0a\x2a\xad\xaa\x74\x90\xa5\xa3\x2e\x1d\x85\xe9\xa8\x7c\x9e\x23\x29\x28\x9f\x84\xca\x00\x70\x8f\xf7\x5c\x7f\xc7\xfa\x2f\x50\xff\x46\xf3\x4f\x5b\xec\xad\xd3\x25\x7e\xad\x35\x6e\xed\x6a\xde\xce\xed\x26\x4b\xaf\xa7\xfd\xf5\xcd\x9d\xc2\xff\xc9\xbe\x9e\xd4\x40\x42\x5e\x97\xdb\xb1\xfa\xbe\xb9\xee\x58\xc0\x01\xe2\x22\xea\xf9\xa2\xdc\x9f\x92\x0e\xe7\x67\x01\x67\xe1\xfc\x38\xdc\x25\xed\xe9\x9c\x1f\x8c\x49\xcc\xda\x95\xe6\x72\x59\x36\xbf\x02\x00\x00\xff\xff\x42\x2f\x25\xf9\xb2\x03\x00\x00")

func templateMinTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMinTmpl,
		"template/min.tmpl",
	)
}

func templateMinTmpl() (*asset, error) {
	bytes, err := templateMinTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/min.tmpl", size: 946, mode: os.FileMode(420), modTime: time.Unix(1478386610, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateNotnilTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xb1\x6a\x43\x31\x0c\x45\xf7\xf7\x15\x17\x43\x48\x32\xd4\x7f\x90\xb1\x1d\xdf\xd4\x31\x8b\x8b\xe5\x22\x50\xe5\x60\xb9\x2d\x45\xf8\xdf\x8b\x93\xf7\xda\xc1\x70\x30\xe8\x9e\xe3\x9e\xa9\xb0\x12\x82\xd6\xae\x2c\x61\x8c\x05\x70\x07\x17\xb0\xad\x9f\x22\xe9\x4d\x08\x11\xf7\x7f\x2e\xb0\xe8\x1e\x5f\x98\x24\xaf\xe9\x83\xc6\xc0\xe5\x02\x65\x81\x2f\xc0\xfd\x30\xe5\xfc\xdc\x5a\x6d\x88\x08\x34\xc1\xe2\x4a\xdf\xa7\x6b\x60\xc3\xca\x72\x0d\xe7\xf0\x18\xdb\x44\x24\x46\xd8\xad\xef\xa4\xd4\x52\xe7\xaa\x8f\x8d\xd3\xad\xb1\xf6\xb2\xd7\x81\x0d\x5a\x3b\xbe\x92\x70\x46\x55\xe8\x7c\x7b\x64\x99\x55\x38\x1e\x0c\x07\x3b\x06\xfc\x57\x6e\xf8\xfa\x73\xa3\xf3\x9f\x8a\x34\x4f\x76\x9f\xf0\x34\xc6\xf2\x1b\x00\x00\xff\xff\x8d\xc2\xfb\xfb\x0b\x01\x00\x00")

func templateNotnilTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateNotnilTmpl,
		"template/notnil.tmpl",
	)
}

func templateNotnilTmpl() (*asset, error) {
	bytes, err := templateNotnilTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/notnil.tmpl", size: 267, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateRequiredTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x92\x31\x6b\x23\x31\x10\x85\x7b\xff\x8a\x61\xaa\xbb\xe2\x16\xec\x3b\x8c\x1b\x97\x97\xd2\x45\x92\xd2\x8d\x6c\xcd\x9a\x01\x59\x5a\x8f\x24\x27\x8e\xd8\xff\x1e\x24\xed\x9a\x10\x12\xc7\x24\x04\xd2\xec\x8c\x96\xa7\xf7\x3e\xc4\x4b\x49\x53\xcb\x96\x00\x85\x0e\x91\x85\x34\xf6\xfd\x04\x20\x25\xe0\x16\xd8\xaf\xa2\x31\x6a\x63\x08\x1a\x28\xff\xb9\x05\xdf\xa4\xd4\xdc\x30\x19\xbd\x52\x7b\xea\x7b\x58\x2e\xc1\xb2\x81\x34\x01\x28\x17\x95\xd6\xff\x45\x9c\x40\x03\x48\x79\xf1\xcd\x8a\x1e\x7e\xad\x91\x3d\x8c\x29\x6b\xfc\x8d\xd5\x71\x48\x23\xe3\x29\xbb\xd3\x01\xaa\xf9\xfd\xa9\x23\x40\x1f\x84\xed\x0e\x2f\x86\x23\x7e\x4f\x36\xdb\x80\xe5\xbb\xa8\x63\x3a\xaf\xf3\xef\xac\xce\xf9\x3f\x04\x8c\x55\x15\x07\x59\x1c\x75\x71\x14\xc6\x51\xb9\x39\x05\x42\x40\x89\x96\xce\x6f\xbc\x23\x4b\xa2\x02\x3b\x5b\xa9\xf1\x76\x60\x04\x67\x81\x6d\xa0\x1d\x09\x1c\x95\x89\xe4\x81\x3d\x58\x17\xc0\xc7\xae\x73\x12\x48\x37\x67\x97\x77\xf8\x5b\xe3\x54\x66\x2b\xb3\xb0\x94\xad\xc0\x6c\xdd\xbe\x33\xf4\xf8\x72\x9f\xce\x16\xc3\xb3\x7c\x04\x66\xe3\x9e\x84\xb7\xca\x7c\x1a\x6d\xe3\x9c\xb9\x32\x2c\x4b\x49\xd9\xeb\xa3\xd8\xdf\x05\x89\xdb\x30\x56\xf6\xa8\x04\x9e\x48\xdc\xab\xe2\x8c\xc7\x0c\x74\xa9\x5d\x6f\x5d\xfd\x7a\xdb\x7e\x26\x99\xd5\xf9\x9c\x52\x5e\xfe\xf4\xfd\xe4\x39\x00\x00\xff\xff\x08\xdc\x4c\x18\x1e\x04\x00\x00")

func templateRequiredTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateRequiredTmpl,
		"template/required.tmpl",
	)
}

func templateRequiredTmpl() (*asset, error) {
	bytes, err := templateRequiredTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/required.tmpl", size: 1054, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateUuidTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x92\xd1\x4a\xc3\x30\x14\x86\xef\xfb\x14\xbf\x81\xb8\x56\xb6\x3e\x80\xb2\x3b\x15\x76\x23\x5e\xb8\xab\x31\x58\x59\x4e\xea\xc1\x2e\x9d\x49\x27\x4a\x38\xef\x2e\x4d\x99\x76\x32\xd9\x55\x4b\x69\xfe\xff\xfb\xce\x49\x16\xa3\x21\xcb\x8e\xa0\x0e\x07\x36\x4a\x24\x8b\x11\x6c\x41\xef\x28\x1f\x99\x1a\xf3\xf2\xb5\x27\xa8\xd0\x79\x76\xb5\x82\xba\x39\xbe\x89\x64\xfd\x6f\xde\xe3\x76\x8e\x9a\xdc\xf6\x95\xb6\x6f\xe5\x22\x2c\x97\x8b\xfb\x3c\x46\xb6\x70\x6d\x87\x9c\xc3\x73\xe7\x51\xa2\x10\xb9\x8e\x91\x9c\x11\x09\x65\x8c\x43\xf6\x53\xb5\x23\x91\xe2\x2e\xc5\x5c\xcd\xe1\xb8\x41\xcc\x80\x18\x51\x19\xf3\xe0\x7d\xdb\x1f\x55\xe4\x7d\xea\x4b\x6c\xd4\x04\x3a\x03\xb8\x5a\xff\x20\xae\xd6\x47\x48\x91\xcc\xb6\x1e\xec\x0c\x7d\x4e\x11\xd8\xd5\x0d\x9d\x76\xf7\xf0\xbe\x72\x35\xe1\x2f\x55\x02\xb9\x60\xf8\x1f\xc3\xc8\xf5\x4c\xe7\x25\xdf\x7c\xef\xd9\x75\x16\x1b\xbb\xeb\xca\xf4\xd1\xe6\x4a\xeb\xb0\xd2\xda\xac\x31\x83\xd6\x41\x4d\x4f\xa5\x74\x98\xf6\x99\xc5\x06\xbf\x45\x45\x3f\x32\x60\x3c\xb6\x61\xbb\x35\x39\xf2\x55\xc7\xad\x1b\x1a\x8f\x7d\xe9\x06\x80\x43\x5a\xdc\x47\xd5\xb0\x41\xeb\x60\xfb\x3c\x4c\x74\x80\x0e\x13\x35\xca\x1f\xa9\x17\x43\x72\x32\x1e\x9e\x98\x89\x64\xdf\x01\x00\x00\xff\xff\xb3\x04\xf3\xf9\x5d\x02\x00\x00")

func templateUuidTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateUuidTmpl,
		"template/uuid.tmpl",
	)
}

func templateUuidTmpl() (*asset, error) {
	bytes, err := templateUuidTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/uuid.tmpl", size: 605, mode: os.FileMode(420), modTime: time.Unix(1478296916, 0)}
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
	"template/bcp47.tmpl": templateBcp47Tmpl,
	"template/dive.tmpl": templateDiveTmpl,
	"template/hex.tmpl": templateHexTmpl,
	"template/len.tmpl": templateLenTmpl,
	"template/main.tmpl": templateMainTmpl,
	"template/max.tmpl": templateMaxTmpl,
	"template/min.tmpl": templateMinTmpl,
	"template/notnil.tmpl": templateNotnilTmpl,
	"template/required.tmpl": templateRequiredTmpl,
	"template/uuid.tmpl": templateUuidTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"bcp47.tmpl": &bintree{templateBcp47Tmpl, map[string]*bintree{}},
		"dive.tmpl": &bintree{templateDiveTmpl, map[string]*bintree{}},
		"hex.tmpl": &bintree{templateHexTmpl, map[string]*bintree{}},
		"len.tmpl": &bintree{templateLenTmpl, map[string]*bintree{}},
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
		"max.tmpl": &bintree{templateMaxTmpl, map[string]*bintree{}},
		"min.tmpl": &bintree{templateMinTmpl, map[string]*bintree{}},
		"notnil.tmpl": &bintree{templateNotnilTmpl, map[string]*bintree{}},
		"required.tmpl": &bintree{templateRequiredTmpl, map[string]*bintree{}},
		"uuid.tmpl": &bintree{templateUuidTmpl, map[string]*bintree{}},
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

