package composites

import "github.com/pandarayc/behavior3go/core"

type Priority struct {
	core.Composite
}

func (node *Priority) OnTick(tick *core.Tick) core.NodeStatus {
	for i := 0; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != core.STATUS_FAILURE {
			return status
		}
	}
	return core.STATUS_FAILURE
}
