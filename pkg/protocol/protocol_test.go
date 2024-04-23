package protocol

import (
	"fmt"
	"testing"
)

func TestHandleAuthService(t *testing.T) {
	authTests := []struct {
		uuid     string
		request  string
		response string
		err      error
	}{
		{"0", "hgfmesf|SIGN_IN|Alpha", "hgfmesf", nil},
		{"0", "kudfemb|WHOAMI", "kudfemb|Alpha", nil},
		{"0", "zxxypev|SIGN_OUT", "zxxypev", nil},
		{"1", "ssdahny|SIGN_IN|Bravo", "ssdahny", nil},
		{"1", "mifyjcp|WHOAMI", "mifyjcp|Bravo", nil},
		{"1", "pzidies|SIGN_OUT", "pzidies", nil},
		{"2", "tsgnvdr|SIGN_IN|Charlie", "tsgnvdr", nil},
		{"2", "krisrfv|WHOAMI", "krisrfv|Charlie", nil},
		{"2", "rujbwob|SIGN_OUT", "rujbwob", nil},
		{"1", "akapqqn|SIGN_IN|Alpha", "akapqqn", nil},
		{"2", "grnssff|SIGN_IN|Bravo", "grnssff", nil},
		{"3", "rzarqfx|SIGN_IN|Charlie", "rzarqfx", nil},
		{"1", "csqhviu|WHOAMI", "csqhviu|Alpha", nil},
		{"2", "yaewjvu|WHOAMI", "yaewjvu|Bravo", nil},
		{"3", "anyzpoh|WHOAMI", "anyzpoh|Charlie", nil},
		{"1", "ykeybfr|SIGN_IN|Alpha", "ykeybfr", nil},
		{"9", "invalid_request_id", "", fmt.Errorf("any error")},
	}

	for _, tt := range authTests {
		response, err := HandleRequest(tt.uuid, tt.request)
		if err != nil && tt.err == nil {
			t.Errorf("HandleRequest(%q, %q) returned error: %v", tt.uuid, tt.request, err)
		}
		if response != tt.response {
			t.Errorf("HandleRequest(%q, %q) = %q, want %q", tt.uuid, tt.request, response, tt.response)
		}
	}
}
