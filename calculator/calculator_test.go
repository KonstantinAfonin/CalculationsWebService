package calculator

import (
	"errors"
	"fmt"
	"math/big"
	"testing"
)

func TestCalculatorNormalExecution(t *testing.T) {
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
		{"POW", "0", "0", "<nil>", errors.New("Invalid operation POW")},
	}

	for _, c := range cases {
		actual, err := Calculate(c.operation, float(c.numberA), float(c.numberB))

		if !equalResults(float(c.expected), actual) || !equalErrors(c.expectedErr, err) {
			t.Errorf("Invalid output on Calculate(%v, %v, %v): %v, %v, while expected %v, %v",
				c.operation, c.numberA, c.numberB, actual, err, c.expected, c.expectedErr)
		}
	}
}

func TestCalculatorPanics(t *testing.T) {
	cases := []struct {
		operation, numberA, numberB, expectedErr string
	}{
		{ADD, "-Inf", "+Inf", "addition of infinities with opposite signs"},
		{SUBSTRACT, "-Inf", "-Inf", "subtraction of infinities with equal signs"},
		{MULTIPLY, "0", "+Inf", "multiplication of zero with infinity"},
		{DIVIDE, "+Inf", "+Inf", "division of zero by zero or infinity by infinity"},
	}

	for _, c := range cases {
		actualErr := panicMessage(c.operation, c.numberA, c.numberB)

		if actualErr != c.expectedErr {
			t.Errorf("expected error: \"%v\", actual error: \"%v\"", c.expectedErr, actualErr)
		}
	}
}

func panicMessage(operation, numberA, numberB string) (err string) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Sprintf("%v", r)
		}
	}()

	Calculate(operation, float(numberA), float(numberB))

	return
}

func equalResults(expected, actual *big.Float) bool {
	if expected == nil || actual == nil {
		return expected == actual
	} else {
		return expected.Cmp(actual) == 0
	}
}

func equalErrors(first, second error) bool {
	if first == nil || second == nil {
		return first == second
	} else {
		return first.(error).Error() == second.(error).Error()
	}
}

func float(from string) *big.Float {
	f, _, _ := new(big.Float).Parse(from, 0)
	return f
}
