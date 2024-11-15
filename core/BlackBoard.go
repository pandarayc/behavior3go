package core

type TreeMemory struct {
	nodeMemory     map[string]interface{} // 节点数据
	openNodes      []INode                // 节点数据
	traversalDepth int32
	traversalCycle int32
}

// 数据黑板
type BlackBoard struct {
	_baseMemory map[string]interface{}
	_treeMemory map[string]*TreeMemory
}

type memory map[string]interface{}

func (m *memory) GetBool() bool {
	return false
}

func (m *memory) GetInt64() int64 {
	return 0
}

func (b *BlackBoard) getTreeMemory(treeId string) *TreeMemory {
	return nil
})




func (b *BlackBoard) Set() {

}

func (b *BlackBoard) Get() {

}
