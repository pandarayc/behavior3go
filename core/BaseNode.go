package core

/**
 * The BaseNode class is used as super class to all nodes in BehaviorJS. It
 * comprises all common variables and methods that a node must have to
 * execute.
 *
 * **IMPORTANT:** Do not inherit from this class, use `Composite`,
 * `Decorator`, `Action` or `Condition`, instead.
 *
 * The attributes are specially designed to serialization of the node in a
 * JSON format. In special, the `parameters` attribute can be set into the
 * visual editor (thus, in the JSON file), and it will be used as parameter
 * on the node initialization at `BehaviorTree.load`.
 *
 * BaseNode also provide 5 callback methods, which the node implementations
 * can override. They are `enter`, `open`, `tick`, `close` and `exit`. See
 * their documentation to know more. These callbacks are called inside the
 * `_execute` method, which is called in the tree traversal.
 *
 * @module b3
 * @class BaseNode
 **/

type INode interface {
	OnEnter(tick *Tick)
	OnOpen(tick *Tick)
	OnClose(tick *Tick)
	OnExit(tick *Tick)
	OnTick(tick *Tick) NodeStatus
}

var _ INode = &BaseNode{}

// b3 节点信息 (数据存储载体)
type BaseNode struct {
	INode
	Id          string // uuid
	Category    NodeCategory
	Name        string // name
	Title       string
	Description string
	Properties  map[string]interface{}
	Parameters  map[string]interface{}
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
