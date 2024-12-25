package core

type NodeStatus int32

const (
	STATUS_SUCCESS NodeStatus = iota
	STATUS_FAILURE
	STATUS_RUNNING
	STATUS_ERROR
)

type NodeCategory string

const (
	CATEGORY_ACTION    NodeCategory = "action"
	CATEGORY_DECORATOR NodeCategory = "decorator"
	CATEGORY_COMPOSITE NodeCategory = "composite"
	CATEGORY_CONDITION NodeCategory = "condition"
)

