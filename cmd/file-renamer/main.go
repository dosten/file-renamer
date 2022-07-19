package main

import (
	"flag"
	"io/fs"
	"path/filepath"

	"github.com/dosten/file-renamer/internal/file"
	"github.com/dosten/file-renamer/internal/guesser"
	"github.com/dosten/file-renamer/internal/namer"
	"github.com/dosten/file-renamer/internal/util"

	log "github.com/sirupsen/logrus"
)

var dryRun bool

func init() {
	flag.BoolVar(&dryRun, "dry-run", false, "")
	flag.Parse()
}

func visit(guesser guesser.TimeGuesser, namer namer.Namer, renamer util.Renamer, path string) error {
	time, err := guesser.Guess(path)
	if err != nil {
		log.WithError(err).WithField("path", path).Warn("unable to guess time, keeping old name")
		return nil
	}

	info := file.NewFileInfo(path, time)

	new, err := namer.Name(info)
	if err != nil {
		log.WithError(err).WithField("path", path).Warn("unable to generate new name, keeping old name")
		return nil
	}

	err = renamer.Rename(path, new)
	if err != nil {
		log.WithError(err).WithField("path", path).Warn("unable to rename")
		return nil
	}

	return nil
}

func main() {
	guesser := guesser.NewMultipleTimeGuesser(guesser.NewFilenameTimeGuesser(), guesser.NewExifTimeGuesser())
	namer := namer.NewDateNamer()

	var renamer util.Renamer
	if dryRun {
		log.Info("dry run mode enabled")
		renamer = util.NewDummyRenamer()
	} else {
		renamer = util.NewOSRenamer()
	}

	root := flag.Arg(0)

	filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if entry.IsDir() {
			return nil
		}
		return visit(guesser, namer, renamer, path)
	})
}
