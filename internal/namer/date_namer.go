package namer

import (
	"errors"
	"fmt"
	"mime"
	"path/filepath"
	"sync"

	"github.com/dosten/file-renamer/internal/file"
)

type dateNamer struct {
	mu       sync.Mutex
	counters map[int64]int
}

func NewDateNamer() Namer {
	return &dateNamer{
		mu:       sync.Mutex{},
		counters: make(map[int64]int),
	}
}

func (namer *dateNamer) Name(info file.FileInfo) (string, error) {
	dir := filepath.Dir(info.GetPath())

	ext, err := namer.normalizeExtension(info.GetPath())
	if err != nil {
		return "", errors.New("unable to normalize extension")
	}

	namer.mu.Lock()
	defer namer.mu.Unlock()

	key := info.GetTime().Unix()

	c, ok := namer.counters[key]
	if ok {
		namer.counters[key] += 1
	} else {
		namer.counters[key] = 1
	}

	return fmt.Sprintf("%s/%s-%04d%s", dir, info.GetTime().Format("2006-01-02"), c, ext), nil
}

func (namer *dateNamer) normalizeExtension(path string) (string, error) {
	ext := filepath.Ext(path)
	if ext == "" {
		return "", errors.New("unrecognized extension")
	}

	mimeType := mime.TypeByExtension(ext)

	switch mimeType {
	case "image/jpeg":
		return ".jpeg", nil
	case "image/png":
		return ".png", nil
	case "image/gif":
		return ".gif", nil
	case "video/mp4":
		return ".mp4", nil
	case "video/quicktime":
		return ".mov", nil
	}

	return ext, nil
}
