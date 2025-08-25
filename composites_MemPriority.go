package behavior3go


type MemPriority struct {
	Composite
}

func (node *MemPriority) OnOpen(tick *Tick) {
	// 初始化运行index
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RUNNING_CHILD, 0)
}

func (node *MemPriority) OnTick(tick *Tick) NodeStatus {
	var childIndex = tick.GetMemory(node.GetId()).GetInt(BLACKBOARD_KEY_RUNNING_CHILD)
	for i := childIndex; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != STATUS_FAILURE {
			if status == STATUS_RUNNING {
				tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RUNNING_CHILD, i)
			}
			return status
		}
	}

	return STATUS_FAILURE
}
