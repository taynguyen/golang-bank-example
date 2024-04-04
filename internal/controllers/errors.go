package controllers

import "errors"

var (
	ErrUnauthorized = errors.New("user not authorized")
)
