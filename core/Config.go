package core

// 配置内相关属性
type IBaseConfig interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetProperties() map[string]interface{}
}
