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
		size:    1737,
		modtime: 1465291897,
		compressed: `
H4sIAAAJbogA/8xUTW/bOBA9i7+CS2ABKpuVsNhbgBzSNGmDFolRF720RUBTQ4eJRAok5cQo+t87JCVb
bpv2Wh9MkfP15r0h65ouhHwQa6DSGqXXtLHgx+/BiaCtoZ0w6NCBCaQ/cCZEd711gXJSsNU2gGf4oboQ
F21rbYeg27hp7TouBkJ9F0LPCG7WOtwNq0rarva9+u//eqN7cIyUhCCst1Y0Z95DoNrTRyd6tFE1GJkg
BUsdiIYq3QJVznZUJF+JpwEautrGHLHUSV2vbWNlZd26npXs7vXKW1ODlySm3RfkvQh31AenzbpMOarL
WOYLKXKRk1N6ueRKtB5KUkQIx/Q2niZzddODSTmwkcJBGJxJOMnX1NhZ359bE+ApjMR6is21WmayZbaR
sO1h7ot4BhkiiAuzobtfhkmKhbP3IMPVy/nhclBKP92oN9o0u8OLIJsluA04T+nHz9Pxue06a87zEBwl
KaoP8X+EHemhLf5N05FpH7+3omtzk4lLLsMTPdqjL1M4B0Q+8YqN4FBUC9yF1nCW8iMREwEssjfVTYE9
SK20nGavwMMRLlKfAS+HVSyCqiCAKhJ1GmPzdk9ROsyx1SsIywSJsz473OqGjRkOCHwmqhFB+GAdVA/o
detTyJRgTvZP45coPHAG6HfQsUxy7HqNuQ4UmnfMsjPGZ6mSbOl+eKpNdpyEWm2pl8JgZRpH1Ge9UgRP
okwA5lGkGItByOWvRYeg830tJ+tZ02TrAhNzVj1vqXPa5x1+64F3uvZO4vIIq3+xI6Vs28yiwLn9WLxD
Kq5G5jgataLR/tcpNbqNTRc9UiI5vlzVhXPWKc4uUdY2utkdd3G+T+jfnn4y7DiaSsyFlO8HWY2koNc/
G/QZsbctMheQc8/LA5FeC9f8Qqg88dDMbtYuLKv1nS7v8dHgLF7GyEFcx5cpv8SJCETD928dGznLj2+1
jZFxDn/gb2QvvfLVNTy+wDnHmYk1yj+A0m8BAAD///HSTzrJBgAA
`,
	},

	"/config/config_test.go": {
		local:   "config/config_test.go",
		size:    132,
		modtime: 1465287110,
		compressed: `
H4sIAAAJbogA/ypITM5OTE9VSM7PS8tM5+LKzC3ILypR0ODiVCpJLS7JzEtX4tLk4korzUtWCAEKhGUW
pBZplChoQWX1QjQVqrk4IcKaXLXoSj0Si1JwKAdLgbQAAgAA//+dCaWbhAAAAA==
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

	"/config/viper.yml": {
		local:   "config/viper.yml",
		size:    575,
		modtime: 1465280095,
		compressed: `
H4sIAAAJbogA/7SRzU7EIBDH730Kwh0Wpl+Wm3cPvkHDUjZF2YW0Y/TxnTbrxjQ2PXki+X/M/MIIIYqX
5Gw0BWN5Sm/eYR8Gw2zO0ZM2WLQzpskvAcbew23o54/LJXwZthZJ9uiosfqC8RExm9NJq0rqrpO6pVc3
Bsq243uZppS6avczdS1B1/R298xzzKPdMp9tjAfIa+8HebNp0YT9E2K1zvuWu1uv07CFchYPmDj/PyCX
rtd0WwZ/jgF9H8OMjzUASmp4kgD0/UrxxwlBKwlVJRtZksp+6ZruQNmlQ9OL7wAAAP//+gz4sj8CAAA=
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
