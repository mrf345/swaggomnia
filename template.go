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

var _tmplSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x51\x6b\xdb\x3c\x14\x7d\xf7\xaf\xb8\x84\x3c\xd6\x26\x7c\xdf\x9b\x61\x0f\xed\x36\xd6\xd2\x76\x84\x35\x1b\x83\x31\x86\x6a\xdf\xc4\xa2\xb1\xa4\xde\x2b\x77\x33\xa9\xff\xfb\x90\xac\x38\x72\xd2\xd2\x94\xb2\x37\xfb\xf8\xe8\xe8\x1c\xdd\x23\x6b\x83\x4a\x18\x99\xc3\xff\xd9\x2c\x9b\x25\x52\x2d\x75\x9e\x00\x94\xc8\x05\x49\x63\xa5\x56\x39\x4c\x36\x1b\xc8\xde\x6b\xb5\x94\xab\xec\xc3\xee\x03\x74\xdd\x24\x01\x78\x40\xe2\x03\xda\xb7\x1e\x0c\x14\x2b\xed\x1a\xc7\x84\x85\x83\xfc\xe7\xcd\x26\x05\xb9\x84\xec\x82\xcf\x50\x10\x12\x74\x5d\x52\xe8\xda\x68\x85\xca\xb2\x33\xc3\x58\x34\x24\x6d\x7b\x53\x54\x58\xa3\x87\x00\x6e\x3d\xf9\xb4\xb1\x55\xff\x0e\x60\x5b\x83\x39\x54\xd6\x9a\x00\xb0\xe7\xe7\x81\x9a\x6c\x65\x1c\x3f\x8d\xd7\xc3\x8f\x9f\xde\x05\xaa\xb2\xeb\x62\x3f\x82\x65\x71\xac\x1d\xc7\x3d\xce\x8d\x63\xee\x9b\x19\x56\x3f\xe7\xe5\xd4\xc8\x4b\x6c\x8f\x34\x23\x3c\xf9\xd0\x4d\x8f\x07\x48\xaa\x1c\x2a\x14\x25\x52\x00\x94\x70\xf6\xbe\xa7\xa7\xf3\x8b\xf4\x12\x45\xbb\xe7\x31\x12\x1d\x9b\x64\x24\xd7\x81\x9e\xd5\xd0\x7a\x3c\xe8\x33\xc1\x38\x17\xb6\xf2\xb3\xb6\x62\xc5\x79\xb2\xd9\x00\x09\xb5\x42\x98\xfe\x3a\x81\x29\x42\xfe\x0e\xb2\x1b\x4d\x16\xcb\x8f\xca\x4a\x2b\x91\x5d\xce\x34\x18\x72\x62\x53\xcc\x3e\x91\x6e\x4c\xe8\xd3\xb8\x9d\xae\x42\xce\x8a\x5b\x64\x84\xad\xfc\x0e\xe9\x31\x5b\x44\x34\xb7\xf0\x04\xa6\x84\xac\x1b\x2a\x90\xdd\x8a\x29\x66\x9e\xed\x8f\x1d\xc0\xf9\x70\x34\x78\x84\x2f\x58\xeb\x07\x1f\x6b\x4e\xb8\x94\x7f\xba\xae\x3f\xe8\x48\xf0\xee\x2e\x92\xf3\x6a\x3b\x6d\x2f\xd7\x0b\x6e\xc1\xec\x1a\x6d\xa5\x4b\x78\x84\x85\xbe\xd2\xbf\xfd\x35\x18\x86\xe7\x4f\xad\x7f\x4e\x9f\x3a\x0f\x5f\xaf\xa6\xae\x05\xb5\xe1\xbc\x06\xd9\xcf\xa2\xc6\x88\x15\x0a\xa5\x29\xa2\xcc\x05\x89\x1a\x2d\x12\x47\xe0\x99\x2e\xdb\xfe\x4b\x8c\x9e\xfb\xbe\x1c\xf0\xae\x65\x8d\x8b\xd6\xe0\x36\xd8\xb0\xcf\x8b\xbc\x83\xff\xcc\x6e\x45\xfc\xa7\x79\x84\xaf\x0a\xb9\x10\x06\xcf\x17\xd7\x57\x7b\x79\xc2\xe8\x7b\xc0\x0c\x61\xf2\x57\x5a\x19\xfa\x76\xab\xcb\x76\xab\xdf\x5f\x93\x3d\xc8\xdf\x64\x91\x0f\xef\xdb\xeb\x35\x61\x4b\x52\xad\x76\xc4\x71\x3a\xc2\xfb\x06\xd9\xc2\x58\xcc\xa1\x92\xb0\xcc\xc1\x52\x83\xcf\xc5\x8a\x9b\x7a\xe2\x6a\x48\xa2\x1e\x95\x2a\x9a\xe2\x61\xa2\xbe\xb9\x24\xea\xfd\x3a\x84\x78\xf7\x0d\xd2\xdb\xf3\x3d\x95\x69\x29\xd6\xfc\x86\x50\xa1\x6f\xaf\x4f\xd4\xff\xd8\xfe\x49\xa4\x97\xc6\x34\x02\x08\xd9\x68\xc5\xc8\xbb\xad\xff\x9b\xcd\x62\x1f\xe3\x0d\xb9\x29\x0a\x64\x5e\x36\x6b\xd0\x06\x49\x38\x78\x92\xec\x49\x3f\xfd\xf8\x37\x00\x00\xff\xff\xd2\x81\x82\x3c\xca\x07\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/swagger.yaml", size: 1994, mode: os.FileMode(420), modTime: time.Unix(1765560328, 0)}
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
