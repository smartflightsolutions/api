package api

// Response for successful requests
type Response struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

// ErrorResponse default response to unspecified errors
type ErrorResponse struct {
	Error string `json:"error"`
}

// UnprocessableEntityErrResponse default response to unspecified errors
type UnprocessableEntityErrResponse struct {
	ErrorResponse
	InvalidAttributes interface{} `json:"invalid_attributes"`
}