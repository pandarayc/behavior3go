package core

import "github.com/pandarayc/behavior3go/config"

// 方便修改
type INodeWrapper interface {
	_execute(tick *Tick) NodeStatus
	_enter(tick *Tick)
	_open(tick *Tick)
	_tick(tick *Tick) NodeStatus
	_close(tick *Tick)
	_exit(tick *Tick)
}

var _ INodeWrapper = &BaseNode{}

// 外层修改不到Base的方法，只能通过interface方式组合
// 节点通用方法
type INode interface {
	// INodeWrapper
	config.IBaseConfig
	SetWorker(IWorker)
	Ctor()
	Initialize(*config.NodeCfg) // 初始化时把外层结构注入
	GetCategory() NodeCategory
	Execute(tick *Tick) NodeStatus
}

var _ INode = &BaseNode{}

// b3 节点信息 (数据存储载体)
type BaseNode struct {
	IWorker // 为了把外部结构重载方法注入进来使用
	*config.NodeCfg
	category NodeCategory
}

// Initialize 初始化节点信息
func (node *BaseNode) Initialize(cfg *config.NodeCfg) {
	// 偷懒偷懒
	node.NodeCfg = cfg
}

func (node *BaseNode) GetCategory() NodeCategory {
	return node.category
}

// 设置工作函数
func (node *BaseNode) SetWorker(worker IWorker) {
	node.IWorker = worker
}

// ctor 构造函数
func (node *BaseNode) Ctor() {
}

// _execute 执行节点
func (node *BaseNode) _execute(tick *Tick) NodeStatus {
	node._enter(tick)
	// 检查自己是否处于已打开状态
	if !tick.GetMemory(node.GetId()).GetBool(BLACKBOARD_KEY_IS_OPEN) {
		node._enter(tick)
	}

	status := node._tick(tick)
	if status != STATUS_RUNNING {
		node._close(tick)
	}
	node._exit(tick)
	return status
}

// Execute 执行节点 (官方是protected go需要导出使用)
func (node *BaseNode) Execute(tick *Tick) NodeStatus {
	return node._execute(tick)
}

// _enter 进入节点
func (node *BaseNode) _enter(tick *Tick) {
	tick._enterNode(node)
	node.OnEnter(tick)
}

// _open 打开节点
func (node *BaseNode) _open(tick *Tick) {
	tick._openNode(node)
	// 设置open
	node.OnOpen(tick)
}

// _tick 节点运行
func (node *BaseNode) _tick(tick *Tick) NodeStatus {
	tick._tickNode(node)
	return node.OnTick(tick)
}

// _close 关闭节点
func (node *BaseNode) _close(tick *Tick) {
	tick._closeNode(node)
	// 设置isopen close
	node.OnClose(tick)
}

// _exit 节点退出
func (node *BaseNode) _exit(tick *Tick) {
	tick._exitNode(node)
	node.OnExit(tick)
}
