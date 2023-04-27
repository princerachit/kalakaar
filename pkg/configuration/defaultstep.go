package configuration

import (
	"os/exec"
)

// DefaultStep is the default implementation of the IStep interface
type DefaultStep struct {
	// id is the ID of the step
	id string
	// name is the name of the step
	name string
	// command is the shell command to execute
	command string
}

// GetID returns the ID of the step
func (s *DefaultStep) GetID() string {
	return s.id
}

// GetName returns the name of the step
func (s *DefaultStep) GetName() string {
	return s.name
}

// SetID sets the ID of the step
func (s *DefaultStep) SetID(id string) {
	s.id = id
}

// SetName sets the name of the step
func (s *DefaultStep) SetName(name string) {
	s.name = name
}

// Execute executes the step
func (s *DefaultStep) Execute() (Result, []error) {
	cmd := exec.Command("/bin/sh", "-c", s.command)
	out, err := cmd.Output()
	if err != nil {
		return nil, []error{err}
	}
	return &DefaultResult{result: []string{string(out)}}, nil
}
