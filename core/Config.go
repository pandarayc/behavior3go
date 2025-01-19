package core

// 配置内相关属性

type BaseConfig struct {
	id          string // uuid
	name        string // name
	title       string
	description string
	properties  map[string]interface{}
}

func (config *BaseConfig) GetId() string {
	return config.id
}

func (node *BaseConfig) GetName() string {
	return node.name
}

func (node *BaseConfig) GetTitle() string {
	return node.title
}

func (node *BaseConfig) GetDescription() string {
	return node.description
}

func (node *BaseConfig) GetProperties() map[string]interface{} {
	return node.properties
}
