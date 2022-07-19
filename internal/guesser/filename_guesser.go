package guesser

import (
	"errors"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

var filenameRe = regexp.MustCompile("^((IMG|VID)[-_])?(?P<year>[0-9]{4})[-_/]?(?P<month>[0-9]{2})[-_/]?(?P<day>[0-9]{2})(.+)?$")

type filenameTimeGuesser struct {
}

func NewFilenameTimeGuesser() TimeGuesser {
	return &filenameTimeGuesser{}
}

func (guesser *filenameTimeGuesser) Guess(path string) (time.Time, error) {
	matches := filenameRe.FindStringSubmatch(filepath.Base(path))

	if matches == nil {
		return time.Time{}, errors.New("unable to guess file date from name")
	}

	year, err := strconv.Atoi(matches[filenameRe.SubexpIndex("year")])
	if err != nil {
		return time.Time{}, errors.New("unable to parse year")
	}

	month, err := strconv.Atoi(matches[filenameRe.SubexpIndex("month")])
	if err != nil {
		return time.Time{}, errors.New("unable to parse month")
	}

	day, err := strconv.Atoi(matches[filenameRe.SubexpIndex("day")])
	if err != nil {
		return time.Time{}, errors.New("unable to parse day")
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}
