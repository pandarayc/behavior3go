package core

// 数据黑板
type BlackBoard struct {
	baseMemory *memory
	treeMemory map[string]*TreeMemory
}

// 作为共享信息层处理
type TreeMemory struct {
	nodeMemory map[string]*memory // 节点数据
	openNodes  []INode            // 节点数据
	*memory
	traversalDepth int32
	traversalCycle int32
}

// 新建黑板数据
func NewBlackBoard() *BlackBoard {
	blackBoard := &BlackBoard{
		baseMemory: newMemory(),
		treeMemory: make(map[string]*TreeMemory),
	}

	return blackBoard
}

func (b *BlackBoard) getMemory(treeId, nodeId string) *memory {
	mem := b.baseMemory
	if len(treeId) == 0 {
		treeMem := b.getTreeMemory(treeId)
		mem = treeMem.memory
		if nodeId != "" {
			mem = b.getNodeMemory(treeId, nodeId)
		}
	}
	return mem
}

func (b *BlackBoard) getNodeMemory(treeId, nodeId string) *memory {

	return nil
}

// 获取树节点
func (b *BlackBoard) getTreeMemory(treeId string) *TreeMemory {
	if b.treeMemory[treeId] == nil {
		b.treeMemory[treeId] = newTreeMemory()
	}

	return b.treeMemory[treeId]
}

// 树状数据
func newTreeMemory() *TreeMemory {
	return &TreeMemory{}
}

func (b *BlackBoard) Set(key string, value interface{}, treeScope, nodeScope string) {
	mem := b.getMemory(treeScope, nodeScope)
	if mem != nil {
		mem.Set(key, value)
	}
}

func (b *BlackBoard) GetMemory(treeScope, nodeScope string) *memory {
	return nil
}

func (b *BlackBoard) Get(key string, treeScope, nodeScope string) interface{} {
	mem := b.GetMemory(treeScope, nodeScope)
	if mem != nil {
		return mem.Get(key)
	}
	return nil
}

// ---------------------------
//  Memory 取值
// ---------------------------

// 内部数据
type memory struct {
	mem map[string]interface{}
}

func newMemory() *memory {
	return &memory{
		mem: make(map[string]interface{}),
	}
}

// 自己转换
func (m *memory) Get(key string) interface{} {
	return m.mem[key]
}

func (m *memory) Set(key string, val interface{}) {
	m.mem[key] = val
}

func (m *memory) GetString(key string) string {
	str, ok := m.mem[key]
	if ok {
		return str.(string)
	}
	return ""
}

func (m *memory) GetBool(key string) bool {
	val, ok := m.mem[key]
	if ok {
		return val.(bool)
	}
	return false
}

func (m *memory) GetInt32(key string) int32 {
	val, ok := m.mem[key]
	if ok {
		return val.(int32)
	}
	return 0
}

func (m *memory) GetInt64(key string) int64 {
	val, ok := m.mem[key]
	if ok {
		return val.(int64)
	}
	return 0
}

func (m *memory) GetFloat64(key string) float64 {
	val, ok := m.mem[key]
	if ok {
		return val.(float64)
	}
	return 0
}
