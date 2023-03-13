package service_error

import "fmt"

type ServiseErrorStruct struct {
	Code    string `json:"code"`
	Message any    `json:"msg"`
}

func (e ServiseErrorStruct) Error() string {
	return fmt.Sprintf("Service Error: %s (code: %s)", e.Message, e.Code)
}

func ServiceError(code string, message any) *ServiseErrorStruct {
	return &ServiseErrorStruct{
		Code:    code,
		Message: message,
	}
}
