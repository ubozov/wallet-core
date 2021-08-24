package dto

import (
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

// BaseDto ...
type BaseDto struct {
	Success      bool     `json:"success"`
	FullMessages []string `json:"full_messages"`
}

// ErrorDto ...
type ErrorDto struct {
	BaseDto
	Errors map[string]interface{} `json:"errors"`
}

// This should only be called when we have an Error that is returned from a ShouldBind which contains a lot of information
// other kind of errors should use other functions such as CreateDetailedErrorDto
func CreateBadRequestErrorDto(err error) ErrorDto {
	res := ErrorDto{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	res.FullMessages = make([]string, len(errs))
	count := 0
	for _, v := range errs {
		if v.ActualTag == "required" {
			var message = fmt.Sprintf("%v is required", v.Field)
			res.Errors[v.Field] = message
			res.FullMessages[count] = message
		} else {
			var message = fmt.Sprintf("%v has to be %v", v.Field, v.ActualTag)
			res.Errors[v.Field] = message
			res.FullMessages = append(res.FullMessages, message)
		}
		count++
	}
	return res
}

// CreateSuccessWithDtoAndMessageDto ...
func CreateSuccessWithDtoAndMessageDto(data interface{}, messages []string) map[string]interface{} {
	result := make(map[string]interface{})
	result["data"] = data
	result["success"] = true
	result["full_messages"] = messages
	return result
}
