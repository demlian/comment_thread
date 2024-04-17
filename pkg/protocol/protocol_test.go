package protocol

import (
	"fmt"
	"testing"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected *Request
		err      error
	}{
		{
			name: "ValidRequest",
			data: "okcypwj|VERB|Data",
			expected: &Request{
				requestId: "okcypwj",
				verb:      "VERB",
				data:      "Data",
			},
			err: nil,
		},
		{
			name:     "InvalidRequestSeparator",
			data:     "invalid-data",
			expected: nil,
			err:      fmt.Errorf("received invalid data format: invalid-data"),
		},
		{
			name:     "InvalidRequestIdLEngth",
			data:     "request-id|verb|data",
			expected: nil,
			err:      fmt.Errorf("received invalid request ID format: request-id"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := newRequest(tt.data)
			if err != nil && tt.err == nil {
				t.Errorf("newRequest returned an unexpected error: %v; expected: %+v", err, tt.expected)
			}
			if err == nil && tt.err != nil {
				t.Errorf("newRequest returned the following request: %+v; expected error: %v", request, tt.err)
			}
			if request != nil && tt.expected != nil {
				if request.requestId != tt.expected.requestId {
					t.Errorf("newRequest returned incorrect requestId. Expected: %s, Got: %s", tt.expected.requestId, request.requestId)
				}
				if request.verb != tt.expected.verb {
					t.Errorf("newRequest returned incorrect verb. Expected: %s, Got: %s", tt.expected.verb, request.verb)
				}
				if request.data != tt.expected.data {
					t.Errorf("newRequest returned incorrect data. Expected: %s, Got: %s", tt.expected.data, request.data)
				}
			}
		})
	}
}
