package dto

type ResponseError struct {
	Status bool     `json:"status"`
	Error  []string `json:"errors"`
}

func NewResponseError(status bool, errors []string) *ResponseError {
	return &ResponseError{
		Status: status,
		Error:  errors,
	}
}