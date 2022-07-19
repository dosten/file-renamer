package file

import "time"

type fileInfo struct {
	path string
	time time.Time
}

func NewFileInfo(path string, time time.Time) FileInfo {
	return &fileInfo{path, time}
}

func (info *fileInfo) GetPath() string {
	return info.path
}

func (info *fileInfo) GetTime() time.Time {
	return info.time
}
