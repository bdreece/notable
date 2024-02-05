package storage

import (
	"io/fs"
	"os"
	"path"

	"github.com/bdreece/notable/pkg/server/config"
)

type localProvider struct {
	cfg *config.Storage
}

// Open implements Provider.
func (p *localProvider) Open(name string) (fs.File, error) {
	return os.Open(path.Join(p.cfg.RootDirectory, path.Clean(name)))
}

// ReadDir implements Provider.
func (p *localProvider) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(path.Join(p.cfg.RootDirectory, path.Clean(name)))
}

// ReadFile implements Provider.
func (p *localProvider) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(path.Join(p.cfg.RootDirectory, path.Clean(name)))
}

func NewLocalProvider(cfg *config.Storage) Provider {
	return &localProvider{cfg}
}
