package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCalculateRequestUnmarshal(t *testing.T) {

	jsonString := "{\"number_a\": 2,\"number_b\": 8,\"operation\": \"ADD\"}"

	request := new(CalculateRequest)

	_ = json.Unmarshal([]byte(jsonString), request)

	fmt.Printf("request: %v, getOperation: %v", request, request.GetOperation())

}

func TestCalculateResponseMarshal(t *testing.T) {

	//small, _ := util.BigFloat("+Inf").Float64()
	//fmt.Printf("small %v, isNaN %v", small, math.IsInf(small, 0))
	//
	////float := util.BigFloat("100")
	//response := CalculateResponse{small, ""}
	////response.setError(errors.New("some error"))
	//
	//bytes, _ := json.Marshal(response)
	//
	//fmt.Printf("response: \"%v\"", string(bytes))
}
