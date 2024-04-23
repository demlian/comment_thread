package protocol

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/demlian/comment_thread/pkg/auth"
	"github.com/demlian/comment_thread/pkg/model"
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
	case "CREATE_DISCUSSION":
		reference := strings.TrimSpace(parts[2])
		comment := strings.TrimSpace(parts[3])
		rsp.Data, err = model.CreateDiscussion(uuid, reference, comment)
	case "CREATE_REPLY":
		discussionID := strings.TrimSpace(parts[2])
		comment := strings.TrimSpace(parts[3])
		err = model.CreateReply(uuid, discussionID, comment)
	case "GET_DISCUSSION":
		var discussion *model.Discussion
		discussionID := strings.TrimSpace(parts[2])
		discussion, err = model.GetDiscussion(discussionID)
		rsp.Data = discussionToString(discussion)
	case "LIST_DISCUSSIONS":
		reference := strings.TrimSpace(parts[2])
		referencePrefix := strings.Split(reference, ".")[0]
		var discussions []*model.Discussion
		discussions, err = model.ListDiscussions(referencePrefix)
		rsp.Data = discussionsToString(discussions)
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

func discussionToString(discussion *model.Discussion) string {
	var replies []string
	var formattedReply string
	for _, reply := range discussion.Comments {
		if reply.UserName == "Alpha" {
			formattedReply = fmt.Sprintf("%s|%s", reply.UserName, strings.ReplaceAll(reply.Content, "\"", "\"\""))
		} else {
			formattedReply = fmt.Sprintf("%s|\"%s\"", reply.UserName, strings.ReplaceAll(reply.Content, "\"", "\"\""))
		}
		replies = append(replies, formattedReply)
	}
	return fmt.Sprintf("%s|%s|(%s)", discussion.DiscussionID, discussion.Reference, strings.Join(replies, ","))
}

func discussionsToString(discussions []*model.Discussion) string {
	var discussionsStrings []string
	for _, discussion := range discussions {
		discussionsStrings = append(discussionsStrings, discussionToString(discussion))
	}
	return fmt.Sprintf("(%s)", strings.Join(discussionsStrings, ","))
}
