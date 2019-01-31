## Calculator Lambda Handler

### Abstract

Simple API which allows to execute basic arithmetic operations (ADD, SUBTRACT, MULTIPLY and DIVIDE) 
over a pair of float numbers. 

Code is divided into layers (basic implementation -> JSON wrapper -> AWS Lambda Handler)
so that each layer can be used individually, regardless of the upper level's
implementation.

### Source Code Reference

#### `calculator` package

**Exposes `calculator.Calculate(string, *big.Float, *big.Float) (*big.Float, error)` function**,
which executes operation provided in the first `string` argument over the numbers, provided
as `*big.Float` parameters and returns result as another `big.Float` value.  

Supported operations are exposed as constants in the package.

`big.Float` was chosen as the type for parameters and result both because it can handle huge numbers 
and because it handles some situation better than primitives (such as division by zero
produces `+Inf` instead of panic as it would be with primitives, which seem more mathematically correct).
The obvious downside might be performance, yet in case of optimization the interface can be
changed with minimum changes of client code, as all the computation-related logic is encapsulated within
`calculator.Calculate` function.

Panics produced by `big.Float` methods are wrapped and returned as `error` so that the function
returns normally in any case.  

`calculator_test.go` file contains testing scenarios.

#### `json` package

**Exposes `CalculateRequest` and `CalculateResponse` structs**, which can be marshalled to and unmarhalled from 
JSON. 

`float64` was chosen as the type for input values and result in order for corresponding fields
to be numbers in JSON format. Methods of `CalculateRequest` and `CalculateResponse` encapsulate 
conversion between `big.Float` and `float64` and normalization of the `operation` argument 
(trim and upper case mapping). 

Also **exposes `json.Calculate(*CalculateRequest) (*CalculateResponse, error)` function** which 
wraps the call to `calculator.Calculate` function into JSON-related types.

`types_test.go` file contains tests for types' methods. `calculator_test.go` contains testing scenarions for `json.Calculate` function.

#### `aws_lambda` package

**Contains `lambda_handler` program** which can be used _AWS Lambda Function Handler_.
`HandleRequest(context.Context, json.CalculateRequest) (*json.CalculateResponse, error)` function delegates call
to `json.Calculate` function.

#### `util` package

Type conversion and comparison utils.