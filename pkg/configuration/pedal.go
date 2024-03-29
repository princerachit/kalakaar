package configuration

import (
	"github.com/princerachit/kalakaar/pkg/common"
)

// Result is an interface that represents a result of some execution
type Result interface {
	// Get returns the result of the execution
	Get() interface{}
}

// Executable is an interface that represents an executable component in kalaakar
type Executable interface {
	// Execute executes the executable
	Execute() (Result, []error)
}

// IStep is an interface type that represents a step
type IStep interface {
	Executable
	common.ICommonIdentifier
}

// IPedal is an interface that represents a set of steps
type IPedal interface {
	IStep
	// AddStep adds a step to the pedal
	AddStep(step IStep)
	// ListSteps lists the steps in the pedal
	ListSteps() []IStep
	// GetDependencies returns the dependencies of the pedal
	// All dependencies must be executed before the pedal can be executed
	GetDependencies() []common.ICommonIdentifier
}
