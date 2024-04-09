package comments

import "fmt"

type CommentNode struct {
	content  string
	userUUID string
	children []*CommentNode
}

type CommentThread struct {
	videoUUID string
	root      *CommentNode
}

func NewCommentNode(content, userUUID string) *CommentNode {
	return &CommentNode{
		content:  content,
		userUUID: userUUID,
		children: []*CommentNode{},
	}
}

func (c *CommentNode) AddChild(content, userUUID string) *CommentNode {
	childNode := NewCommentNode(content, userUUID)
	c.children = append(c.children, childNode)
	return childNode
}

func (c *CommentNode) UpdateContent(content, userUUID string) error {
	print(c.userUUID)
	if c.userUUID == userUUID {
		c.content = content
		return nil
	}
	return fmt.Errorf("only the user who created the content can update it")
}

func NewCommentThread(videoUUID string, root *CommentNode) *CommentThread {
	return &CommentThread{
		videoUUID: videoUUID,
		root:      root,
	}
}

func (c *CommentThread) AddComment(comment, userUUID string, parent *CommentNode) *CommentNode {
	if parent == nil {
		if c.root == nil {
			c.root = NewCommentNode(comment, userUUID)
		} else {
			return c.root.AddChild(comment, userUUID)
		}
	} else {
		return parent.AddChild(comment, userUUID)
	}
	return nil
}

func (c *CommentThread) UpdateComment(comment, userUUID string, node *CommentNode) error {
	if node == nil {
		return fmt.Errorf("no comments exist in the thread")
	}
	return node.UpdateContent(comment, userUUID)
}
