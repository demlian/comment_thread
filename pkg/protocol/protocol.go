package protocol

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/demlian/comment_thread/pkg/auth"
)

type Request struct {
	id       string
	verb     string
	clientId string
}

type Response struct {
	ID   string
	Data string
}

func isValidRequestId(requestChunk string) (bool, error) {
	pattern := "^[a-z]{7}$"
	return regexp.MatchString(pattern, requestChunk)
}

func HandleRequest(uuid string, data string) (string, error) {
	request := &Request{}
	parts := strings.Split(data, "|")
	rid := parts[0]
	valid, err := isValidRequestId(rid)
	if !valid || err != nil {
		return "", fmt.Errorf("invalid request ID format: %s", rid)
	}
	request.id = rid
	request.verb = strings.TrimSpace(parts[1])
	rsp := &Response{ID: request.id}
	switch request.verb {
	case "SIGN_IN":
		request.clientId = strings.TrimSpace(parts[2])
		err = auth.HandleSignIn(uuid, request.clientId)
	case "SIGN_OUT":
		err = auth.HandleSignOut(uuid)
	case "WHOAMI":
		rsp.Data, err = auth.HandleWhoAmI(uuid)
	default:
		err = fmt.Errorf("unknown verb: %s", request.verb)
	}

	var responseString string
	if rsp.Data != "" {
		responseString = fmt.Sprintf("%s|%s", rsp.ID, rsp.Data)
	} else {
		responseString = rsp.ID
	}
	return responseString, err
}
