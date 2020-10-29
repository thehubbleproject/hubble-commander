package core

// Result mirrors Types.Result in the contracts
type Result int

const (
	Ok Result = iota
	InvalidTokenAddress
	InvalidTokenAmount
	NotEnoughTokenBalance
	BadFromTokenType
	BadToTokenType
	BadFromIndex
	NotOnDesignatedStateLeaf
	BadSignature
	MismatchedAmount
	BadWithdrawRoot
)

// ParseResult parses a uint8 result returned by the contract, and returns a proper golang error
func ParseResult(result uint8) error {
	var r = Result(result)
	switch r {
	case Ok:
		return nil
	case InvalidTokenAddress:
		return ErrInvalidTokenAddress
	case InvalidTokenAmount:
		return ErrInvalidTokenAmount
	case NotEnoughTokenBalance:
		return ErrNotEnoughTokenBalance
	case BadFromTokenType:
		return ErrBadFromTokenType
	case BadToTokenType:
		return ErrBadToTokenType
	case BadFromIndex:
		return ErrBadFromIndex
	case NotOnDesignatedStateLeaf:
		return ErrNotOnDesignatedStateLeaf
	case BadSignature:
		return ErrBadSignature
	case MismatchedAmount:
		return ErrMismatchedAmount
	case BadWithdrawRoot:
		return ErrBadWithdrawRoot
	default:
		return GenericError("Undefined error")
	}
}
