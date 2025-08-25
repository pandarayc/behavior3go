package behavior3go

import (
	"github.com/pandarayc/behavior3go/config"
)

type RepeatUntilFailure struct {
	Decorator
	maxLoop int
}

func (node *RepeatUntilFailure) Initialize(cfg *config.NodeCfg) {
	node.Decorator.Initialize(cfg)
	node.maxLoop = cfg.GetPropertyAsInt(PROPERTY_KEY_MAX_LOOP)
}

func (node *RepeatUntilFailure) OnOpen(tick *Tick) {
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RANGE_INDEX, 0)
}

func (node *RepeatUntilFailure) OnTick(tick *Tick) NodeStatus {
	if node.GetChild() == nil {
		return STATUS_ERROR
	}

	index := tick.GetMemory(node.GetId()).GetInt(BLACKBOARD_KEY_RANGE_INDEX)
	status := STATUS_FAILURE

	for node.maxLoop < 0 || index < node.maxLoop {
		status = node.GetChild().Execute(tick)
		if status == STATUS_SUCCESS {
			index++
		} else {
			break
		}
	}
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RANGE_INDEX, index)
	return status
}
