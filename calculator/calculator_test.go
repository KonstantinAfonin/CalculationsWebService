package calculator

import (
	"errors"
	"github.com/KonstantinAfonin/CalculationsWebService/util"
	"testing"
)

func TestCalculate(t *testing.T) {
	cases := []struct {
		operation, numberA, numberB, expected string
		expectedErr                           error
	}{
		{ADD, "100", "200", "300", nil},
		{ADD, "100.5", "200.22", "300.72", nil},
		{ADD, "+Inf", "+Inf", "+Inf", nil},
		{SUBSTRACT, "100", "100", "0", nil},
		{SUBSTRACT, "100", "100.5", "-0.5", nil},
		{SUBSTRACT, "+Inf", "5", "+Inf", nil},
		{MULTIPLY, "1e+500", "1e+500", "1e+1000", nil},
		{MULTIPLY, "100.5", "0.5", "50.25", nil},
		{MULTIPLY, "+Inf", "5", "+Inf", nil},
		{DIVIDE, "100", "4", "25", nil},
		{DIVIDE, "1.2", "3", "0.4", nil},
		{DIVIDE, "10", "0", "+Inf", nil},
		{"POW", "0", "0", "<nil>", errors.New("invalid operation \"POW\"")},
		{ADD, "-Inf", "+Inf", "<nil>", errors.New("addition of infinities with opposite signs")},
		{SUBSTRACT, "-Inf", "-Inf", "<nil>", errors.New("subtraction of infinities with equal signs")},
		{MULTIPLY, "0", "+Inf", "<nil>", errors.New("multiplication of zero with infinity")},
		{DIVIDE, "+Inf", "+Inf", "<nil>", errors.New("division of zero by zero or infinity by infinity")},
	}

	for _, c := range cases {
		actual, err := Calculate(c.operation, util.StringToBigFloat(c.numberA), util.StringToBigFloat(c.numberB))

		if !util.EqualBigFloats(util.StringToBigFloat(c.expected), actual) || !util.EqualErrors(c.expectedErr, err) {
			t.Errorf("Invalid output on Calculate(%v, %v, %v): %v, %v, while expected %v, %v",
				c.operation, c.numberA, c.numberB, actual, err, c.expected, c.expectedErr)
		}
	}
}
