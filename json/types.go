package json

import (
	"errors"
	"fmt"
	"github.com/KonstantinAfonin/CalculatorLambdaHandler/util"
	"math"
	"math/big"
	"strings"
)

type CalculateRequest struct {
	NumberA   float64 `json:"number_a"`
	NumberB   float64 `json:"number_b"`
	Operation string  `json:"operation"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func (request *CalculateRequest) GetOperation() string {
	return strings.TrimSpace(strings.ToUpper(request.Operation))
}

func (request *CalculateRequest) GetNumberA() *big.Float {
	return util.Float64ToBigFloat(request.NumberA)
}

func (request *CalculateRequest) GetNumberB() *big.Float {
	return util.Float64ToBigFloat(request.NumberB)
}

func (response *CalculateResponse) setResult(result *big.Float) error {

	floatNumber, err := util.BigFloatToFloat64(result)
	if err != nil {
		return err
	}

	if math.IsInf(floatNumber, 0) {
		return errors.New(fmt.Sprintf("%v value doesn't fit JSON into number type", floatNumber))
	}

	response.Result = floatNumber
	return nil
}
