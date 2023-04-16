package configuration

// Setup is a struct that represents the setup for the pipeline. A setup can have a list of pedals.
type Setup struct {
	// Pedals is a list of pedals
	Pedals []Pedal
}

// think about how to inject secrets or use the current user to fetch secrets from a tpp etc.
