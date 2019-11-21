// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// graphiql.html
// schema.gql
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

var _graphiqlHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x51\x4f\xdb\x30\x10\x7e\xef\xaf\xf0\x2c\x4d\x4a\x25\xb0\x1b\x36\xf1\x90\xa6\x7d\x60\xb0\x69\x13\x1b\xb0\xed\x65\x8f\xc6\xbe\xd4\x66\x8e\x1d\xce\x4e\xa1\x42\xfc\xf7\x29\x09\x85\xd0\xa6\xd3\xa6\x0d\xbf\xd4\x3e\xdf\x7d\xdf\x77\xd7\xcf\x4a\xfe\xea\xf8\xec\xdd\xf7\x1f\xe7\x27\x44\xc7\xd2\xce\x47\x79\xf7\x43\x08\x21\xb9\x06\xa1\xba\x6d\x7b\xb4\xc6\xfd\x24\x1a\xa1\x98\x51\x1d\x63\x15\x32\xce\xa5\x72\x57\x81\x49\xeb\x6b\x55\x58\x81\xc0\xa4\x2f\xb9\xb8\x12\xb7\xdc\x9a\xcb\xc0\x17\x28\x2a\x6d\xae\x2d\x9f\xb0\x34\x65\x69\xfa\x18\x60\xa5\x71\x4c\x86\x40\x09\x82\x9d\xd1\x10\x57\x16\x82\x06\x88\x94\xf0\x1e\x63\x90\x68\xaa\x48\x02\xca\x3f\xa6\x84\x70\xb8\x5f\xa1\x2f\x4d\x00\xfe\x96\xa5\x2c\xed\x47\x98\xa8\xa3\x6f\xb9\xaf\x02\x9d\xe7\xbc\xc3\xff\x37\xc2\x02\xa2\xd4\xfc\x80\x4d\xd8\x9b\x6e\xff\x9f\xf1\x11\x84\x8c\x3c\x3d\x64\x07\x6c\xc2\xeb\x52\x75\x01\x56\xa1\x57\xb5\x8c\xc6\xbb\x97\xe0\xdb\x57\xbe\xdc\xe2\x6c\x82\x2f\xc7\xfb\x7b\xaf\x6c\xf3\xe4\xfc\xc9\x9e\xf9\xa5\x57\x2b\xd2\xba\x68\x46\x6f\x8c\x8a\x3a\x23\xe9\x64\xf2\x7a\x4a\x34\x98\x85\x8e\xeb\x53\x29\x70\x61\x5c\x46\x26\x53\xe2\x97\x80\x85\xf5\x37\x19\xd1\x46\x29\x70\x53\xda\x93\xaf\xcc\x92\x18\x35\xa3\x6b\x09\x74\x8d\xdd\x83\x5b\xea\x29\x9d\x9f\x7a\xa1\x8c\x5b\x30\xc6\x72\xae\xcc\x72\x6b\x02\x4f\x81\x66\x15\xb5\x6b\x27\x47\x5a\xdc\x8b\xd3\xf7\x8d\x5d\x00\x93\x87\xe3\xb9\x40\x51\x86\x31\xb9\x7b\x56\xd4\x2c\x84\x58\xa3\x23\xad\xbd\x12\xda\x4d\xe6\xda\xd2\xbd\x81\xd4\x66\x95\x10\xb5\x57\x19\xa1\x95\x0f\x91\xee\x0d\xe6\x34\x13\xcb\xc8\xa7\x6f\x67\x5f\x58\x88\x68\xdc\xc2\x14\xab\x0d\x21\xc3\x85\x12\x41\x81\x8b\x46\xd8\x90\x11\x6a\x9c\xb4\xb5\x82\x01\x92\xfb\x31\x8b\x1a\x5c\xf2\xd8\x74\x82\x10\x2a\xef\x02\x0c\x75\xd8\xeb\x72\x9d\xc6\x22\xdc\xc6\x64\x3c\xfd\x0b\xe4\x23\xaf\x56\xbb\xd0\x23\xae\x76\xdc\xf4\xb8\xdb\x79\x54\x02\x03\x3c\xc7\xdc\x16\xd1\x0a\x21\x52\x44\xa9\x49\x02\x88\x1e\x77\x11\x0f\xb4\xd6\x60\xee\x80\x1c\xe8\xf6\x79\xe6\xfd\xe8\x6b\xf3\x18\x8f\xcf\x3e\x33\x04\xa7\x00\x93\xf6\xba\x0d\x32\x89\x20\x22\x9c\x58\x28\xc1\xc5\xe4\x43\x6b\xdf\x8b\xd3\x3d\x72\x57\x74\x56\xcb\x36\xac\x77\xff\xf0\x27\x2b\x2f\xeb\xa6\x84\x2d\x20\x3e\x54\x1f\xad\x3e\xaa\xe4\xe9\x05\x8c\x47\x3d\x1d\x9b\x2f\xb1\x31\xd3\x7c\x94\xf3\xee\xdb\xf1\x2b\x00\x00\xff\xff\xe6\x51\xa1\x55\x53\x06\x00\x00")

func graphiqlHtmlBytes() ([]byte, error) {
	return bindataRead(
		_graphiqlHtml,
		"graphiql.html",
	)
}

func graphiqlHtml() (*asset, error) {
	bytes, err := graphiqlHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "graphiql.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\x0a\x0a\xb9\xa5\x25\x89\x25\x99\xf9\x79\x56\x0a\xbe\x50\x16\x57\x2d\x17\x57\x49\x65\x41\x2a\x44\x09\x58\x4f\x46\x6a\x4e\x4e\xbe\x46\x6e\x71\xba\x55\x70\x49\x51\x66\x5e\xba\xa6\x95\x02\x84\xa1\xc8\xa5\xa0\x50\x94\x9a\x9c\x9a\x59\x96\xea\x9b\x5a\x5c\x9c\x98\x9e\x8a\x90\x81\x19\x03\x33\x17\x6c\x52\x71\x6a\x5e\x0a\x54\x25\xaa\x79\x9e\x2e\x20\x1d\x80\x00\x00\x00\xff\xff\x3f\x3d\xe6\xd6\xa8\x00\x00\x00")

func schemaGqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGql,
		"schema.gql",
	)
}

func schemaGql() (*asset, error) {
	bytes, err := schemaGqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.gql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"graphiql.html": graphiqlHtml,
	"schema.gql":    schemaGql,
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
	"graphiql.html": &bintree{graphiqlHtml, map[string]*bintree{}},
	"schema.gql":    &bintree{schemaGql, map[string]*bintree{}},
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
