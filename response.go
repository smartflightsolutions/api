package api

// NewNotFoundResponse returns a response with message not found
func NewNotFoundResponse() *ErrorResponse {
	r := &ErrorResponse{}
	r.Error = "Object not found"
	return r
}

// NewResponse return a default response for successful
func NewResponse(data interface{}, meta ...interface{}) *Response {
	response := &Response{
		Data: data,
		Meta: meta,
	}
	return response
}

// NewErrorResponse return unsuccessful response
func NewErrorResponse(errorMessage string) *ErrorResponse {
	response := &ErrorResponse{
		Error: errorMessage,
	}
	return response
}

// NewUnprocessableEntityResponse return unprocessable entity response
func NewUnprocessableEntityResponse(errorMessage string, invalidAttributes interface{}) *UnprocessableEntityErrResponse {
	response := &UnprocessableEntityErrResponse{}
	response.Error =  errorMessage
	response.InvalidAttributes = invalidAttributes
	return response
}