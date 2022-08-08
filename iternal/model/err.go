package model

import "errors"

var (
	ErrWeekNotCreated         = errors.New("training week is not in the database")
	ErrWeekAlreadyCreated     = errors.New("the week is already in the database")
	ErrIncorrectDaysOfWeek    = errors.New("incorrect names of the days of the week")
	ErrNoEqualSeven           = errors.New("wrong number of days passed, expected 7")
	ErrCharacterLimitExceeded = errors.New("description cannot be more than 50 characters")
	ErrNotEnoughCharacters    = errors.New("description must be at least 2 characters long")
)

type Err struct {
	Status int
	Error  error
}
