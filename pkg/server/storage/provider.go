package storage

import "io/fs"

type Provider interface {
    fs.ReadDirFS
    fs.ReadFileFS
}
