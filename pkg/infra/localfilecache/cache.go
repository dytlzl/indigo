package localfilecache

import (
	"os"
	"path/filepath"
)

type cacheFile struct {
	name string
}

func NewCacheFile(name string) cacheFile {
	return cacheFile{name: name}
}

const cacheSubDir = "indigo"

func (f cacheFile) Write(data []byte) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(cacheDir, cacheSubDir), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(cacheDir, cacheSubDir, f.name), data, 0644)
}

func (f cacheFile) Read() ([]byte, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(filepath.Join(cacheDir, cacheSubDir, f.name))
}
