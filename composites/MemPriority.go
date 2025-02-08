package composites

import "github.com/pandarayc/behavior3go/core"
type MemPriority struct {
	core.Composite
}

func (node *MemPriority) OnOpen(tick *core.Tick) {
	// 初始化运行index
	tick.GetBlackBoard().GetMemory(tick.GetTree().GetId(), node.GetId()).Set(core.BLACKBOARD_KEY_RUNNING_CHILD, 0)
}

func (node *MemPriority) OnTick(tick *core.Tick) core.NodeStatus {
	var childIndex = tick.GetBlackBoard().GetMemory(tick.GetTree().GetId(), node.GetId()).GetInt(core.BLACKBOARD_KEY_RUNNING_CHILD)
	for i := childIndex; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != core.STATUS_FAILURE {
			if status == core.STATUS_RUNNING {
				tick.GetBlackBoard().GetMemory(tick.GetTree().GetId(), node.GetId()).Set(core.BLACKBOARD_KEY_RUNNING_CHILD, i)
			}
			return status
		}
	}

	return core.STATUS_FAILURE
}
