package behavior3go


type Priority struct {
	Composite
}

func (node *Priority) OnTick(tick *Tick) NodeStatus {
	for i := 0; i < node.GetChildrenCount(); i++ {
		status := node.GetChild(i).Execute(tick)

		if status != STATUS_FAILURE {
			return status
		}
	}
	return STATUS_FAILURE
}
