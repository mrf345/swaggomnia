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

var _tmplSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x16\x86\x8f\xb1\x60\xb4\x37\x01\xbd\xa4\x2d\x92\x43\x5c\x18\x85\xdb\x3b\x2b\xad\x6c\x22\xe6\x23\xbb\xa4\x5b\xc1\xd6\xbf\x17\xa4\x64\x3d\x13\xa4\x41\x11\x9f\xac\xe1\xec\xcc\x0e\x39\xc6\xa2\x16\x56\x66\xf0\x31\x5d\xa7\xeb\x44\xea\xd2\x64\x09\x40\x81\x9c\x93\xb4\x4e\x1a\x9d\xc1\xe2\x7c\x86\xf4\xb3\xd1\xa5\xdc\xa7\x5f\xfa\x03\xa8\xeb\x45\x02\x70\x42\xe2\x19\xed\x67\x03\xb6\x14\x27\xdd\x11\xc7\x84\x5d\x80\xe2\x31\x23\x05\x89\xe0\xba\x02\x4f\xc7\x31\xef\x56\x30\x6e\x85\x3b\x44\xaa\x13\x7b\xce\x92\xf3\x19\x48\xe8\x3d\xc2\xf2\x11\xab\x1b\x58\x9e\xc4\xd1\x23\x64\x9f\x20\xfd\xaa\x9d\x74\x12\x19\xea\x3a\x59\x81\x16\xaa\x35\x0d\x44\xb8\xc0\x1d\xba\x3b\x32\xde\x7e\x13\x0a\xdb\xcd\xc6\x39\x17\x41\x1b\x75\x11\xe6\xad\x70\x87\x68\xb6\x1a\xbb\x45\x7c\xe6\x36\xa0\x05\xc2\x0d\x2c\x09\xd9\x78\xca\x31\x72\xdb\xa9\xba\x4e\x00\xc2\x3e\xe1\x13\x2e\xf0\x1d\x95\x39\xc5\x78\x5b\xc2\x52\xfe\xa9\xeb\x70\x0b\x81\xd2\x9b\x3e\x0e\xb4\xa2\x54\x2f\x1c\xe5\x1a\xc1\x2b\x98\x6e\xd0\x1d\x4c\x01\x17\xd8\x99\x07\xf3\x1b\x09\xae\x9a\x00\xcd\xed\x35\xff\x57\xaf\xdc\x4b\xf8\xb1\x57\x4a\x50\xd5\x5e\x61\xe7\x30\x61\x85\x5d\x65\x09\x86\x06\x94\xad\x20\xa1\xd0\x21\xf1\x00\xbc\x35\x45\xd5\x9c\x0c\xd1\x7b\x14\x45\xe0\xb5\x59\x3a\xbd\xc9\xdc\x46\x2a\xdc\x55\x16\x7b\xde\xac\xa1\xfd\xc4\xb0\xa3\x17\xf8\xa1\x91\x73\x61\xf1\x7e\xb7\x79\x98\xec\xdd\x3e\x75\x03\xd8\x6e\xe9\xec\x8d\xab\x74\x55\xfb\x65\x8a\xea\xaa\x0f\x20\xf5\x0c\xe2\xfc\x80\x4a\x64\xdd\x37\x80\xab\x2c\x4e\xf7\x9f\xda\xf4\xf3\xe3\xd0\x84\x4f\x1e\xd9\xc1\xd8\x23\xa0\x92\xb0\xc8\xc0\x91\xc7\x97\xd2\x0e\x0b\x1b\x5b\x4d\x42\x8d\xea\x35\x78\xc4\x79\xd0\xa6\xc3\x24\xd4\xb4\x0d\x6d\xea\x27\x8f\xf4\x2f\xb1\xd9\x91\xd4\xfb\x97\xf2\x3d\x97\xa9\x14\x47\xfe\x8f\x50\x6d\xdd\xde\x9e\xe8\x10\x07\xdf\x25\xd2\x6b\xcf\x34\x02\x08\xd9\x1a\xcd\xc8\xbd\xf5\x87\xf5\x7a\xb8\xc7\xd8\x90\x7d\x9e\x23\x73\xe9\x8f\x60\x2c\x92\x08\xf0\x22\x99\x48\x3f\xff\xf7\x6f\x00\x00\x00\xff\xff\x78\x07\x79\xff\x1b\x06\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/swagger.yaml", size: 1563, mode: os.FileMode(420), modTime: time.Unix(1765550961, 0)}
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
