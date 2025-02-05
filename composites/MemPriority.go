package composites

import "github.com/pandarayc/behavior3go/core"

type IMemPriority interface {
	core.IWorker
}

type MemPriority struct {
	core.Composite
}


func (node *MemPriority) OnOpen(tick *core.Tick) {


}