package core

// Result mirrors Types.Result in the contracts
type Result int

const (
	Ok Result = iota
	InvalidTokenAmount
	NotEnoughTokenBalance
	BadFromTokenType
	BadToTokenType
	BadSignature
	MismatchedAmount
	BadWithdrawRoot
	BadCompression
	TooManyTx
)

// ParseResult parses a uint8 result returned by the contract, and returns a proper golang error
func ParseResult(result uint8) error {
	var r = Result(result)
	switch r {
	case Ok:
		return nil
	case InvalidTokenAmount:
		return ErrInvalidTokenAmount
	case NotEnoughTokenBalance:
		return ErrNotEnoughTokenBalance
	case BadFromTokenType:
		return ErrBadFromTokenType
	case BadToTokenType:
		return ErrBadToTokenType
	case BadSignature:
		return ErrBadSignature
	case MismatchedAmount:
		return ErrMismatchedAmount
	case BadWithdrawRoot:
		return ErrBadWithdrawRoot
	case BadCompression:
		return ErrBadCompression
	case TooManyTx:
		return ErrTooManyTx
	default:
		return GenericError("Undefined error")
	}
}
