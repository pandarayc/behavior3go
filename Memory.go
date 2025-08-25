package behavior3go

type IMemory interface {
	Exist(key string) bool
	Get(key string) interface{}
	Set(key string, val interface{})
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat32(key string) float32
	GetFloat64(key string) float64
}

var _ IMemory = &memory{}

// 内部数据
type memory struct {
	mem map[string]interface{}
}

func newMemory() *memory {
	return &memory{
		mem: make(map[string]interface{}),
	}
}

func (m *memory) Exist(key string) bool {
	_, ok := m.mem[key]
	return ok
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

func (m *memory) GetInt(key string) int {
	val, ok := m.mem[key]
	if ok {
		return val.(int)
	}
	return 0
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

func (m *memory) GetFloat32(key string) float32 {
	val, ok := m.mem[key]
	if ok {
		return val.(float32)
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
