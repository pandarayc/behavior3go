package behavior3go

import (
)

var defaultRegisterHandlers *RegisterHandlers

func GetDefaultRegisterHandlers() *RegisterHandlers {
	if defaultRegisterHandlers == nil {
		handlers := NewRegisterHandlers()

		// actions
		handlers.Add("Error", &Error{})
		handlers.Add("Failer", &Failer{})
		handlers.Add("Runner", &Runner{})
		handlers.Add("Succeeder", &Succeeder{})
		handlers.Add("Wait", &Wait{})

		// composites
		handlers.Add("Sequence", &Sequence{})
		handlers.Add("MemSequence", &MemSequence{})
		handlers.Add("Priority", &Priority{})
		handlers.Add("MemPriority", &MemPriority{})

		// decorators
		handlers.Add("Inverter", &Inverter{})
		handlers.Add("Limiter", &Limiter{})
		handlers.Add("MaxTime", &MaxTime{})
		handlers.Add("Repeater", &Repeater{})
		handlers.Add("RepeatUntilFailure", &RepeatUntilFailure{})
		handlers.Add("RepeatUntilSuccess", &RepeatUntilSuccess{})

		defaultRegisterHandlers = handlers
	}
	return defaultRegisterHandlers
}
