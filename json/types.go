package json

import (
	"github.com/KonstantinAfonin/CalculatorLambdaHandler/util"
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

func (response *CalculateResponse) setResult(result *big.Float) (err error) {
	var floatNumber float64
	floatNumber, err = util.BigFloatToFloat64(result)

	if err == nil {
		response.Result = floatNumber
	}

	return
}
