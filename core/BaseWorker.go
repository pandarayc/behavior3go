package core

// go语言的问题，莫得方法重载，用worker 来代替

type IWorker interface {
	OnEnter(tick *Tick)
	OnOpen(tick *Tick)
	OnClose(tick *Tick)
	OnExit(tick *Tick)
	OnTick(tick *Tick) NodeStatus
}

var _ IWorker = &BaseWorker{}

type BaseWorker struct {
}

func (w *BaseWorker) OnEnter(tick *Tick) {
}

func (w *BaseWorker) OnOpen(tick *Tick) {
}

func (w *BaseWorker) OnTick(tick *Tick) NodeStatus {
	// 默认成功
	return STATUS_SUCCESS
}

func (w *BaseWorker) OnClose(tick *Tick) {
}

func (w *BaseWorker) OnExit(tick *Tick) {
}
