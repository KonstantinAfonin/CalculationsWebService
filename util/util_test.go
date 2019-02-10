package util

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestBigFloatToFloat64(t *testing.T) {
	cases := []struct {
		input          *big.Float
		expectedOutput float64
		expectedError  error
	}{
		{StringToBigFloat("100"), 100, nil},
		{StringToBigFloat("-100.5"), -100.5, nil},

		{StringToBigFloat("-Inf"), math.Inf(-1), nil},
		{StringToBigFloat("+Inf"), math.Inf(1), nil},
		{StringToBigFloat("1e+500"), math.Inf(1), errors.New("1e+500 value doesn't fit float64 type")},
	}

	for _, c := range cases {
		actualOutput, actualError := BigFloatToFloat64(c.input)
		if actualOutput != c.expectedOutput || !EqualErrors(actualError, c.expectedError) {
			fmt.Printf(`Invalid output for %v input. Expected: output %v, error %v. Actual: output %v, error %v`,
				c.input, c.expectedOutput, c.expectedError, actualOutput, actualError)
		}
	}
}
