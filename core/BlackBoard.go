package core

// 数据黑板
type BlackBoard struct {
	baseMemory *memory
	treeMemory map[string]*TreeMemory
}

// 新建黑板数据
func NewBlackBoard() *BlackBoard {
	blackBoard := &BlackBoard{
		baseMemory: newMemory(),
		treeMemory: make(map[string]*TreeMemory),
	}

	return blackBoard
}

// 获取树节点
func (b *BlackBoard) GetTreeMemory(treeId string) *TreeMemory {
	if b.treeMemory[treeId] == nil {
		b.treeMemory[treeId] = newTreeMemory()
	}

	return b.treeMemory[treeId]
}

func (b *BlackBoard) GetMemory(treeScope, nodeScope string) *memory {
	mem := b.baseMemory
	if len(treeScope) == 0 {
		treeMem := b.GetTreeMemory(treeScope)
		mem = treeMem.memory
		if nodeScope != "" {
			mem = treeMem.GetNodeMemory(nodeScope)
		}
	}
	return mem
}

func (b *BlackBoard) Get(key string, treeScope, nodeScope string) interface{} {
	mem := b.GetMemory(treeScope, nodeScope)
	if mem != nil {
		return mem.Get(key)
	}
	return nil
}

func (b *BlackBoard) Set(key string, value interface{}, treeScope, nodeScope string) {
	mem := b.GetMemory(treeScope, nodeScope)
	if mem != nil {
		mem.Set(key, value)
	}
}

// 作为共享信息层处理
type TreeMemory struct {
	nodeMemory map[string]*memory // 节点数据
	*memory                       // 树节点数据
}

// 树状数据
func newTreeMemory() *TreeMemory {
	return &TreeMemory{
		nodeMemory: make(map[string]*memory),
		memory:     newMemory(),
	}
}
func (tm *TreeMemory) GetNodeMemory(key string) *memory {
	if tm == nil {
		return nil
	}
	return tm.nodeMemory[key]
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
