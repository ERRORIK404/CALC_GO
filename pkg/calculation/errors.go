package calculation

import (
	"errors"
)

var (
    ErrEmptyExpression        = errors.New("empty expression")
    ErrIncorrectBracketPlacement = errors.New("incorrect placement of brackets")
    ErrBracketMismatch         = errors.New("bracket mismatch")
    ErrInvalidCharacter       = errors.New("invalid character")
    ErrNotEnoughOperands      = errors.New("not enough operands")
    ErrIncorrectExpression    = errors.New("incorrect expression")
    ErrUnknownOperation       = errors.New("unknown operation")
    ErrDivisionByZero         = errors.New("division by zero")
)