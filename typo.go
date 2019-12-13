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

// messageBroker provides an interface so we can swap out the
// implementation of SendEmail under tests.
type messageBroker interface {
	send(topicID string, msg interface{}) error
}

// googlePubSub message broker
type googlePubSub struct {
	ProjectID string
}

// messageAPI standard message api
type messageAPI struct {
	API messageBroker
}
