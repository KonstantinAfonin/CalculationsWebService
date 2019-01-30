package json

import (
	"encoding/json"
	"github.com/KonstantinAfonin/CalculationsWebService/util"
	"testing"
)

func TestCalculateRequest(t *testing.T) {

	cases := []struct {
		jsonString, expectedOperation, expectedNumberA, expectedNumberB string
	}{
		{"", "", "0", "0"},
		{"{\"number_a\": 2,\"number_b\": 8,\"operation\": \"ADD\"}", "ADD", "2", "8"},
		{"{\"number_a\": 1.5,\"number_b\": 2.5,\"operation\": \"  DiViDe  \"}", "DIVIDE", "1.5", "2.5"},
	}

	for _, c := range cases {
		request := new(CalculateRequest)
		_ = json.Unmarshal([]byte(c.jsonString), request)

		requestA := request.GetNumberA()
		requestB := request.GetNumberB()
		requestOperation := request.GetOperation()

		wrongNumberA := !util.EqualBigFloats(requestA, util.StringToBigFloat(c.expectedNumberA))
		wrongNumberB := !util.EqualBigFloats(requestB, util.StringToBigFloat(c.expectedNumberB))
		wrongOperation := requestOperation != c.expectedOperation

		if wrongNumberA || wrongNumberB || wrongOperation {
			t.Errorf("Invalid output for \"%v\" JSON. Expected: \"%v\", %v, %v; Actual: \"%v\", %v, %v",
				c.jsonString,
				c.expectedOperation, c.expectedNumberA, c.expectedNumberB,
				requestOperation, requestA, requestB)
		}
	}
}

func TestCalculateResponse(t *testing.T) {
	//cases := []struct {
	//	result, expectedJson string
	//	expectedError        error
	//}{
	//	{},
	//}
	//
	//for _, c := range cases {
	//
	//	response := new(CalculateResponse)
	//	err := response.setResult(util.StringToBigFloat(c.result))
	//	bytes, _ := json.Marshal(response)
	//
	//	if (!util.EqualErrors(c.expectedError, err) || )
	//}
}
