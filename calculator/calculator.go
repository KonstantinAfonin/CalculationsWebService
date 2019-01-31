package calculator

import (
	"errors"
	"fmt"
	"math/big"
)

const (
	ADD      = "ADD"
	DIVIDE   = "DIVIDE"
	SUBTRACT = "SUBTRACT"
	MULTIPLY = "MULTIPLY"
)

func Calculate(operation string, numberA *big.Float, numberB *big.Float) (result *big.Float, err error) {
	defer func() {
		if r := recover(); r != nil {
			result, err = nil, errors.New(fmt.Sprint(r))
		}
	}()

	switch operation {
	case ADD:
		result, err = new(big.Float).Add(numberA, numberB), nil
	case SUBTRACT:
		result, err = new(big.Float).Sub(numberA, numberB), nil
	case MULTIPLY:
		result, err = new(big.Float).Mul(numberA, numberB), nil
	case DIVIDE:
		result, err = new(big.Float).Quo(numberA, numberB), nil
	default:
		result, err = nil, errors.New(fmt.Sprintf("invalid operation '%v'", operation))
	}

	return
}
