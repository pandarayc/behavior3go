package core

type BehaviorTree struct {
	Id string // uuid
	Title string // treeName
	Description string // 描述
	Properties map[string]interface{}
	Root INode
	Debug interface{}
}


func (bt *BehaviorTree) Load() {

}
 
func (bt *BehaviorTree) Dump() {

}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *BlackBoard) NodeStatus {
	return 0 
}