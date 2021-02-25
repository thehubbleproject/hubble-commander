package common

import (
	"fmt"
)

const (
	mantissaBits uint16 = 12
	exponentBits uint16 = 4
	exponentMax         = 1<<exponentBits - 1
	mantissaMax         = 1<<mantissaBits - 1
)

// Compress compress an unsinged 64 bits integer input to float16 2 bytes
func Compress(input uint64) (output [2]byte, err error) {
	mantissa := input
	exponent := uint16(0)
	for exponent < exponentMax {
		if mantissa == 0 || mantissa%10 != 0 {
			break
		}
		mantissa /= 10
		exponent++
	}

	if mantissa > mantissaMax {
		return output, fmt.Errorf("mantissa %d is greater than mantissaMax %d", mantissa, mantissaMax)
	}
	combined := (exponent << mantissaBits) | uint16(mantissa)
	// combined is a little endian golang type
	output[0] = byte(combined >> 8)
	output[1] = byte(combined)
	return
}

// Decompress decompress a float16 2 bytes to an unsinged 64 bits integer
func Decompress(input [2]byte) (output uint64) {
	mantissa := uint16(input[0]&0x0F)<<8 | uint16(input[1])
	exponent := (input[0] & 0xF0) >> 4
	result := uint64(mantissa)
	for i := uint8(0); i < exponent; i++ {
		result *= 10
	}
	return result
}

// Round converts an input to the largest possible compressible integer
func Round(input uint64) (output uint64, err error) {
	mantissa := input
	exponent := uint16(0)
	powOfTen := uint64(1)
	for exponent < exponentMax {
		if mantissa <= mantissaMax {
			return mantissa * powOfTen, nil
		}
		mantissa /= 10
		powOfTen *= 10
		exponent++
	}
	return 0, fmt.Errorf("Can't round the input %d", input)
}
