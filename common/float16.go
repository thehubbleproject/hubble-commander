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
	output[0] = byte(combined >> 8)
	output[1] = byte(combined)
	return
}

func Decompress() {

}

func Round() {

}
