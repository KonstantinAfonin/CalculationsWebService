package json

import (
	"errors"
	"github.com/KonstantinAfonin/CalculatorLambdaHandler/util"
	"testing"
)

func TestCalculateRequestUnmarshal(t *testing.T) {

	cases := []struct {
		jsonString, expectedOperation, expectedNumberA, expectedNumberB string
	}{
		{"", "", "0", "0"},
		{"{\"number_a\": 2,\"number_b\": 8,\"operation\": \"ADD\"}", "ADD", "2", "8"},
		{"{\"number_a\": 1.5,\"number_b\": 2.5,\"operation\": \"  DiViDe  \"}", "DIVIDE", "1.5", "2.5"},
	}

	for _, c := range cases {

		request := unmarshalJson(c.jsonString)

		numberA := request.GetNumberA()
		numberB := request.GetNumberB()
		operation := request.GetOperation()

		wrongNumberA := !util.EqualBigFloats(numberA, util.StringToBigFloat(c.expectedNumberA))
		wrongNumberB := !util.EqualBigFloats(numberB, util.StringToBigFloat(c.expectedNumberB))
		wrongOperation := operation != c.expectedOperation

		if wrongNumberA || wrongNumberB || wrongOperation {
			t.Errorf("Invalid output for \"%v\" JSON. Expected: \"%v\", %v, %v; Actual: \"%v\", %v, %v",
				c.jsonString,
				c.expectedOperation, c.expectedNumberA, c.expectedNumberB,
				operation, numberA, numberB)
		}
	}
}

func TestCalculateResponseMarshal(t *testing.T) {
	cases := []struct {
		result, expectedJson string
		expectedError        error
	}{
		{"+Inf", "{\"result\":0}", errors.New("+Inf value doesn't fit into JSON number type")},
		{"1e+500", "{\"result\":0}", errors.New("1e+500 value doesn't fit into float64 type")},
		{"100", "{\"result\":100}", nil},
		{"100.5", "{\"result\":100.5}", nil},
	}

	for _, c := range cases {
		response := new(CalculateResponse)
		err := response.setResult(util.StringToBigFloat(c.result))
		actualJson := marshalJson(response)

		if !util.EqualErrors(c.expectedError, err) || actualJson != c.expectedJson {
			t.Errorf("Invalid output for '%v' result. Expected JSON: '%v', expected error: '%v', actual JSON: '%v', actual error: '%v'",
				c.result, c.expectedJson, c.expectedError, actualJson, err)
		}
	}
}
