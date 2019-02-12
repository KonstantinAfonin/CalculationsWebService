package json

import (
	"errors"
	"fmt"
	"github.com/KonstantinAfonin/CalculatorLambdaHandler/util"
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {

	cases := []struct {
		requestJson, expectedResponseJson string
		expectedError                     error
	}{
		{`{"number_a": 2,"number_b": 8,"operation": "ADD"}`, `{"result":10}`, nil},
		{`{"number_a": 2,"number_b": 8,"operation": "MULTIPLY"}`, `{"result":16}`, nil},
		{`{"number_a": 999.999,"number_b": 0.001,"operation": "    aDd    "}`, `{"result":1000}`, nil},

		{`{"number_a": 2,"number_b": "8","operation": "ADD"}`, `null`,
			errors.New("invalid JSON provided: invalid 'number_b' type: string")},
		{`{"number_a": 2,"number_b": 8,"operation": "POW"}`, `null`, errors.New("invalid operation 'POW'")},
		{"", "null", errors.New(`invalid JSON provided: missing fields: number_a, number_b, operation`)},
		{fmt.Sprintf(`{"number_a": %v,"number_b": 8,"operation": "MULTIPLY"}`, math.MaxFloat64), `{"result":0}`,
			errors.New("1.4381545078898526e+309 value doesn't fit into float64 type")},
	}

	for _, c := range cases {

		request := unmarshalJson(c.requestJson)

		response, err := Calculate(request)

		actualResponseJson := marshalJson(response)

		if c.expectedResponseJson != actualResponseJson || !util.EqualErrors(c.expectedError, err) {
			fmt.Printf("Invalid result. request: '%v', expected response: '%v', expected error: '%v', actual response: '%v', actual error: '%v'",
				c.requestJson, c.expectedResponseJson, c.expectedError, actualResponseJson, err)
		}
	}

}
