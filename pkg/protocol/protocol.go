package protocol

import (
	"fmt"
	"strings"

	"github.com/demlian/comment_thread/pkg/auth"
)

type Request struct {
	ID   string
	Type string
	Data string
}

type Response struct {
	ID   string
	Data string
}

func newRequest(data string) (*Request, error) {
	parts := strings.Split(data, "|")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid request format: %s", data)
	}

	request := &Request{
		ID:   "",
		Type: "",
		Data: data,
	}
	if len(parts) > 0 && parts[0] != "" {
		request.ID = parts[0]
	}
	if len(parts) > 1 && parts[1] != "" {
		request.Type = parts[1]
	}
	if len(parts) > 2 && parts[2] != "" {
		request.Data = parts[2]
	}

	return request, nil
}

func HandleRequest(connHash string, data string) (*Response, error) {
	req, err := newRequest(data)
	if err != nil {
		return nil, err
	}
	var rsp Response
	switch strings.TrimSpace(req.Type) {
	case "SIGN_IN":
		err = auth.HandleSignIn(connHash, req.Data)
		rsp = Response{ID: req.ID}
	case "SIGN_OUT":
		err = auth.HandleSignOut(connHash)
		rsp = Response{ID: req.ID}
	case "WHOAMI":
		userName := auth.HandleWhoAmI(connHash)
		rsp = Response{ID: req.ID, Data: userName}
	default:
		rsp = Response{ID: "request ID: " + req.ID, Data: " request type not supported: " + req.Type}
	}

	return &rsp, err
}
