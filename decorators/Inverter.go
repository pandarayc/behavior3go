package decorators

import (
	"github.com/pandarayc/behavior3go/core"
)

type Inverter struct {
	core.Decorator
}

func (node *Inverter) OnTick(tick *core.Tick) core.NodeStatus {
	if node.GetChild() == nil {
		return core.STATUS_ERROR
	}

	status := node.GetChild().Execute(tick)
	if status == core.STATUS_SUCCESS {
		status = core.STATUS_FAILURE
	} else if status == core.STATUS_FAILURE {
		status = core.STATUS_SUCCESS
	}
	return status
}
