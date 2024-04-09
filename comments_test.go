package comments

import "testing"

func TestAddComment(t *testing.T) {
	t.Run("add root content when none exists", func(t *testing.T) {
		ct := &CommentThread{}
		content := "Root comment"
		userUUID := "user-1"
		ct.AddComment(content, userUUID, nil)
		if ct.root == nil {
			t.Error("Root content was not added")
		} else if ct.root.content != content || ct.root.userUUID != userUUID {
			t.Errorf("Root content was not added correctly, got: %v", ct.root)
		}
	})

	t.Run("add child content to root", func(t *testing.T) {
		ct := &CommentThread{}
		rootComment := "Root content"
		rootUserUUID := "user-1"
		ct.AddComment(rootComment, rootUserUUID, nil)

		childComment := "Child content"
		childUserUUID := "user-2"
		ct.AddComment(childComment, childUserUUID, ct.root)

		if len(ct.root.children) != 1 {
			t.Error("Child content was not added to root")
		} else if ct.root.children[0].content != childComment || ct.root.children[0].userUUID != childUserUUID {
			t.Errorf("Child content was not added correctly, got: %v", ct.root.children[0])
		}
	})

	t.Run("add child content to another comment", func(t *testing.T) {
		ct := &CommentThread{}
		rootComment := "Root content"
		rootUserUUID := "user-1"
		ct.AddComment(rootComment, rootUserUUID, nil)

		childComment := "Child content"
		childUserUUID := "user-2"
		childNode := ct.AddComment(childComment, childUserUUID, ct.root)

		grandChildComment := "Grandchild content"
		grandChildUserUUID := "user-3"
		ct.AddComment(grandChildComment, grandChildUserUUID, childNode)

		if len(childNode.children) != 1 {
			t.Error("Grandchild content was not added to child")
		} else if childNode.children[0].content != grandChildComment || childNode.children[0].userUUID != grandChildUserUUID {
			t.Errorf("Grandchild content was not added correctly, got: %v", childNode.children[0])
		}
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("update root content", func(t *testing.T) {
		ct := &CommentThread{}
		content := "Root comment"
		userUUID := "user-1"
		ct.AddComment(content, userUUID, nil)

		newComment := "Updated root comment"
		ct.UpdateComment(newComment, userUUID, ct.root)

		if ct.root.content != newComment {
			t.Errorf("Root comment was not updated correctly, got: %v", ct.root.content)
		}
	})

	t.Run("update child content", func(t *testing.T) {
		ct := &CommentThread{}
		rootComment := "Root content"
		rootUserUUID := "user-1"
		ct.AddComment(rootComment, rootUserUUID, nil)

		childComment := "Child content"
		childUserUUID := "user-2"
		childNode := ct.AddComment(childComment, childUserUUID, ct.root)

		newChildComment := "Updated child content"
		ct.UpdateComment(newChildComment, childUserUUID, childNode)

		if ct.root.children[0].content != newChildComment {
			t.Errorf("Child content was not updated correctly, got: %v", ct.root.children[0].content)
		}
	})
}
