package behavior3go

import (
)

type Inverter struct {
	Decorator
}

func (node *Inverter) OnTick(tick *Tick) NodeStatus {
	if node.GetChild() == nil {
		return STATUS_ERROR
	}

	status := node.GetChild().Execute(tick)
	if status == STATUS_SUCCESS {
		status = STATUS_FAILURE
	} else if status == STATUS_FAILURE {
		status = STATUS_SUCCESS
	}
	return status
}
