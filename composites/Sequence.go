package composites

import "github.com/pandarayc/behavior3go/core"

type Sequence struct {
	core.Composite
}

func (node *Sequence) OnTick(tick *core.Tick) core.NodeStatus {
	for i := 0; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != core.STATUS_SUCCESS {
			return status
		}
	}
	return core.STATUS_SUCCESS
}
