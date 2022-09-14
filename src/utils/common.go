package utils

import (
	"gin-api/src/models"
)

func FactorialNumber(param int64) (result int64, err error) {
	result = 1
	if param == 1 || param == 0 {
		return
	}

	data, err := FactorialNumber(param - 1)
	result = result * data
	return
}

var ReadJsonFile = func(filepath string) (*byte, error) {
	return nil, nil
}

func HttpResponse(code int, message string, data interface{}) models.WebResponse {
	response := models.WebResponse{
		Code:    code,
		Message: message,
		Payload: data,
	}
	return response
}
