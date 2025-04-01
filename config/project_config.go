package config

import (
	"encoding/json"
)

// b3.json
type B3ProjectCfg struct {
	Trees        []*B3TreeCfg     `json:"trees"`
	CustomNodes  []*CustomNodeCfg `json:"custom_nodes"`
	SelectedTree string           `json:"selectedTree"`
}

func (p *B3ProjectCfg) String() string {
	str, _ := json.Marshal(p)
	return string(str)
}
