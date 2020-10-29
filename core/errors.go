package core

import (
	"errors"
	"fmt"
)

// corresponding errors to Types.Results
var ErrInvalidTokenAddress = errors.New("invalid token address")
var ErrInvalidTokenAmount = errors.New("invalid token amount")
var ErrNotEnoughTokenBalance = errors.New("not enough token balance")
var ErrBadFromTokenType = errors.New("bad sender token type")
var ErrBadToTokenType = errors.New("bad receiver token type")
var ErrBadFromIndex = errors.New("bad sender state index")
var ErrNotOnDesignatedStateLeaf = errors.New("Not on designated state leaf")
var ErrBadSignature = errors.New("bad signature")
var ErrMismatchedAmount = errors.New("mismatched amount")
var ErrBadWithdrawRoot = errors.New("bad withdraw root")

func ErrRecordNotFound(msg string) error {
	return fmt.Errorf("Error: Record not found. Msg: %s", msg)
}

func GenericError(msg string) error {
	return fmt.Errorf("Error: %v", msg)
}

func ErrUnableToCreateRecord(msg string) error {
	return fmt.Errorf("Error: Unable to crete record. Msg: %s", msg)
}
