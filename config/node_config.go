package config

import "encoding/json"

// 节点配置
type NodeCfg struct {
	Id          string                 `json:"id"`
	Name        string                 `json:"name"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Properties  map[string]interface{} `json:"properties"`
	// Display     map[string]int32       `json:"display"`
}

// 方便打印
func (t *NodeCfg) String() string {
	str, _ := json.Marshal(t)
	return string(str)
}
