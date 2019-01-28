package calculator

import (
	"errors"
	"fmt"
	"math/big"
)

const (
	ADD       = "ADD"
	DIVIDE    = "DIVIDE"
	SUBSTRACT = "SUBSTRACT"
	MULTIPLY  = "MULTIPLY"
)

func Calculate(operation string, numberA *big.Float, numberB *big.Float) (result *big.Float, err error) {
	switch operation {
	case ADD:
		result, err = new(big.Float).Add(numberA, numberB), nil
	case SUBSTRACT:
		result, err = new(big.Float).Sub(numberA, numberB), nil
	case MULTIPLY:
		result, err = new(big.Float).Mul(numberA, numberB), nil
	case DIVIDE:
		result, err = new(big.Float).Quo(numberA, numberB), nil
	default:
		result, err = nil, errors.New(fmt.Sprintf("Invalid operation %v", operation))
	}

	return
}