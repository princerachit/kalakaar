package configuration

import "github.com/princerachit/kalakaar/pkg/common"

// Cycle is a struct that represents a cycle in the pipeline. A cycle can have a list of pedals.
type Cycle struct {
	common.CommonIdentifier
	// Pedals is a list of pedals
	Pedals []IPedal
}

// Setup is a struct that represents the setup for the pipeline. A setup can have a list of pedals.
type Setup struct {
	// Pedals is a list of pedals
	Pedals []IPedal
}

// Pipeline is a struct that represents the pipeline. It can have a setup and multiple cycles
type Pipeline struct {
	// Setup is the setup for the pipeline
	Setup []IPedal
	// Cycles is the list of cycles for the pipeline
	Cycles []Cycle
}

// Validate validates the pipeline
func (p *Pipeline) Validate() []error {

	return nil
}
