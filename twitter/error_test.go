package twitter

import (
	"fmt"
	"reflect"
	"testing"
)

var testAPIError = APIError{
	Errors: []ErrorDetail{
		ErrorDetail{Message: "Status is a duplicate", Code: 187},
	},
}

func TestApiError_Empty(t *testing.T) {
	err := APIError{}
	if !err.Empty() {
		t.Errorf("expected Empty() to return true for %v", err)
	}

	err = APIError{
		Errors: []ErrorDetail{
			ErrorDetail{Message: "Status is a duplicate", Code: 187},
		},
	}
	if err.Empty() {
		t.Errorf("expected Empty() to return false for %v", err)
	}
}

func TestRelevantError(t *testing.T) {
	cases := []struct {
		httpError error
		apiError  APIError
		expected  error
	}{
		{nil, APIError{}, nil},
		{nil, testAPIError, testAPIError},
		{fmt.Errorf("unknown host"), APIError{}, fmt.Errorf("unknown host")},
		{fmt.Errorf("unknown host"), testAPIError, fmt.Errorf("unknown host")},
	}
	for _, c := range cases {
		err := relevantError(c.httpError, c.apiError)
		if !reflect.DeepEqual(c.expected, err) {
			t.Errorf("not DeepEqual: expected %v, got %v", c.expected, err)
		}
	}
}
