package behavior3go


type MemSequence struct {
	Composite
}

func (node *MemSequence) OnOpen(tick *Tick) {
	// 初始化运行index
	tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RUNNING_CHILD, 0)
}

func (node *MemSequence) OnTick(tick *Tick) NodeStatus {
	var childIndex = tick.GetMemory(node.GetId()).GetInt(BLACKBOARD_KEY_RUNNING_CHILD)
	for i := childIndex; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != STATUS_SUCCESS {
			if status == STATUS_RUNNING {
				tick.GetMemory(node.GetId()).Set(BLACKBOARD_KEY_RUNNING_CHILD, i)
				return status
			}
		}
	}

	return STATUS_SUCCESS
}
