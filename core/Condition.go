package core


type Condition struct {
    BaseNode
}

func NewCondition(name string, title string, properties map[string]interface{}) *Condition {
    if name == "" {
        name = "Condition"
    }

    return &Condition{
        BaseNode: BaseNode{
            Name:        name,
            Title:       title,
            Properties:  properties,
            Category:    CATEGORY_CONDITION,
        },
    }
}