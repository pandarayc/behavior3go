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

func (node *BaseNode) _execute(tick *Tick) {
	node._enter(tick)
	// if (!tick.BlackBoard)

}

func (node *BaseNode) _enter(tick *Tick) {
	tick._enterNode(node)
	node.Enter(tick)
}

func (node *BaseNode) _open(tick *Tick) {
	node._openNode(tick)
	node.Open(tick)
}

/**
 * Enter method, override this to use. It is called every time a node is
 * asked to execute, before the tick itself.
 *
 * @method enter
 * @param {Tick} tick A tick instance.
 **/
func (node *BaseNode) Enter(tick *Tick) {
}

/**
 * Open method, override this to use. It is called only before the tick
 * callback and only if the not isn't closed.
 *
 * Note: a node will be closed if it returned `RUNNING` in the tick.
 *
 * @method open
 * @param {Tick} tick A tick instance.
 **/
func (node *BaseNode) Open(tick *Tick) {
}

/**
 * Tick method, override this to use. This method must contain the real
 * execution of node (perform a task, call children, etc.). It is called
 * every time a node is asked to execute.
 *
 * @method tick
 * @param {Tick} tick A tick instance.
 **/
func (node *BaseNode) Tick(tick *Tick) {
}

/**
 * Close method, override this to use. This method is called after the tick
 * callback, and only if the tick return a state different from
 * `RUNNING`.
 *
 * @method close
 * @param {Tick} tick A tick instance.
 **/
func (node *BaseNode) Close(tick *Tick) {
}

/**
 * Exit method, override this to use. Called every time in the end of the
 * execution.
 *
 * @method exit
 * @param {Tick} tick A tick instance.
 **/
func (node *BaseNode) Exit(tick *Tick) {
}
