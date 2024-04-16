package model

import "github.com/google/uuid"

type StorageAdapter interface {
	AddContent(comment *CommentNode, parentID string)
	UpdateContent(comment *CommentNode)
}

type InMemoryStorage struct {
	comments map[string]*CommentNode
}

func (i *InMemoryStorage) AddContent(comment *CommentNode, parentID string) {
	comment.parentID = parentID
	i.comments[comment.id] = comment
}

func (i *InMemoryStorage) UpdateContent(comment *CommentNode) {
	if node, exists := i.comments[comment.id]; exists {
		node.content = comment.content
	}
}

type CommentNode struct {
	id       string
	content  string
	userUUID string
	parentID string
}

func NewCommentNode(content, userUUID string) *CommentNode {

	return &CommentNode{
		id:       uuid.New().String(),
		content:  content,
		userUUID: userUUID,
	}
}

type CommentThread struct {
	videoUUID string
	storage   StorageAdapter
}

func NewCommentThread(videoUUID string, storage StorageAdapter) *CommentThread {
	return &CommentThread{
		videoUUID: videoUUID,
		storage:   storage,
	}
}

func (c *CommentThread) AddComment(comment, userUUID string, parent *CommentNode) *CommentNode {
	newComment := NewCommentNode(comment, userUUID)
	if parent == nil {
		c.storage.AddContent(newComment, "")
	} else {
		c.storage.AddContent(newComment, parent.id)
	}
	return newComment
}

func (c *CommentThread) UpdateComment(comment, userUUID string, node *CommentNode) *CommentNode {
	// TODO: Ensure the node exists before updating its content.
	if node.userUUID != userUUID {
		return nil
	}
	node.content = comment
	c.storage.UpdateContent(node)
	return node
}
