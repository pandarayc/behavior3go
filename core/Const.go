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

const (
	PROPERTY_KEY_WAIT_MS                = "milliseconds" // ACTION_WAIT 参数 毫秒
	PROPERTY_KEY_REPEATER_LOOP          = "maxLoop"      // decorator
	PROPERTY_KEY_REPEATER_UNTIL_FAILURE = "maxLoop"
	PROPERTY_KEY_REPEATER_UNTIL_SUCCESS = "maxLoop"
	PROPERTY_KEY_LIMITER                = "maxLoop"
)

const (
	BLACKBOARD_KEY_WAIT_START = "startTime"
	BLACKBOARD_KEY_IS_OPEN    = "isOpen"
	BLACKBOARD_KEY_RUNNING_CHILD = "runningChild"
)
