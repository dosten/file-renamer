package guesser

import (
	"fmt"
	"time"
)

type multipleTimeGuesser struct {
	guessers []TimeGuesser
}

func NewMultipleTimeGuesser(guessers ...TimeGuesser) TimeGuesser {
	return &multipleTimeGuesser{guessers}
}

func (guesser *multipleTimeGuesser) Guess(path string) (time.Time, error) {
	var err error
	for _, g := range guesser.guessers {
		time, err := g.Guess(path)
		if err == nil {
			return time, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to guess time: %v", err)
}
