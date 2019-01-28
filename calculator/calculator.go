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

func Calculate(operation string, operandA *big.Float, operandB *big.Float) (result *big.Float, err error) {
	switch operation {
	case ADD:
		result, err = new(big.Float).Add(operandA, operandB), nil
	case SUBSTRACT:
		result, err = new(big.Float).Sub(operandA, operandB), nil
	case MULTIPLY:
		result, err = new(big.Float).Mul(operandA, operandB), nil
	case DIVIDE:
		result, err = new(big.Float).Quo(operandA, operandB), nil
	default:
		result, err = nil, errors.New(fmt.Sprintf("Invalid operation %v", operation))
	}

	return
}
