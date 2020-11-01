package core

import (
	"errors"
	"fmt"
)

// corresponding errors to Types.Results
var ErrInvalidTokenAmount = errors.New("invalid token amount")
var ErrNotEnoughTokenBalance = errors.New("not enough token balance")
var ErrBadFromTokenType = errors.New("bad sender token type")
var ErrBadToTokenType = errors.New("bad receiver token type")
var ErrBadSignature = errors.New("bad signature")
var ErrMismatchedAmount = errors.New("mismatched amount")
var ErrBadWithdrawRoot = errors.New("bad withdraw root")
var ErrBadCompression = errors.New("bad transaction compression")
var ErrTooManyTx = errors.New("too many transactions in a commitment")

func ErrRecordNotFound(msg string) error {
	return fmt.Errorf("Error: Record not found. Msg: %s", msg)
}

func GenericError(msg string) error {
	return fmt.Errorf("Error: %v", msg)
}

func ErrUnableToCreateRecord(msg string) error {
	return fmt.Errorf("Error: Unable to crete record. Msg: %s", msg)
}
