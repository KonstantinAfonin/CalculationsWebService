package json

import (
	"errors"
	"fmt"
	"github.com/KonstantinAfonin/CalculatorLambdaHandler/util"
	"math"
	"math/big"
	"reflect"
	"strings"
)

const (
	number_a  = "number_a"
	number_b  = "number_b"
	operation = "operation"
)

type CalculateRequest map[string]interface{}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func (request *CalculateRequest) Validate() (err error) {

	var validationErrors = make([]string, 0, 0)
	var missingFields = make([]string, 0, 0)
	var extraFields = make([]string, 0, 0)

	underlyingMap := map[string]interface{}(*request)

	validateField(number_a, underlyingMap[number_a], reflect.Float64, &missingFields, &validationErrors)
	validateField(number_b, underlyingMap[number_b], reflect.Float64, &missingFields, &validationErrors)
	validateField(operation, underlyingMap[operation], reflect.String, &missingFields, &validationErrors)

	for key := range underlyingMap {
		if key != number_a && key != number_b && key != operation {
			extraFields = append(extraFields, key)
		}
	}

	if len(missingFields) != 0 {
		validationErrors = append(validationErrors, "missing fields: "+strings.Join(missingFields, ", "))
	}

	if len(extraFields) != 0 {
		validationErrors = append(validationErrors, "extra fields: "+strings.Join(extraFields, ", "))
	}

	if len(validationErrors) == 0 {
		err = nil
	} else {
		err = errors.New("invalid JSON provided: " + strings.Join(validationErrors, "; "))
	}

	return
}

func validateField(name string, value interface{}, expectedType reflect.Kind, missingFields *[]string, validationErrors *[]string) {
	if value != nil {
		valueKind := reflect.TypeOf(value).Kind()
		if valueKind != expectedType {
			*validationErrors = append(*validationErrors, fmt.Sprintf("invalid '%v' type: %v", name, valueKind))
		}
	} else {
		*missingFields = append(*missingFields, name)
	}
}

func (request *CalculateRequest) GetOperation() string {
	return strings.TrimSpace(strings.ToUpper(request.getField(operation).(string)))
}

func (request *CalculateRequest) GetNumberA() *big.Float {
	return util.Float64ToBigFloat(request.getField(number_a).(float64))
}

func (request *CalculateRequest) GetNumberB() *big.Float {
	return util.Float64ToBigFloat(request.getField(number_b).(float64))
}

func (request *CalculateRequest) getField(fieldName string) interface{} {
	underlyingMap := map[string]interface{}(*request)
	return underlyingMap[fieldName]
}

func (response *CalculateResponse) setResult(result *big.Float) error {

	floatNumber, err := util.BigFloatToFloat64(result)
	if err != nil {
		return err
	}

	if math.IsInf(floatNumber, 0) {
		return errors.New(fmt.Sprintf("%v value doesn't fit into JSON number type", floatNumber))
	}

	response.Result = floatNumber
	return nil
}
