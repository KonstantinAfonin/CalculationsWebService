package json

import "encoding/json"

func marshalJson(response *CalculateResponse) string {
	actualJsonBytes, _ := json.Marshal(response)
	return string(actualJsonBytes)
}

func unmarshalJson(requestString string) (request *CalculateRequest) {
	request = new(CalculateRequest)
	_ = json.Unmarshal([]byte(requestString), request)
	return
}
