package calculator

import "errors"

var (
	ErrMismatchedParentheses = errors.New("mismatched parentheses")
	ErrInvalidExpression      = errors.New("invalid expression")
	ErrDivisionByZero        = errors.New("division by zero")
	ErrInvalidNumber          = errors.New("invalid number")
	ErrInternalServerError    = errors.New("internal server error")
)
