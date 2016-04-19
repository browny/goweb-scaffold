package config

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
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
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
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

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(f)
	}
	f, err := _escStatic.prepare(name)
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

var _escData = map[string]*_escFile{

	"/config/config.go": {
		local:   "config/config.go",
		size:    742,
		modtime: 1457324388,
		compressed: `
H4sIAAAJbogA/2RSzWrbQBA+a59iqkORIEh3Qw6lJaFQWkMfIFntjuR1pR2xGhFC8Lt3ZuXgiJwsvv1+
Zr5x28LRun92QHAU+zCAJ1yu32uyHCjCZKMQJoxs5h3ZmDDNlBgqU5QYHfkQh/a8UCwF6CfWn4jcnpjn
0tTGSN7jSJ0dv29pYQE+IXQhqhSoO6NjoP59GvWCPoxo+HXGvXbhtAr5zRTHRKr76RVTn2fVHcp5w5+C
L5/NJaf/Iusf3Uirv7qMArwv3OQ4pl2O6dfoPukqHQp0r+ZBvur9aDKTR6kDExzu8xLNb3z5sUFZWpti
X7LwPlq8XUyBKcuvTs0mr77udOITelDml3uIYdToQppvjlIEj7Eq5Y3SobxTkrClhiIhrynur/yhn2/L
gqyneUl2nmUHbSAPKdUktD5fBPpEE9jMdYIyeuhe1UNbObTtQJ5cQ2loh8CntWscTe10Dp3U0eLibsXm
wGq2fLoesL41qwttIVLFw9+qt+Oi7ekId/CkaH5u/swYs0d92zD/cy7mfwAAAP//G9zq/uYCAAA=
`,
	},

	"/config/config.json": {
		local:   "config/config.json",
		size:    36,
		modtime: 1457596613,
		compressed: `
H4sIAAAJbogA/6rm4lQqKMrPSk0uic9MUbJSUErPL09N0i1OTkxLy89JUeKq5QIEAAD//1urXUMkAAAA
`,
	},

	"/config/config_test.go": {
		local:   "config/config_test.go",
		size:    237,
		modtime: 1457596613,
		compressed: `
H4sIAAAJbogA/2SPv0oEMRCH652nGFJt5Ex6wUJERLCwuBfITf5czr3MmUwQEd/deLud1cDHx+9jLo7e
XQpIXGJOAPl84So4w6RSlmM/GOKzbVKD0LFaCU1y/LKutVBFDetKSlKgAWIvhPsBXtn5Z1q4+8fr7ix4
s4lmr/EbprWHd/f4z/0DD2NfZmVXbTvm1LgoPULT2jdPH90ts+xQJf4Mh9tGLkZevNptD5m3yqdA8uI1
/MBvAAAA//8mzT567QAAAA==
`,
	},

	"/config/seelog.xml": {
		local:   "config/seelog.xml",
		size:    364,
		modtime: 1457596613,
		compressed: `
H4sIAAAJbogA/1yQsWrDMBCG5+QpjgNBArUlm7pDsDyVTO3UsXRQ7LNrkCUjyWnz9pVi00I33a+f+z6u
9kTaDhBuM0n0N9MiTKPRdCUtsaPLMsRAfW8BOWcdNvtdbZcwL8FDb92kwthJnNRo0teubq3xVhNv9mly
VuvRDP2o6X97w3YqEEIqGDXFmV+V49GKD/aLLplvVd9b3eUxQkjlWYVAzkgUZS6KvBTiCSFpJpaXWCHc
4TXfNJPxyvZ3w/UNfyJrIJE9x/WHtDATRSZKKKqTeDyJKhdCHOGdvaRDfAB79QMc2Dk6PwA7L6Z9+7Qu
HJlBnmD8l1bz9cTN/icAAP//234GRGwBAAA=
`,
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
