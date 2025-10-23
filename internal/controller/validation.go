package controller

import (
	"errors"
)

func ValidateComment(text string) error {
	if len(text) == 0 || len(text) > 2000 {
		return errors.New("invalid comment length")
	}
	return nil
}

func ValidatePost(title, body string) error {
	if len(title) == 0 || len(title) > 300 {
		return errors.New("invalid title length")
	}
	if len(body) == 0 || len(body) > 2000 {
		return errors.New("invalid body length")
	}
	return nil
}

func ValidateUser(name string) error {
	if len(name) == 0 || len(name) > 100 {
		return errors.New("name must be between 1 and 100 characters")
	}
	return nil
}
