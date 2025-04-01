package behavior3go

import (
	"testing"

	"github.com/pandarayc/behavior3go/core"
)

func TestXxx(t *testing.T) {
	mgr := New()
	mgr.LoadCfg("example/project2.json")
	// mgr.EchoConfig()
	err := 	mgr.Load()
	if err != nil {
		t.Fatal(err)
	}
	mgr.DumpTree()
}

func TestNewNode(t *testing.T) {
	name := "Wait"
	handlers := GetDefaultRegisterHandlers()
	node, err := handlers.New(name)
	if err != nil {
		t.Fatal(err)
	}
	node.SetWorker(node.(core.IWorker))
	_ = node
}
