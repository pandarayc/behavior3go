package behavior3go


type Runner struct {
	Action
}

func (node *Runner) OnTick(tick *Tick) NodeStatus {
	return STATUS_RUNNING
}
