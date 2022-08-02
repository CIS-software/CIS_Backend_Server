package model

import "errors"

var (
	ErrWeekNotCreated     = errors.New("training week is not in the database")
	ErrWeekAlreadyCreated = errors.New("the week is already in the database")
)

type Err struct {
	Status int
	Error  error
}
