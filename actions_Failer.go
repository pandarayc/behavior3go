package behavior3go


type Failer struct {
	Action
}

func (node *Failer) OnTick(tick *Tick) NodeStatus {
	return STATUS_FAILURE
}
