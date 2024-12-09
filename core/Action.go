package core

type Action struct {
	BaseNode
}

// NewAction creates a new Action.
func NewAction(name string, title string, properties map[string]interface{}) *Action {
	if name == "" {
		name = "Action"
	}

	return &Action{
		BaseNode: BaseNode{
			Name:       name,
			Title:      title,
			Properties: properties,
			Category:   CATEGORY_ACTION,
		},
	}
}
