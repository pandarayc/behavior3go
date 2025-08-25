package behavior3go


type Error struct {
	Action
}

//  通过配置表注入的正常不用重新赋值name

func (node *Error) OnTick(tick *Tick) NodeStatus {
	return STATUS_ERROR
}
