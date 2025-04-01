package config

import (
	"encoding/json"
)

type IBaseConfig interface {
	GetId() string
	GetTitle() string
	GetDescription() string
	GetProperties() map[string]any
	GetProperty(name string) interface{}
	GetPropertyAsString(name string) string
	GetPropertyAsBool(name string) bool
	GetPropertyAsInt32(name string) int32
	GetPropertyAsInt64(name string) int64
	GetPropertyAsFloat32(name string) float32
	GetPropertyAsFloat64(name string) float64
}

// 节点配置
type NodeCfg struct {
	BaseConfig
	Children []string `json:"children"`
	Child    string   `json:"child"`
}

// 方便打印
func (t *NodeCfg) String() string {
	str, _ := json.Marshal(t)
	return string(str)
}
