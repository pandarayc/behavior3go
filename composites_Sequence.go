package behavior3go


type Sequence struct {
	Composite
}

func (node *Sequence) OnTick(tick *Tick) NodeStatus {
	for i := 0; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != STATUS_SUCCESS {
			return status
		}
	}
	return STATUS_SUCCESS
}
