package json

import "github.com/KonstantinAfonin/CalculatorLambdaHandler/calculator"

func Calculate(request *CalculateRequest) (response *CalculateResponse, err error) {

	err = request.Validate()
	if err != nil {
		return
	}

	operation := request.GetOperation()
	numberA := request.GetNumberA()
	numberB := request.GetNumberB()

	bigFloatResult, err := calculator.Calculate(operation, numberA, numberB)

	if err == nil {
		response = new(CalculateResponse)
		err = response.setResult(bigFloatResult)
	}

	return
}
