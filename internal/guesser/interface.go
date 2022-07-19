package guesser

import (
	"time"
)

type TimeGuesser interface {
	Guess(string) (time.Time, error)
}
