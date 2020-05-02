// Package templates Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// readme.tpl
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

var _readmeTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xd1\x6f\x1a\x39\x10\xc6\xdf\xfd\x57\x8c\xc4\x3d\x04\x14\x96\x77\xa2\x3b\x29\x4a\xee\x74\xd1\x71\xa7\x28\x21\x27\x55\x08\x69\x27\xbb\x03\xeb\xc6\x6b\xbb\xf6\x2c\x29\x42\xfc\xef\x95\xed\x65\xb3\x50\xa2\xaa\x2f\x6d\xdf\xd8\x31\xfe\xc6\xf3\xfd\xbe\x19\xc0\x6e\x07\xd9\xbd\x33\x1f\xa9\xe0\xec\xb1\x32\x8e\xff\xc3\x9a\x60\xbf\x17\x62\xd1\x3f\x6a\xab\xcb\x8b\x7e\xf1\xe9\x61\x06\xfb\xfd\x10\xc6\x47\x2a\xb7\xe4\x0b\x27\x2d\x4b\xa3\xa3\xce\x60\x00\xf3\xd9\xd5\xed\xc3\x95\x10\x79\x9e\x17\x46\x7b\xa3\x48\xfc\x06\x15\xa9\x1a\x1c\x59\x03\x58\x96\x51\xe1\x81\xac\xf1\x92\x8d\xdb\x1e\xfa\x9d\x96\x53\xc7\xa3\xcb\x8d\x2d\x91\x3b\x3d\xa9\x3d\xa3\x52\xed\x3d\x45\xe8\xe9\x3d\xad\xb6\x3c\x09\xe5\x9b\x0a\x5d\x37\x24\x8c\xf5\x57\xf7\xbd\xc5\x22\xda\x92\xe7\x79\x1c\xe9\x4e\xb3\x33\x65\x53\x84\x31\x85\x98\x57\xd2\x43\x11\x44\xa0\x24\xab\xcc\xd6\x1f\x59\x72\x6d\x6d\xd0\x35\x1a\x10\x16\xff\x34\xcf\xe4\x34\x31\xf9\xe5\x45\xc5\x6c\xa7\x93\xc9\x4b\x57\xca\xa4\x19\x42\xa1\x1a\xcf\xe4\xa0\xf1\x52\xaf\x81\x2b\x82\xc5\xdf\xa4\xea\xf4\x77\x3f\x9d\x4c\xc2\xa8\x99\xaf\x86\x60\xb1\x78\xc1\x35\x41\x8d\x1a\xd7\xe4\xb2\xf8\xb4\x7b\x47\x8e\x3e\x35\xd2\x4b\x26\x2f\xc4\x6e\x07\x0e\xf5\x9a\xc2\x73\x7a\x07\x30\xde\xef\x45\x22\x17\xe6\xda\xed\xc6\x40\xba\x3c\x10\xbb\x4b\x3e\x1e\xfa\x47\x7f\x84\x98\x9b\xce\xe0\x50\x4d\x03\xbf\x4a\xae\xe2\xa7\x4b\x7e\x81\x0e\x36\xe6\x67\x08\xe4\xd3\x73\x11\xf8\x31\xc8\xe6\xe1\xbd\xa6\xae\x51\x97\xdf\x40\x14\x46\x79\x63\xd4\xc1\x90\xe9\xa4\xa4\x15\x36\x8a\xa1\x30\x7a\x25\xd7\x8d\xc3\x10\x80\x0c\x82\xfc\xe2\xa8\xb6\xbc\x18\x1c\x7d\x0f\xc1\x53\x4c\x0b\x28\xe9\xd9\x47\x31\x8b\x0e\x6b\x62\x72\xe1\x13\x19\x0a\xd4\xf0\x4c\x9d\x36\x95\x50\x36\x2e\x30\x68\x3d\x4a\xcd\x84\xf8\x03\x46\xa3\xb9\xb4\xa3\xd1\x14\x66\xd2\x33\x04\xfb\x5a\xf7\x7d\x9b\x9a\x3c\x7a\x1b\x5a\xa5\xbc\x3e\x69\xf9\x1e\xd2\xe6\x70\x34\x29\x49\x11\x53\x3c\xce\x3d\xa3\xaf\xc6\xc6\x92\x43\x36\xee\x3c\xba\xf6\xff\xe7\xc8\x7d\x17\x10\x47\xb5\xd9\x90\x87\x43\xb0\xfa\xf6\x9b\xda\x1a\x4d\x9a\x3d\xa0\xf7\xa6\x90\xc8\x54\xbe\x65\x2e\x45\x30\x41\x55\xf1\x42\x2f\x89\x69\x1d\x6e\xfa\x14\x52\xdf\x95\x51\xca\xbc\x46\x27\xf0\x59\x51\x8f\x48\x87\x2c\x94\x7b\x78\xcc\x2a\x1e\x3f\x06\x53\x7a\x4d\xb9\x22\xe9\xba\x4c\x6c\x50\x35\xe4\xb3\xb8\x72\x6d\x3e\xff\x8f\xa5\xb8\x55\x8f\x96\x0a\xb9\xda\x02\x61\x51\xbd\x49\xf7\x96\x3c\x1f\x8f\x3d\x31\xbc\xd0\xf6\xf7\xa8\xb4\xb8\xec\x7e\x2e\x73\x40\xb7\x6e\x6a\xd2\x0c\x6c\x5a\xb8\x2d\xb5\x3c\x83\xbf\x8c\x03\xfa\x8c\xb5\x55\xf4\x53\x56\x0c\xd2\xcb\x4f\xc7\xfe\x33\x3d\xa9\x43\x7e\xad\x98\x9c\x46\x96\x1b\x52\xdb\x4b\x40\xf8\x70\xfd\xef\x0c\x56\x52\x51\x8a\xbf\x8f\x0e\xc9\x16\x62\x72\x13\x56\xc6\x9d\xee\x4a\xbb\x26\xd6\x99\x8d\x2c\x43\x1a\x2a\xa9\x48\x9c\xe4\x3b\x42\xfa\x25\xac\x69\x07\x69\xd3\xb1\xc5\x5a\x45\x3b\xbe\x04\x00\x00\xff\xff\x5f\xae\x5b\x7f\x7c\x07\x00\x00")

func readmeTplBytes() ([]byte, error) {
	return bindataRead(
		_readmeTpl,
		"readme.tpl",
	)
}

func readmeTpl() (*asset, error) {
	bytes, err := readmeTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "readme.tpl", size: 1916, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"readme.tpl": readmeTpl,
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
	"readme.tpl": {readmeTpl, map[string]*bintree{}},
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