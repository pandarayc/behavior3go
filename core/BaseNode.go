package core

import "github.com/genet9496/behavior3go/config"

// 外层修改不到Base的方法，只能通过interface方式组合
// 节点通用方法
type INode interface {
	Initialize(*config.NodeCfg) // 初始化时把外层结构注入
	GetId() string
	GetCategory() NodeCategory
	GetName() string
	GetTitle() string
}

// b3 节点信息 (数据存储载体)
type BaseNode struct {
	IWorker            // 为了把外部结构重载方法注入进来使用
	id          string // uuid
	category    NodeCategory
	name        string // name
	title       string
	description string
	properties  map[string]interface{}
	parameters  map[string]interface{}
}

func (node *BaseNode) Initialize(cfg *config.NodeCfg) {

}

func (node *BaseNode) _execute(tick *Tick) NodeStatus {
	node._enter(tick)
	// 检查自己是否处于已打开状态
	if !tick.BlackBoard.GetMemory(tick.Tree.Id, node.Id).GetBool("isOpen") {
		node._enter(tick)
	}

	status := node._tick(tick)
	if status != STATUS_RUNNING {
		node._close(tick)
	}
	node._exit(tick)
	return status
}

func (node *BaseNode) _enter(tick *Tick) {
	tick._enterNode(node)
	node.OnEnter(tick)
}

func (node *BaseNode) _open(tick *Tick) {
	tick._openNode(node)
	// 设置open
	node.OnOpen(tick)
}

func (node *BaseNode) _tick(tick *Tick) NodeStatus {
	tick._tickNode(node)
	return node.OnTick(tick)
}

func (node *BaseNode) _close(tick *Tick) {
	tick._closeNode(node)
	// 设置isopen close
	node.OnClose(tick)
}

func (node *BaseNode) _exit(tick *Tick) {
	tick._exitNode(node)
	node.OnExit(tick)
}
