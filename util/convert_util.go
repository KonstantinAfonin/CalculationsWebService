package util

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

func StringToBigFloat(from string) *big.Float {
	f, _, _ := new(big.Float).Parse(from, 0)
	return f
}

func Float64ToBigFloat(number float64) *big.Float {
	return new(big.Float).SetFloat64(number)
}

func BigFloatToFloat64(from *big.Float) (result float64, err error) {

	if from.IsInf() {
		result = math.Inf(from.Sign())
		return
	}

	result, _ = from.Float64()
	if math.IsInf(result, 0) {
		err = errors.New(fmt.Sprintf("%v value doesn't fit into float64 type", from))
	}

	return
}

func EqualBigFloats(expected, actual *big.Float) bool {
	if expected == nil || actual == nil {
		return expected == actual
	} else {
		return expected.Cmp(actual) == 0
	}
}

func EqualErrors(first, second error) bool {
	if first == nil || second == nil {
		return first == second
	} else {
		return first.(error).Error() == second.(error).Error()
	}
}
