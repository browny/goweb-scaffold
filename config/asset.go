package config

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _esc_localFS struct{}

var _esc_local _esc_localFS

type _esc_staticFS struct{}

var _esc_static _esc_staticFS

type _esc_file struct {
	compressed string
	size       int64
	local      string
	isDir      bool

	data []byte
	once sync.Once
	name string
}

func (_esc_localFS) Open(name string) (http.File, error) {
	f, present := _esc_data[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_esc_staticFS) prepare(name string) (*_esc_file, error) {
	f, present := _esc_data[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		gr, err = gzip.NewReader(bytes.NewBufferString(f.compressed))
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _esc_staticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (f *_esc_file) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_esc_file
	}
	return &httpFile{
		Reader:    bytes.NewReader(f.data),
		_esc_file: f,
	}, nil
}

func (f *_esc_file) Close() error {
	return nil
}

func (f *_esc_file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_esc_file) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_esc_file) Name() string {
	return f.name
}

func (f *_esc_file) Size() int64 {
	return f.size
}

func (f *_esc_file) Mode() os.FileMode {
	return 0
}

func (f *_esc_file) ModTime() time.Time {
	return time.Time{}
}

func (f *_esc_file) IsDir() bool {
	return f.isDir
}

func (f *_esc_file) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _esc_local
	}
	return _esc_static
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _esc_local.Open(name)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(f)
	}
	f, err := _esc_static.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _esc_data = map[string]*_esc_file{

	"/config/.DS_Store": {
		local: "config/.DS_Store",
		size:  6148,
		compressed: "" +
			"\x1f\x8b\b\x00\x00\tn\x88\x00\xff\xec\x97\xcdJ\xc3@\x14\x85ύ\x11\xa6\x16$K\x97\xb3t%\xf8\x06\xa1T\xc1\xb5{\u007fZ\xa3(\xd1,Tp\x99\a\xf4\x01|\x00\xdfC\xefd\x8e\x92N\xe3\xc2UK{\xbf2|m:g~H\xc8\xcc\x00\x90\xc9\xeb\xcd1P\xe8W\x87\xe8l\x84A\x1c\xcb\x12\x19\xbd\x1b\xda\xeb\xdax\xc1\x11\xee\xd0" +
			"\x9c\xd5\xcd|\xb8-c\xcd\b\xf7\xce\xe1\x1a\xcf\xfa\xa9\xfa\xf7oV73\xc4\a\xe3T\xcb\xe1WG\x97\x19a\x8e\x06O\xb8Ž\xd6\x1eHd\x17Ib\x9c$\x1e\xb47\xfd\x9d\xa4>\x93\xd4\xfeB\xeaRGWinp\x8c\xf2\x91d\xf7\xba\xf9T\xa8\xb5\x85\xd0\xe3\x1b\x1eQ'\xf3zל\xebe\f\xc30\xb6\x05\x89r\xe3\xd5\x0e\xc3" +
			"0\x8c5$\xbc\x1f<]\xd2m\xb4\xf0\xff\x8c\xce{\x99\x82\xf6tI\xb7\xd1\xc2z\x19\x9dӎ.hO\x97t\x1b͗\x96\xf0\xf0!\xecYxB\x91\x82\xf6t\xf9\xcfI\x1bƖ\xb0\x13U\x84\xf5\xff\xe4\xef\xf3\xbfa\x18\x1b\x8c\xe4\xd3\xf3\xe9\x04\xbf\a\x82%\xc2Z\xeb\xb5\\\xfd\x04\xb0\xb8\x11\u0d5cu\xc3R|л\xee\xe9\x92" +
			"n\xa3m#`\x18\xab\xe2;\x00\x00\xff\xff\xfcץ\xfa\x04\x18\x00\x00",
	},

	"/config/config.go": {
		local: "config/config.go",
		size:  742,
		compressed: "" +
			"\x1f\x8b\b\x00\x00\tn\x88\x00\xffdR\xcdj\xdb@\x10>k\x9fb\xaaC\x91 HwC\x0e\xa5%\xa1PZC\x1f Y\xed\x8e\xe4u\xa5\x1d\xb1\x1a\x11B\xf0\xbbwf\xe5\xe0\x88\x9c,\xbe\xfd~f\xbeq\xdb\xc2Ѻ\u007fv@p\x14\xfb0\x80'\\\xae\xdfk\xb2\x1c(\xc2d\xa3\x10&\x8cl\xe6\x1d٘0͔\x18*S" +
			"\x94\x18\x1d\xf9\x10\x87\xf6\xbcP,\x05\xe8'֟\x88ܞ\x98\xe7\xd2\xd4\xc6H\xde\xe3H\x9d\x1d\xbfoia\x01>!t!\xaa\x14\xa8;\xa3c\xa0\xfe}\x1a\xf5\x82>\x8ch\xf8uƽv\xe1\xb4\n\xf9\xcd\x14\xc7D\xaa\xfb\xe9\x15S\x9fg\xd5\x1d\xcayß\x82/\x9f\xcd%\xa7\xff\"\xeb\x1f\xddH\xab\xbf\xba\x8c\x02\xbc/\xdc\xe4" +
			"8\xa6]\x8e\xe9\xd7\xe8>\xe9*\x1d\nt\xaf\xe6A\xbe\xea\xfdh2\x93G\xa9\x03\x13\x1c\xee\xf3\x12\xcdo|\xf9\xb1AYZ\x9bb_\xb2\xf0>Z\xbc]L\x81)˯N\xcd&\xaf\xbe\xeet\xe2\x13zP\xe6\x97{\x88a\xd4\xe8B\x9ao\x8eR\x04\x8f\xb1*\xe5\x8dҡ\xbcS\x92\xb0\xa5\x86\"!\xaf)\xee\xaf\xfc\xa1\x9fo\xcb" +
			"\x82\xac\xa7yIv\x9ee\am \x0f)\xd5$\xb4>_\x04\xfaD\x13\xd8\xccu\x822z\xe8^\xd5C[9\xb4\xed@\x9e\\Cih\x87\xc0\xa7\xb5k\x1cM\xedt\x0e\x9d\xd4\xd1\xe2\xe2n\xc5\xe6\xc0j\xb6|\xba\x1e\xb0\xbe5\xab\vm!R\xc5\xc3ߪ\xb7\xe3\xa2\xed\xe9\bw\xf0\xa4h~n\xfe\xcc\x18\xb3G}\xdb0\xffs" +
			".\xe6\u007f\x00\x00\x00\xff\xff\x1b\xdc\xea\xfe\xe6\x02\x00\x00",
	},

	"/config/config.json": {
		local: "config/config.json",
		size:  36,
		compressed: "" +
			"\x1f\x8b\b\x00\x00\tn\x88\x00\xff\xaa\xe6\xe2T*(\xca\xcfJM.\x89\xcfLQ\xb2RPJ\xcf/OM\xd2-NNLK\xcb\xcfIQ\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff[\xab]C$\x00\x00\x00",
	},

	"/config/config_test.go": {
		local: "config/config_test.go",
		size:  237,
		compressed: "" +
			"\x1f\x8b\b\x00\x00\tn\x88\x00\xffd\x8f\xbfJ\x041\x10\x87띧\x18Rm\xe4Lz\xc1BDD\xb0\xb0\xb8\x17\xc8M\xfe\\ν̙L\x10\x11\xdf\xddx\xbb\x9d\xd5\xc0\xc7\xc7\xefc.\x8e\xde]\nH\\bN\x00\xf9|\xe1*8äR\x96c?\x18\xe2\xb3mR\x83бZ\tMr\xfc\xb2\xae\xb5PE\r\xebJJ" +
			"R\xa0\x01b/\x84\xfb\x01^\xd9\xf9gZ\xb8\xfb\xc7\xeb\xee,x\xb3\x89f\xaf\xf1\x1b\xa6\xb5\x87w\xf7\xf8\xcf\xfd\x03\x0fc_feWm;\xe6Ը(=B\xd3\xda7O\x1f\xdd-\xb3\xecP%\xfe\f\x87\xdbF.F^\xbc\xdam\x0f\x99\xb7ʧ@\xf2\xe25\xfc\xc0o\x00\x00\x00\xff\xff&\xcd>z\xed\x00\x00\x00",
	},

	"/config/seelog.xml": {
		local: "config/seelog.xml",
		size:  355,
		compressed: "" +
			"\x1f\x8b\b\x00\x00\tn\x88\x00\xff\\\x90\xb1j\xc30\x10\x86\xe7\xe4)\x8e\x03A\x02\x8du\tu\x87`y*\x99کc\xe9\xa0\xd8\x17\xd7 KF\x92\xd3\xe6\xed+Ŧ\x85nw\xbf\u007f\xee\xfb\xac*0\x1b\xd7A\xbc\x8d\xac0\xdcl\x830\xf4\xd6\xf0\x95\x8d\u0096\xcfS\x97\x02\xfd\xbd\x04\xec\xbd\xf3X\xafW\x95\x9b\xe28\xc5\x00\x17" +
			"\xe7\a\x1d\xfbV\xe1\xa0{\x9b?\xad\xaa\xc6\xd9\xe0\f\xcbz\x9d7\xef\x8c\xe9mw\xe9\r\xffo/\xd8VGF\xc8\x05\xab\x87\xb4˫\xf62Y\xc9\xce}\xf1\xb9H\x13B\xee\x8c:F\xf6V!\x1d\n\xda\x17\a\xa2'\x84l\x97\x11Aa\x89pgVr\xb1ˢ32\xdc\xc5\xe6\x19\xfe\xf8s\xa0P<\xa7\xf3\x9b|pG\xfb\x1d" +
			"\x1d`_\x1e\xe9\xf1HeAD[x\x17/\xf9\xff?@\xbc\x86\x0e6\xe2\x94T\x1f@\x9c&ۼ}:\x1f\xb7¢\xcc0\xf9K\xab\xe4\xfc\xb2\xf5\xfa'\x00\x00\xff\xffL\xf3\x0e\xd7c\x01\x00\x00",
	},

	"/": {
		isDir: true,
		local: "/",
	},

	"/config": {
		isDir: true,
		local: "/config",
	},
}
