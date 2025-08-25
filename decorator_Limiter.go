package behavior3go

import (
	"github.com/pandarayc/behavior3go/config"
)

type Limiter struct {
	Decorator
	maxLoop int
}

func (node *Limiter) Initialize(cfg *config.NodeCfg) {
	node.Decorator.Initialize(cfg)
	node.maxLoop = cfg.GetPropertyAsInt(PROPERTY_KEY_MAX_LOOP)
	if node.maxLoop < 1 {
		panic("Limiter maxLoop must > 0")
	}
}

func (node *Limiter) OnOpen(tick *Tick) {
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RANGE_INDEX, 0)
}

func (node *Limiter) OnTick(tick *Tick) NodeStatus {
	if node.GetChild() == nil {
		return STATUS_ERROR
	}

	index := tick.GetMemory(node.GetId()).GetInt(BLACKBOARD_KEY_RANGE_INDEX)
	if index < node.maxLoop {
		status := node.GetChild().Execute(tick)
		if status == STATUS_SUCCESS || status == STATUS_FAILURE {
			tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RANGE_INDEX, index+1)
			return status
		}
	}
	return STATUS_FAILURE
}
