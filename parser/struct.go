package parser

import "fmt"

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

func (r *RugFile) HasScript(name string) bool {
	_, ok := r.Scripts[name]
	return ok
}

func (r *RugFile) GetScript(name string) string {
	script, ok := r.Scripts[name]
	if !ok {
		return ""
	}
	return script
}

func (r *RugFile) GetEnv() []string {
	envs := []string{}
	for k, v := range r.Env {
		envs = append(envs, fmt.Sprintf("%v=%v", k, v))
	}
	return envs
}
