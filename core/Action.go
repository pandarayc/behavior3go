package core

type Action struct {
	Worker
}

// NewAction creates a new Action.
func NewAction(name string, title string, properties map[string]interface{}) INode {
	if name == "" {
		name = "Action"
	}

	return &Node{
		Name:       name,
		Title:      title,
		Properties: properties,
		Category:   CATEGORY_ACTION,
		INode:      &Action{},
	}
}
