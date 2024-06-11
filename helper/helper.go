package helper

import "github.com/go-playground/validator/v10"

// response api dengan return meta dan data
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
	//menggunakan interface agar output field jadi flexible
}

// output status api
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidatationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
