package configuration

import "strings"

// DefaultPedal is the default implementation of the IPedal interface
type DefaultPedal struct {
	// steps is a list of IStep
	steps    []IStep
	id, name string
}

// AddStep adds a step to the pedal
func (p *DefaultPedal) AddStep(step IStep) {
	p.steps = append(p.steps, step)
}

// ListSteps lists the steps in the pedal
func (p *DefaultPedal) ListSteps() []IStep {
	return p.steps
}

// GetID returns the ID of the pedal
func (p *DefaultPedal) GetID() string {
	return p.id
}

// GetName returns the name of the pedal
func (p *DefaultPedal) GetName() string {
	return p.name
}

// SetID sets the ID of the pedal
func (p *DefaultPedal) SetID(id string) {
	p.id = id
}

// SetName sets the name of the pedal
func (p *DefaultPedal) SetName(name string) {
	p.name = name
}

// DefaultResult is the default implementation of the Result interface
type DefaultResult struct {
	// result is the result of the execution
	result []string
}

// Get returns the result of the execution
func (r *DefaultResult) Get() interface{} {
	return strings.Join(r.result, ",")
}

// Execute executes the pedal
func (p *DefaultPedal) Execute() (Result, []error) {
	results := make([]string, 0)
	for _, step := range p.steps {
		result, err := step.Execute()
		if err != nil {
			return nil, err
		}
		results = append(results, result.Get().(string))
	}
	return &DefaultResult{result: results}, nil
}
