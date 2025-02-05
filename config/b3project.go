package config

import (
	"encoding/json"
	"os"
)

// b3.json
type B3ProjectCfg struct {
	Trees        []*TreeCfg `json:"trees"`
	CustomNodes  []*NodeCfg `json:"custom_nodes"`
	SelectedTree string     `json:"selectedTree"`
	Version      string     `json:"version"`
	Scope        string     `json:"scope"`
}

// 都转换至project级别的配置
func LoadB3ProjectCfg(path string) *B3ProjectCfg {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	cfg := &B3ProjectCfg{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
