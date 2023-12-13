package input_config

import "sinhasoftware.solutions/kalakaar/pkg/common"

// Step is a struct that represents a step
type Step struct {
	common.CommonIdentifier
	// Command is the command to be executed
	Command string
}

// Pedal is a struct that represents a set of steps
type Pedal struct {
	common.CommonIdentifier
	Steps        []Step                    `json:"steps" yaml:"steps"`
	Dependencies []common.CommonIdentifier `json:"dependencies" yaml:"dependencies"`
}

// Setup is a struct that represents the setup for the pipeline. A setup can have a list of pedals.
// Setup should be a common setup for the pipeline e.g. setup a mock system that will be used by all the cycles.
// Setup will be enriched as we go along. Therefore, some features will be only restricted to Setup.
type Setup struct {
	// Pedals is a list of pedals
	Pedals []Pedal `json:"pedals" yaml:"pedals"`
}

// Cycle is a struct that represents a cycle in the pipeline. A cycle can have a list of pedals.
type Cycle struct {
	common.CommonIdentifier
	// Pedals is a list of pedals
	Pedals []Pedal `json:"pedals" yaml:"pedals"`
}

// Cycles is a struct that represents the cycles in the pipeline. Cycles is a list of cycles.
// It is recommended that you group as many pedals as you can into one cycle based on the phase of your pipeline.
// e.g. it is better to have 1 cycle called "Build and Deploy" than to have 2 cycles called "Build" and "Deploy",
// you can have another cycle called "Post Deployment Test", "Post Deployment Monitoring". Use multiple cycles only
// when Pedals cannot be used in a single cycle.
type Cycles struct {
	// Cycles is a list of cycles
	Cycles []Cycle `json:"cycles" yaml:"cycles"`
}

// Kalakaar is a struct that represents the pipeline. It can have a setup and multiple cycles
type Kalakaar struct {
	// Setup is the setup for the pipeline
	Setup Setup `json:"setup" yaml:"setup"`
	// Cycles is the cycles for the pipeline
	Cycles Cycles `json:"cycles" yaml:"cycles"`
}

// todo: think about how to inject secrets or use the current user to fetch secrets from a tpp etc.
