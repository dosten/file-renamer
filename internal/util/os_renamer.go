package util

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type osRenamer struct {
}

func NewOSRenamer() Renamer {
	return &osRenamer{}
}

func (renamer *osRenamer) Rename(old, new string) error {
	log.WithField("old", filepath.Base(old)).WithField("new", filepath.Base(new)).Info("renaming file")
	return os.Rename(old, new)
}
