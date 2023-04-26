package configuration

// Result is an interface that represents a result of some execution
type Result interface {
	// Get returns the result of the execution
	Get() interface{}
}

// Executable is an interface that represents an executable component in kalaakar
type Executable interface {
	// Execute executes the executable
	Execute() (Result, error)
}

// IStep is an interface type that represents a step
type IStep interface {
	Executable
	// GetId returns the id of the step
	GetId() string
	// SetId sets the id of the step
	SetId(id string) IStep
}

// IPedal is an interface that represents a set of steps
type IPedal interface {
	IStep
	// AddStep adds a step to the pedal
	AddStep(step IStep) IPedal
	// ListSteps lists the steps in the pedal
	ListSteps() []IStep
}
