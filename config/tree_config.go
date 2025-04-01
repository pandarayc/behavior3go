package config

// 解析配置使用
// https://www.behaviortrees.com/#/dash/home

// 树配置
type B3TreeCfg struct {
	NodeCfg
	Root          string              `json:"root"`
	Nodes         map[string]*NodeCfg `json:"nodes"`
	CustomNodeCfg []*CustomNodeCfg    `json:"custom_nodes"`
}

func (t *B3TreeCfg) ToProjectCfg() *B3ProjectCfg {
	return &B3ProjectCfg{
		Trees:        []*B3TreeCfg{t},
		SelectedTree: t.Id,
		CustomNodes:  t.CustomNodeCfg,
	}
}
