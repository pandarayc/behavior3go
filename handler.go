package behavior3go

import (
	"github.com/pandarayc/behavior3go/actions"
	"github.com/pandarayc/behavior3go/composites"
	"github.com/pandarayc/behavior3go/core"
	"github.com/pandarayc/behavior3go/decorators"
)

var defaultRegisterHandlers *core.RegisterHandlers

func GetDefaultRegisterHandlers() *core.RegisterHandlers {
	if defaultRegisterHandlers == nil {
		handlers := core.NewRegisterHandlers()

		// actions
		handlers.Add("Error", &actions.Error{})
		handlers.Add("Failer", &actions.Failer{})
		handlers.Add("Runner", &actions.Runner{})
		handlers.Add("Succeeder", &actions.Succeeder{})
		handlers.Add("Wait", &actions.Wait{})

		// composites
		handlers.Add("Sequence", &composites.Sequence{})
		handlers.Add("MemSequence", &composites.MemSequence{})
		handlers.Add("Priority", &composites.Priority{})
		handlers.Add("MemPriority", &composites.MemPriority{})

		// decorators
		handlers.Add("Inverter", &decorators.Inverter{})
		handlers.Add("Limiter", &decorators.Limiter{})
		handlers.Add("MaxTime", &decorators.MaxTime{})
		handlers.Add("Repeater", &decorators.Repeater{})
		handlers.Add("RepeatUntilFailure", &decorators.RepeatUntilFailure{})
		handlers.Add("RepeatUntilSuccess", &decorators.RepeatUntilSuccess{})

		defaultRegisterHandlers = handlers
	}
	return defaultRegisterHandlers
}
