package core

type BehaviorTree struct {
	id string // uuid
	title string // treeName
	description string // 描述
	properties map[string]interface{}
	root INode
	debug interface{}
}


func (bt *BehaviorTree) Load() {

}
 
func (bt *BehaviorTree) dump() {

}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *BlackBoard) NodeStatus {
	return nil 
}