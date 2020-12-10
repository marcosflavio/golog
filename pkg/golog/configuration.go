package golog

// GlogConfiguration represents the Glog client configuration
type GlogConfiguration struct {
	ApplicationID      string
	ApplicationVersion string
	ServerURL          string
	Local              bool
	Type               LoggerManagementType
	CustomData         map[string]interface{}
}
