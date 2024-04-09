package comments

import (
	"testing"
)

func TestAddComment(t *testing.T) {
	storage := &InMemoryStorage{
		comments: make(map[string]*CommentNode),
	}
	commentThread := NewCommentThread("video-uuid", storage)

	rootComment := commentThread.AddComment("Root comment", "user-uuid", nil)
	childComment := commentThread.AddComment("Child comment", "user-uuid", rootComment)
	if childComment.parentID != rootComment.id {
		expectedParentID := rootComment.id
		gotParentID := childComment.parentID
		t.Errorf("child comment has incorrect parent ID. Expected: %s, Got: %s", expectedParentID, gotParentID)
	}
}

func TestUpdateComment(t *testing.T) {
	storage := &InMemoryStorage{
		comments: make(map[string]*CommentNode),
	}
	commentThread := NewCommentThread("video-uuid", storage)

	comment := commentThread.AddComment("Original comment", "user-uuid", nil)
	updatedComment := commentThread.UpdateComment("Updated comment", "user-uuid", comment)
	if updatedComment.content != "Updated comment" {
		expectedContent := "Updated comment"
		gotContent := updatedComment.content
		t.Errorf("Comment content was not updated. Expected: %s, Got: %s", expectedContent, gotContent)
	}
}
