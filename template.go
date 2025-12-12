// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// tmpl/swagger.yaml
package main

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

// Mode return file modify time
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

var _tmplSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x51\x6f\xd3\x30\x10\x7e\xcf\xaf\x38\x55\x7d\x5c\xa2\x0a\xde\x22\xf1\xd0\x31\xb4\x4d\xdb\x50\xb5\x15\x84\x84\x10\xf2\x92\x6b\x63\xad\xb1\x3d\x9f\x33\x88\xba\xfe\x77\x64\xc7\x49\x9c\xb4\x83\x0e\x1e\x78\x6b\xbe\x7c\xbe\xfb\xee\xbb\xcf\xa9\x54\x28\x98\xe2\x29\xbc\x4d\x66\xc9\x2c\xe2\x62\x25\xd3\x08\x20\x47\xca\x34\x57\x86\x4b\x91\xc2\x64\xbb\x85\xe4\xbd\x14\x2b\xbe\x4e\xce\xfa\x17\xb0\xdb\x4d\x22\x80\x27\xd4\xb4\x47\xfb\xdc\x80\x9e\x62\xb8\xd9\xe0\x90\xb0\xb4\x90\x7b\xbd\xdd\xc6\xc0\x57\x90\x5c\xd2\x29\x32\x8d\x1a\x76\xbb\x28\x93\xa5\x92\x02\x85\x21\x2b\x86\x30\xab\x34\x37\xf5\x5d\x56\x60\x89\x0e\x02\xb8\x77\xe4\x79\x65\x8a\xe6\x19\xc0\xd4\x0a\x53\x28\x8c\x51\x1e\x20\xc7\x4f\x3d\x35\x6a\xcb\x58\x7e\x1c\x9e\x87\xaf\xdf\x9c\x0a\x14\xf9\x6e\x17\xea\x61\xc4\xb3\x63\xe5\x58\xee\x71\x6a\x2c\x73\x2c\xa6\x3b\xfd\x92\x96\xb9\xe2\x57\x58\x1f\x29\x86\x39\xf2\xbe\x9a\x06\xf7\x10\x17\x29\x14\xc8\x72\xd4\x1e\x10\xcc\xca\xfb\x12\xcf\x17\x97\xf1\x15\xb2\x7a\xa4\x31\x28\x3a\x14\x49\xa8\x6d\x06\x1a\x56\xa5\x37\xc3\x45\x9f\x32\xc2\x05\x33\x85\xdb\xb5\x61\x6b\x4a\xa3\xed\x16\x34\x13\x6b\x84\xe9\xf7\x13\x98\x22\xa4\xef\x20\xb9\x93\xda\x60\xfe\x41\x18\x6e\x38\x92\x9d\x33\xf6\x82\x6c\xb1\x29\x26\xe7\x5a\x56\xca\xe7\x69\x98\x4e\x1b\x21\x2b\xc5\x1e\x52\xcc\x14\xae\x43\x7c\x4c\x8b\x80\x66\x0f\x9e\xc0\x54\x23\xc9\x4a\x67\x48\xf6\xc4\x14\x13\xc7\x76\xb6\x03\x58\x1d\x96\x06\xcf\x70\x8b\xa5\x7c\x72\x63\x2d\x34\xae\xf8\xcf\xb9\xc8\x6f\x51\x6d\x58\x86\x0b\xa6\x59\x69\x8b\x37\xce\x07\x1d\x1e\x1e\x82\xfa\xae\x7c\xdf\xcc\xd5\x6f\x3a\xb4\x60\x72\x83\xa6\x90\x39\x3c\xc3\x52\x5e\xcb\x1f\xee\x5e\x74\xdb\x74\x36\x36\xbf\xe3\x43\x06\xb9\xbc\x55\x65\xc9\x74\xed\x0d\xec\xca\x7e\x64\x25\x06\x2c\x9f\x30\xa9\x03\x8a\x9b\x01\x0d\x6a\x0a\xc0\x53\x99\xd7\x89\x9f\xae\x47\x2f\x5c\x80\xf6\x78\x37\xbc\xc4\x65\xad\xb0\x1d\xec\xc0\x07\xa5\x3f\x11\x7e\x52\x9e\xe1\x93\x40\xca\x98\xc2\x8b\xe5\xcd\x75\xa0\x53\x75\x9a\xd2\x40\x79\xb0\x63\xa5\x9c\xa7\x7e\x41\xe7\x68\xdc\x76\xda\x6d\x74\x66\x05\x99\x52\xa1\x5b\xcd\x7d\x98\xd8\xe3\x3d\xe4\xae\x2c\x4b\xbb\xe7\xf6\x1e\x4d\xc8\x68\x2e\xd6\x3d\x71\x1c\xc8\x16\xd7\xf8\x58\x71\x8d\x79\x0a\x46\x57\x18\x08\xf7\x71\x1d\xec\xe0\x8f\x1e\x76\xea\xef\x65\x5e\x8f\x84\x0f\xa1\xbf\x14\x6e\xe5\x22\x19\x18\x16\x3b\x7e\x88\xf6\x2e\xd9\x75\x58\xe7\x07\x29\x0f\x62\x35\x9e\x5c\x48\xe3\x4f\x24\x67\x9c\xd8\xfd\x06\xf3\x97\x76\xe6\x48\xa3\x0c\x7b\x0b\x1e\x2b\xd4\xff\xee\x41\xdf\x65\xff\xaf\x6e\xec\xc7\x8a\x6d\xe8\xb7\x86\xbc\xd2\x21\x7f\x9b\x5e\x3f\x7a\xf3\x1d\xff\x2f\xc1\x1d\x00\x1a\x49\x49\x41\x48\x7d\xeb\x37\xb3\x59\xa8\x63\xd8\x90\xaa\x2c\x43\xa2\x55\xb5\x01\xa9\x50\x33\x0b\x4f\xa2\x51\xe9\xc3\x3f\x7f\x05\x00\x00\xff\xff\x23\xef\xf9\x17\xb9\x08\x00\x00")

func tmplSwaggerYamlBytes() ([]byte, error) {
	return bindataRead(
		_tmplSwaggerYaml,
		"tmpl/swagger.yaml",
	)
}

func tmplSwaggerYaml() (*asset, error) {
	bytes, err := tmplSwaggerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/swagger.yaml", size: 2233, mode: os.FileMode(420), modTime: time.Unix(1765572598, 0)}
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
	"tmpl/swagger.yaml": tmplSwaggerYaml,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"swagger.yaml": &bintree{tmplSwaggerYaml, map[string]*bintree{}},
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
