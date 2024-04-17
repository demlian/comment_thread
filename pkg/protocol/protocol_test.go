package protocol

import (
	"testing"
)

func TestHandleRequest(t *testing.T) {
	tests := []struct {
		connHash string
		request  string
		response string
	}{
		{"0", "hgfmesf|SIGN_IN|Alpha", "hgfmesf"},
		{"0", "kudfemb|WHOAMI", "kudfemb|Alpha"},
		{"0", "zxxypev|SIGN_OUT", "zxxypev"},
		{"1", "ssdahny|SIGN_IN|Bravo", "ssdahny"},
		{"1", "mifyjcp|WHOAMI", "mifyjcp|Bravo"},
		{"1", "pzidies|SIGN_OUT", "pzidies"},
		{"2", "tsgnvdr|SIGN_IN|Charlie", "tsgnvdr"},
		{"2", "krisrfv|WHOAMI", "krisrfv|Charlie"},
		{"2", "rujbwob|SIGN_OUT", "rujbwob"},
		{"1", "akapqqn|SIGN_IN|Alpha", "akapqqn"},
		{"2", "grnssff|SIGN_IN|Bravo", "grnssff"},
		{"3", "rzarqfx|SIGN_IN|Charlie", "rzarqfx"},
		{"1", "csqhviu|WHOAMI", "csqhviu|Alpha"},
		{"2", "yaewjvu|WHOAMI", "yaewjvu|Bravo"},
		{"3", "anyzpoh|WHOAMI", "anyzpoh|Charlie"},
		{"1", "ykeybfr|SIGN_IN|Alpha", "ykeybfr"},
	}

	for _, tt := range tests {
		response, err := HandleRequest(tt.connHash, tt.request)
		if err != nil {
			t.Errorf("HandleRequest(%q, %q) returned error: %v", tt.connHash, tt.request, err)
		}
		if response != tt.response {
			t.Errorf("HandleRequest(%q, %q) = %q, want %q", tt.connHash, tt.request, response, tt.response)
		}
	}
}
