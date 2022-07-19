package namer

import "github.com/dosten/file-renamer/internal/file"

type Namer interface {
	Name(file.FileInfo) (string, error)
}
