package config

import (
	"encoding/json"
)

// 解析配置使用
// https://www.behaviortrees.com/#/dash/home

// 树配置
type TreeCfg struct {
	Version    string                 `json:"version"`
	Scope      string                 `json:"scope"`
	Id         string                 `json:"id"`
	Title      string                 `json:"title"`
	Root       string                 `json:"root"`
	Properties map[string]interface{} `json:"properties"`
	Nodes      map[string]*NodeCfg    `json:"nodes"`
	Children   []string               `json:"children"`
	Child      string                 `json:"child"`
}

func (t *TreeCfg) String() string {
	str, _ := json.Marshal(t)
	return string(str)
}

