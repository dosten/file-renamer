package file

import "time"

type FileInfo interface {
	GetPath() string
	GetTime() time.Time
}
