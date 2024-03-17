package assets

import (
	"embed"
	"io/fs"
)

//go:embed all:web/*
var WebFiles embed.FS

func LoadFile() fs.FS {
	subFs, _ := fs.Sub(WebFiles, "web/assets")
	return subFs
}
