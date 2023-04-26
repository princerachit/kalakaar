package common

// CommonIdentifier is a struct that represents a common identifier for many components in kalaakar
type CommonIdentifier struct {
	// Name is the name of the component
	Name string `json:"name" yaml:"name"`
	// Id is the id of the component
	Id string `json:"id" yaml:"id"`
}

// ICommonIdentifier is an interface that represents a common identifier for many components in kalaakar
type ICommonIdentifier interface {
	// GetName returns the name of the component
	GetName() string
	// SetName sets the name of the component
	SetName(name string) CommonIdentifier
	// GetId returns the id of the component
	GetId() string
	// SetId sets the id of the component
	SetId(id string) CommonIdentifier
}
