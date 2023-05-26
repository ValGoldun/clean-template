package errors

import "errors"

type Error struct {
	text       string
	isCritical bool
}

func New(text string) error {
	return errors.New(text)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func (e Error) Error() string {
	return e.text
}

func (e Error) IsCritical() bool {
	return e.isCritical
}
