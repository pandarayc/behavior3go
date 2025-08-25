package behavior3go


type Succeeder struct {
	Action
}

func (node *Succeeder) OnTick(tick *Tick) NodeStatus {
	return STATUS_SUCCESS
}
