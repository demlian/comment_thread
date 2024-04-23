package model

import (
	"fmt"
	"strings"
	"sync"

	"github.com/demlian/comment_thread/pkg/auth"
	"github.com/google/uuid"
)

func generateShortUUID() string {
	uuid := uuid.New().String()
	uuid = strings.ReplaceAll(uuid, "-", "")
	uuid = uuid[:7]
	return uuid
}

type Discussion struct {
	Comments     []*Comment
	DiscussionID string
	Reference    string
}

type Comment struct {
	Content  string
	UserName string
}

var discussionsById sync.Map              // Map to store discussion ID -> *Discussion
var discussionsByReferencePrefix sync.Map // Map to store reference -> discussion ID

func CreateDiscussion(userUUID, reference, comment string) (string, error) {
	discussionId := generateShortUUID()
	userName, err := auth.HandleWhoAmI(userUUID)
	if err != nil {
		return "", err
	}
	discussion := &Discussion{
		Comments:     []*Comment{{comment, userName}},
		DiscussionID: discussionId,
		Reference:    reference,
	}
	discussionsById.Store(discussionId, discussion)
	referencePrefix := strings.Split(reference, ".")[0]
	discussionsByReferencePrefix.Store(referencePrefix, discussionId)
	return discussionId, nil
}

func CreateReply(userUUID, discussionID, comment string) error {
	userName, err := auth.HandleWhoAmI(userUUID)
	if err != nil {
		return err
	}
	discussion, ok := discussionsById.Load(discussionID)
	if !ok {
		return nil
	}
	discussion.(*Discussion).Comments = append(discussion.(*Discussion).Comments, &Comment{comment, userName})
	return nil
}

func GetDiscussion(discussionID string) (*Discussion, error) {
	discussion, ok := discussionsById.Load(discussionID)
	if !ok {
		return nil, fmt.Errorf("discussion ID %s not found", discussionID)
	}
	return discussion.(*Discussion), nil
}

func ListDiscussions(referencePrefix string) ([]*Discussion, error) {
	discussionID, ok := discussionsByReferencePrefix.Load(referencePrefix)
	if !ok {
		return nil, fmt.Errorf("reference %s not found", referencePrefix)
	}
	discussion, ok := discussionsById.Load(discussionID.(string))
	if !ok {
		return nil, fmt.Errorf("discussion ID %s not found", discussionID)
	}
	return []*Discussion{discussion.(*Discussion)}, nil
}
