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
}

// b3 节点信息
type BaseNode struct {
	Id          string // uuid
	Category    int32
	Name        string // name
	Title       string
	Description string
	Properties  map[string]interface{}
	Parameters  map[string]interface{}
}

func (node *BaseNode) GetCategory() int32 {
	return node.Category
}

func (node *BaseNode) Execute(tick *Tick) {
}

func (node *BaseNode) _enter(tick *Tick) {
}