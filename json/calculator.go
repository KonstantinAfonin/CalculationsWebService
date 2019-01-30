package json

import "github.com/KonstantinAfonin/CalculationsWebService/calculator"

func Calculate(request *CalculateRequest) (response *CalculateResponse, err error) {

	operation := request.GetOperation()
	numberA := request.GetNumberA()
	numberB := request.GetNumberB()

	response = new(CalculateResponse)

	bigFloatResult, err := calculator.Calculate(operation, numberA, numberB)

	if err == nil {
		err = response.setResult(bigFloatResult)
	}

	return
}
