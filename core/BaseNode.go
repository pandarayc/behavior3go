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
	Enter(tick *Tick)
	Open(tick *Tick)
	Close(tick *Tick)
	Exit(tick *Tick)
	Tick(tick *Tick) NodeStatus
}

var _ INode = &BaseNode{}

// b3 节点信息
type BaseNode struct {
	id          string // uuid
	category    int32
	name        string // name
	title       string
	description string
	properties  map[string]interface{}
	parameters  map[string]interface{}
}

func (node *BaseNode) GetCategory() int32 {
	return node.category
}

func (node *BaseNode) _execute(tick *Tick) NodeStatus {
	node._enter(tick)
	// 检查自己是否处于已打开状态
	// if (!tick.BlackBoard)
	{
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
	node.Enter(tick)
}

func (node *BaseNode) _open(tick *Tick) {
	tick._openNode(tick)
	// 设置open
	node.Open(tick)
}

func (node *BaseNode) _tick(tick *Tick) NodeStatus {
	tick._tickNode(node)
	return node.Tick(tick)
}

func (node *BaseNode) _close(tick *Tick) {
	tick._closeNode(node)
	// 设置isopen close
	node.Close(tick)
}

func (node *BaseNode) _exit(tick *Tick) {
	tick._exitNode(node)
	node.Exit(tick)
}

func (node *BaseNode) Enter(tick *Tick) {
}

func (node *BaseNode) Open(tick *Tick) {
}

func (node *BaseNode) Tick(tick *Tick) NodeStatus {
	// 默认成功
	return STATUS_SUCCESS
}

func (node *BaseNode) Close(tick *Tick) {
}

func (node *BaseNode) Exit(tick *Tick) {
}
