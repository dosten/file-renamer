package guesser

import (
	"errors"
	"mime"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

type exifTimeGuesser struct {
	guessers []TimeGuesser
}

func NewExifTimeGuesser() TimeGuesser {
	return &exifTimeGuesser{}
}

func (guesser *exifTimeGuesser) Guess(path string) (time.Time, error) {
	ext := filepath.Ext(path)
	if ext == "" {
		return time.Time{}, errors.New("missing extension")
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return time.Time{}, errors.New("unrecognized type")
	}

	switch mimeType {
	case "image/jpeg", "image/png":
		f, err := os.Open(path)
		if err != nil {
			return time.Time{}, errors.New("unable to open file")
		}
		defer f.Close()

		x, err := exif.Decode(f)
		if err != nil {
			return time.Time{}, errors.New("unable to decode exif data")
		}

		date, err := x.DateTime()
		if err != nil {
			return time.Time{}, err
		}

		return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC), nil
	case "video/mp4":
		return time.Time{}, errors.New("not implemented")
	}

	return time.Time{}, errors.New("type not supported")
}
