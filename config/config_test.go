package config

import (
	"fmt"
	"testing"
)

func TestLoadB3ProjectCfg(t *testing.T) {
	path := "../project.json"
	cfg := LoadB3ProjectCfg(path)
	if cfg == nil {
		t.Error("load b3 project cfg failed")
	}
	fmt.Printf("cfg: %+v\n", cfg)

}
