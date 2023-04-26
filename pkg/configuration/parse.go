package configuration

import "gopkg.in/yaml.v2"

// Parse function accepts a byte array yaml file as input and unmarshals it to return a Kalakaar object
func Parse(data []byte) (*Kalakaar, error) {
	k := Kalakaar{}
	err := yaml.Unmarshal(data, &k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// Validate function accepts a Kalakaar object and validates it
func Validate(k *Kalakaar) []error {
	return nil
}
