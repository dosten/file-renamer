package util

type Renamer interface {
	Rename(old, new string) error
}
