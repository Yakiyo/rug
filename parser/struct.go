package parser

// Struct representing a rug file (rug.json)
type RugFile struct {
	// Run pre/post hooks
	Hooks bool `json:"hooks"`
	// A map of environment variables
	Env map[string]string `json:"env,omitempty"`
	// paths to add to $PATH before running
	Path []string `json:"path,omitempty"`
	// List of scripts
	Scripts map[string]string `json:"scripts"`
}
