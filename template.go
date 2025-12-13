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

var _tmplSwaggerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\xc1\x6e\xd3\x40\x10\xbd\xfb\x2b\x46\x51\x8e\xb5\x15\xc1\xcd\x12\x87\x94\xa2\xb6\x6a\x8b\xa2\x36\x20\x24\x84\xd0\xd6\x9e\xc4\xab\xc6\xbb\xdb\x9d\x75\xc1\x4a\xf2\xef\x68\xd7\x6b\x7b\xed\xb4\x90\xc2\x81\x5b\xf2\xf2\x66\xe6\xed\x9b\x37\xad\x54\x28\x98\xe2\x29\xbc\x4d\x66\xc9\x2c\xe2\x62\x25\xd3\x08\x20\x47\xca\x34\x57\x86\x4b\x91\xc2\x64\xbb\x85\xe4\xbd\x14\x2b\xbe\x4e\xce\xfa\x1f\x60\xbf\x9f\x44\x00\x4f\xa8\xe9\x80\xf6\xb9\x01\x3d\xc5\x70\xb3\xc1\x21\x61\x69\x21\xf7\xf3\x76\x1b\x03\x5f\x41\x72\x49\xa7\xc8\x34\x6a\xd8\xef\xa3\x4c\x96\x4a\x0a\x14\x86\xac\x18\xc2\xac\xd2\xdc\xd4\x77\x59\x81\x25\x3a\x08\xe0\xde\x91\xe7\x95\x29\x9a\xef\x00\xa6\x56\x98\x42\x61\x8c\xf2\x00\x39\x7e\xea\xa9\x51\xdb\xc6\xf2\xe3\xb0\x1e\xbe\x7e\x73\x2a\x50\xe4\xfb\x7d\xa8\x87\x11\xcf\x8e\x95\x63\xb9\xc7\xa9\xb1\xcc\xb1\x98\xae\xfa\x25\x2d\x73\xc5\xaf\xb0\x3e\x52\x0c\x73\xe4\x43\x35\x0d\xee\x21\x2e\x52\x28\x90\xe5\xa8\x3d\x20\x98\x95\xf7\x25\x9e\x2f\x2e\xe3\x2b\x64\xf5\x48\x63\xd0\x74\x28\x92\x50\xdb\x0c\x34\xac\x4a\x6f\x86\x8b\x3e\x65\x84\x0b\x66\x0a\xb7\x6b\xc3\xd6\x94\x46\xdb\x2d\x68\x26\xd6\x08\xd3\xef\x27\x30\x45\x48\xdf\x41\x72\x27\xb5\xc1\xfc\x83\x30\xdc\x70\x24\xfb\xce\xd8\x0b\xb2\xcd\xa6\x98\x9c\x6b\x59\x29\x9f\xa7\x61\x3a\x6d\x84\xac\x14\x5b\xa4\x98\x29\xdc\x84\xf8\x98\x11\x01\xcd\x16\x9e\xc0\x54\x23\xc9\x4a\x67\x48\xb6\x62\x8a\x89\x63\x3b\xdb\x01\xac\x0e\x4b\x83\x1d\xdc\x62\x29\x9f\xdc\xb3\x16\x1a\x57\xfc\xe7\x5c\xe4\xb7\xa8\x36\x2c\xc3\x05\xd3\xac\xb4\xcd\x1b\xe7\x83\x09\x0f\x0f\x41\x7f\xd7\xbe\x1f\xe6\xfa\x37\x13\x5a\x30\xb9\x41\x53\xc8\x1c\x76\xb0\x94\xd7\xf2\x87\xbb\x8b\x6e\x9b\xce\xc6\xe6\x73\xfc\x9c\x41\x2e\x6f\x55\x59\x32\x5d\x7b\x03\xbb\xb6\x1f\x59\x89\x01\xeb\xe0\xd0\x7b\x66\x78\xea\x3b\xf8\x24\x90\x32\xa6\xf0\x62\x79\x73\x1d\xd4\xfb\x84\x4a\x1d\x14\x3a\x0f\xd0\xa0\xa6\x00\x3c\x95\x79\x9d\x78\x77\x7a\xf4\xc2\x05\xf0\x80\x77\xc3\x4b\x5c\xd6\x0a\x5b\x63\x00\x54\xd7\x33\x0d\x26\x07\x3b\x56\xca\x79\xea\x17\x74\x8e\xc6\x6d\xa7\xdd\x46\x67\x56\x90\x29\x15\xba\xd5\xdc\xc3\xc4\x96\xf7\x90\x3b\x59\x96\x76\xdf\xdb\x3b\x9a\x90\xd1\x5c\xac\x7b\xe2\x38\x90\x2d\xae\xf1\xb1\xe2\x1a\xf3\x14\x8c\xae\x30\x10\xee\xe3\x3a\xf0\xf0\x8f\x1e\x74\xea\xef\x65\x5e\x8f\x84\x0f\xa1\xbf\x14\x6e\xe5\x22\x19\x18\x36\x3b\xfe\x11\xed\x2d\xd9\x75\x58\xe7\x07\x29\x0f\x62\x31\x7e\xb9\x90\xc6\x57\x24\x67\x9c\xd8\xfd\x06\xf3\x97\x76\xe6\x48\xa3\x0c\x7b\x0b\x1e\x2b\xd4\xff\xee\x41\x3f\xe5\xf0\x5f\xdd\xd8\x8f\x15\xdb\xd0\x6f\x0d\x79\xa5\x43\xfe\x1a\x5e\xff\xf4\xe6\xef\xf8\x7f\x09\xee\x00\xd0\x48\x4a\x0a\x42\xea\x47\xbf\x99\xcd\x42\x1d\xc3\x81\x54\x65\x19\x12\xad\xaa\x0d\x48\x85\x9a\x59\x78\x12\x8d\x5a\x3f\xff\xf1\x57\x00\x00\x00\xff\xff\x02\x14\x8d\xd4\xb9\x08\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/swagger.yaml", size: 2233, mode: os.FileMode(420), modTime: time.Unix(1765605483, 0)}
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
