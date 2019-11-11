package api

import "testing"

func Test_NewNotFoundResponse(t *testing.T) {
	response := NewNotFoundResponse()
	if response.Error != "Object not found" {
		t.Errorf("Expected \"Object not found\" but got %s", response.Error)
	}
}

func Test_NewResponse(t *testing.T) {
	response := NewResponse("data should be here", "meta 1", "meta 2")
	if response.Data != "data should be here" {
		t.Errorf("Expected \"data should be here\" but got %s", response.Data)
	}
	if len(response.Meta.([]interface{})) != 2 {
		t.Errorf("Expected len 1 but got %d", len(response.Meta.([]interface{})))
	}
}

func Test_NewErrorResponse(t *testing.T) {
	response := NewErrorResponse("error message should be here")
	if response.Error != "error message should be here" {
		t.Errorf("Expected \"error message should be here\" but got %s", response.Error)
	}
}

func Test_NewUnprocessableEntityResponse(t *testing.T) {
	response := NewUnprocessableEntityResponse("error message should be here", []string{"validation 1", "validation 2"})
	if response.Error != "error message should be here" {
		t.Errorf("Expected \"error message should be here\" but got %s", response.Error)
	}
	if len(response.InvalidAttributes.([]string)) != 2 {
		t.Errorf("Expected len 1 but got %d", len(response.InvalidAttributes.([]string)))
	}
}
