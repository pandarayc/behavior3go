package config

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	SCOPE_PROJECT = "project"
	SCOPE_TREE    = "tree"
)

// 因为project级别的配置与Tree级别配置的结构体不一样，所以需要分开处理
type TestLoad struct {
	Scope   string `json:"scope"` // project or tree
	Version string `json:"version"`
}

func LoadConfig(path string) (*B3ProjectCfg, error) {
	// 读取文件
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	test := &TestLoad{}

	// json Test
	err = json.Unmarshal(data, &test)
	if err != nil {
		return nil, err
	}
	switch test.Scope {
	case SCOPE_PROJECT:
		// 读取project.json
		projectCfg := &B3ProjectCfg{}
		err = json.Unmarshal(data, &projectCfg)
		if err != nil {
			return nil, err
		}
		return projectCfg, nil
	case SCOPE_TREE:
		// 读取tree.json
		treeCfg := &B3TreeCfg{}
		err = json.Unmarshal(data, &treeCfg)
		if err != nil {
			return nil, err
		}
		return treeCfg.ToProjectCfg(), nil
	default:
		return nil, errors.New("invalid scope")
	}
}
