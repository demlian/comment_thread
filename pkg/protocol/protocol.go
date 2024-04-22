package protocol

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/demlian/comment_thread/pkg/auth"
)

type Request struct {
	requestId string
	verb      string
	data      string
}

type Response struct {
	ResponseId string
	Data       string
}

func isValidRequestId(request *Request) bool {
	pattern := "^[a-z]{7}$"
	matched, _ := regexp.MatchString(pattern, request.requestId)
	return matched
}

func adaptTextToRequest(data string) (*Request, error) {
	parts := strings.Split(data, "|")
	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("invalid request format: %s", data)
	}
	request := &Request{requestId: strings.TrimSpace(parts[0])}
	if !isValidRequestId(request) {
		return nil, fmt.Errorf("invalid request ID format: %s", request.requestId)
	}
	request.verb = strings.TrimSpace(parts[1])
	if len(parts) > 2 {
		request.data = strings.TrimSpace(parts[2])
	}

	return request, nil
}

func HandleRequest(uuid string, data string) (string, error) {
	req, err := adaptTextToRequest(data)
	if err != nil {
		return "", err
	}
	var rsp Response
	switch req.verb {
	case "SIGN_IN":
		err = auth.HandleSignIn(uuid, req.data)
		rsp = Response{ResponseId: req.requestId}
	case "SIGN_OUT":
		err = auth.HandleSignOut(uuid)
		rsp = Response{ResponseId: req.requestId}
	case "WHOAMI":
		var userName string
		userName, err = auth.HandleWhoAmI(uuid)
		rsp = Response{ResponseId: req.requestId, Data: userName}
	default:
		rsp = Response{ResponseId: "request ID: " + req.requestId, Data: " request type not supported: " + req.verb}
	}

	responseString := adaptResponseToText(&rsp)
	return responseString, err
}

func adaptResponseToText(rsp *Response) string {
	var responseString string
	if rsp.Data != "" {
		responseString = fmt.Sprintf("%s|%s", rsp.ResponseId, rsp.Data)
	} else {
		responseString = rsp.ResponseId
	}

	return responseString
}
