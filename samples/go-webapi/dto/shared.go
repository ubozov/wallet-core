package dto

// BaseDto ...
type BaseDto struct {
	Success      bool     `json:"success"`
	FullMessages []string `json:"full_messages"`
}

// ErrorDto ...
type ErrorDto struct {
	BaseDto
	Error string `json:"errors"`
}

// This should only be called when we have an Error that is returned from a ShouldBind which contains a lot of information
// other kind of errors should use other functions such as CreateDetailedErrorDto
func CreateBadRequestErrorDto(err error) ErrorDto {
	res := ErrorDto{}
	res.Error = err.Error()
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
