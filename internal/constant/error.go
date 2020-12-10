package constant

// GlogError represents a type for Glog errors
type GlogError string

// Errors
const (
	ClientNotImplemented GlogError = "type.not.implemented"
	ConfigurationNotNil  GlogError = "configuration.not.nil"
)
