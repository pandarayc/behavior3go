package actions

import "github.com/pandarayc/behavior3go/core"

type Wait struct {
	core.Action
	endTime int64
}

func OnOpen()
