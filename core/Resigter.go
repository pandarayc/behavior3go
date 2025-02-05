package core

var DefaultRegisterHandlers RegisterHandlers

func GetDefaultRegisterHandlers() RegisterHandlers {
	if DefaultRegisterHandlers == nil {
		DefaultRegisterHandlers = make(RegisterHandlers)
	}
	return DefaultRegisterHandlers
}



// 方法注册器
type RegisterHandlers map[string]INode

// Add 添加对应管理方法
func (h RegisterHandlers) Add(name string, node INode) {

}

