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

var _tmplSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\xcd\x8a\xdb\x4c\x10\xbc\xeb\x29\x1a\xe3\xe3\x5a\x98\xef\xbb\x09\x72\xd9\x24\xec\x1e\xd6\xc1\x04\x27\xf7\x89\xd4\xb2\x86\xf5\xfc\x6c\xf7\x8c\x13\x61\xeb\xdd\xc3\x8c\x64\xfd\x66\x49\x16\x02\xd1\x49\xaa\xa9\xae\xea\x1a\x95\xb1\xa8\x85\x95\x19\xfc\x9f\x6e\xd3\x6d\x22\x75\x69\xb2\x04\xa0\x40\xce\x49\x5a\x27\x8d\xce\x60\x75\xb9\x40\xfa\xde\xe8\x52\x1e\xd3\x0f\xc3\x01\x34\xcd\x2a\x01\x38\x23\xf1\x82\xf6\xb5\x05\x3b\x8a\x93\xee\x84\x53\xc2\x21\x40\xf1\x98\x91\x82\x44\x70\xdd\x80\xa7\xd3\x94\x77\x2f\x18\xf7\xc2\x55\x91\xea\xc4\x91\xb3\xe4\x72\x01\x12\xfa\x88\xb0\x7e\xc6\xfa\x0e\xd6\x67\x71\xf2\x08\xd9\x3b\x48\x3f\x6a\x27\x9d\x44\x86\xa6\x49\x36\xa0\x85\xea\x4c\x03\x11\xae\xf0\x80\xee\x81\x8c\xb7\x9f\x84\xc2\x6e\xb3\x69\xce\x55\xd0\x46\x5d\x84\x79\x2b\x5c\x15\xcd\x36\x53\xb7\x88\x2f\xdc\x46\xb4\x40\xb8\x83\x35\x21\x1b\x4f\x39\x46\x6e\x37\xd5\x34\x09\x40\xd8\x27\x7c\xc2\x15\x3e\xa3\x32\xe7\x18\x6f\x4f\x58\xca\x1f\x4d\x13\x6e\x21\x50\x06\xd3\xe7\x91\x56\x94\x1a\x84\xa3\x5c\x2b\x78\x03\xd3\x1d\xba\xca\x14\x70\x85\x83\x79\x32\xdf\x91\xe0\xa6\x09\xd0\xde\x5e\xfb\xbe\xf9\xcd\xbd\x84\x87\xbd\x52\x82\xea\xee\x0a\x7b\x87\x19\x2b\xec\x2a\x4b\x30\x34\xa2\xec\x05\x09\x85\x0e\x89\x47\xe0\xbd\x29\xea\xf6\x64\x8c\x3e\xa2\x28\x02\xaf\xcb\xd2\xeb\xcd\xe6\x76\x52\xe1\xa1\xb6\x38\xf0\x16\x0d\x1d\x26\xc6\x1d\xbd\xc2\x17\x8d\x9c\x0b\x8b\x8f\x87\xdd\xd3\x6c\xef\xee\x57\xb7\x80\xed\x97\xce\xde\xb8\x4a\x5f\xb5\x6f\xa6\xa8\x6f\xfa\x00\x52\x2f\x20\xce\x2b\x54\x22\xeb\xbf\x01\x5c\x6d\x71\xbe\xff\xdc\x66\x98\x9f\x86\x26\x7c\xf1\xc8\x0e\xa6\x1e\x01\x95\x84\x45\x06\x8e\x3c\xbe\x96\x76\x5c\xd8\xd8\x6a\x12\x6a\x52\xaf\xd1\x4f\x5c\x06\x6d\x3b\x4c\x42\xcd\xdb\xd0\xa5\x7e\xf1\x48\x7f\x12\x9b\x1d\x49\x7d\x7c\x2d\xdf\xdf\xcf\xd4\xb5\xed\xed\x81\xaa\x38\xf8\x4f\x12\x4d\x00\x42\xb6\x46\x33\xf2\x60\xfd\xdf\x76\x3b\xde\x63\x6a\xc8\x3e\xcf\x91\xb9\xf4\x27\x30\x16\x49\x04\x78\x95\xcc\xa4\x7f\xfd\xfa\x33\x00\x00\xff\xff\x7d\x24\xd4\x9e\x1a\x06\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/swagger.yaml", size: 1562, mode: os.FileMode(420), modTime: time.Unix(1765491033, 0)}
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
