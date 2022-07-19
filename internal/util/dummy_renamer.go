package util

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type dummyRenamer struct {
}

func NewDummyRenamer() Renamer {
	return &dummyRenamer{}
}

func (renamer *dummyRenamer) Rename(old, new string) error {
	log.WithField("old", filepath.Base(old)).WithField("new", filepath.Base(new)).Info("rename file")
	return nil
}
