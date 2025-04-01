package decorators

import (
	"github.com/pandarayc/behavior3go/config"
	"github.com/pandarayc/behavior3go/core"
)

type RepeatUntilSuccess struct {
	core.Decorator
	maxLoop int
}

func (node *RepeatUntilSuccess) Initialize(cfg *config.NodeCfg) {
	node.Decorator.Initialize(cfg)
	node.maxLoop = cfg.GetPropertyAsInt(core.PROPERTY_KEY_MAX_LOOP)
}

func (node *RepeatUntilSuccess) OnOpen(tick *core.Tick) {
	tick.GetMemory(node.GetId()).Set(core.BlackBoard_KEY_RANGE_INDEX, 0)
}

func (node *RepeatUntilSuccess) OnTick(tick *core.Tick) core.NodeStatus {
	if node.GetChild() == nil {
		return core.STATUS_ERROR
	}

	index := tick.GetMemory(node.GetId()).GetInt(core.BlackBoard_KEY_RANGE_INDEX)
	status := core.STATUS_FAILURE

	for node.maxLoop < 0 || index < node.maxLoop {
		status = node.GetChild().Execute(tick)
		if status == core.STATUS_FAILURE {
			index++
		} else {
			break
		}
	}

	tick.GetMemory(node.GetId()).Set(core.BlackBoard_KEY_RANGE_INDEX, index)
	return status
}
