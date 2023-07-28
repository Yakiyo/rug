package parser

// Struct representing a rug file (rug.json)
type RugFile struct {
	// Run pre/post hooks
	Hooks bool
	// A map of environment variables
	Env map[string]string
	// paths to add to $PATH before running
	Path []string
	// List of scripts
	Scripts map[string]string
}
