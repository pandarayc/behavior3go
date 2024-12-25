package actions

import "github.com/genet9496/behavior3go/core"

type Wait struct {
	core.Action
	endTime int64
}

func OnOpen()
