package model

import "errors"

//calendar errors
var (
	ErrWeekNotCreated         = errors.New("training week is not in the database")
	ErrWeekAlreadyCreated     = errors.New("the week is already in the database")
	ErrIncorrectDaysOfWeek    = errors.New("incorrect names of the days of the week")
	ErrNoEqualSeven           = errors.New("wrong number of days passed, expected 7")
	ErrCharacterLimitExceeded = errors.New("description cannot be more than 50 characters")
)

//user errors
var (
	ErrUserNotFound          = errors.New("user by id was not found")
	ErrWrongEmailOrPassword  = errors.New("wrong email or password")
	ErrIDLessOne             = errors.New("id cannot be less than one")
	ErrWrongToken            = errors.New("wrong token")
	ErrEmailIsBusy           = errors.New("a user with this email has already been created")
	ErrEmailPasswordNotValid = errors.New("email and/or password do not meet the requirements")
)

//news errors
var (
	ErrWrongContentType         = errors.New("invalid file type, expected png or jpeg")
	ErrTitleDescriptionNotValid = errors.New("title cannot be empty and must not exceed 50 characters, and description 1500 characters")
	ErrLongFileName             = errors.New("character limit exceeded in file name")
)

type Err struct {
	Status int
	Error  error
}
