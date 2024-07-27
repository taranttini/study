package entity

type CustomError struct {
	Message string
	Code    int
}

func ErrorZipcodeInvalid() *CustomError {
	return &CustomError{
		Message: "invalid zipcode",
		Code:    422,
	}
}

func ErrorZipcodeNotFound() *CustomError {
	return &CustomError{
		Message: "can not find zipcode",
		Code:    404,
	}
}

func ErrorInternal() *CustomError {
	return &CustomError{
		Message: "internal error",
		Code:    500,
	}
}
