package decorators

import (
	"github.com/pandarayc/behavior3go/config"
	"github.com/pandarayc/behavior3go/core"
)

type Limiter struct {
	core.Decorator
	maxLoop int
}

func (node *Limiter) Initialize(cfg *config.NodeCfg) {
	node.Decorator.Initialize(cfg)
	node.maxLoop = cfg.GetPropertyAsInt(core.PROPERTY_KEY_MAX_LOOP)
	if node.maxLoop < 1 {
		panic("Limiter maxLoop must > 0")
	}
}

func (node *Limiter) OnOpen(tick *core.Tick) {
	tick.GetMemory(node.GetId()).Set(core.BlackBoard_KEY_RANGE_INDEX, 0)
}

func (node *Limiter) OnTick(tick *core.Tick) core.NodeStatus {
	if node.GetChild() == nil {
		return core.STATUS_ERROR
	}

	index := tick.GetMemory(node.GetId()).GetInt(core.BlackBoard_KEY_RANGE_INDEX)
	if index < node.maxLoop {
		status := node.GetChild().Execute(tick)
		if status == core.STATUS_SUCCESS || status == core.STATUS_FAILURE {
			tick.GetMemory(node.GetId()).Set(core.BlackBoard_KEY_RANGE_INDEX, index+1)
			return status
		}
	}
	return core.STATUS_FAILURE
}
