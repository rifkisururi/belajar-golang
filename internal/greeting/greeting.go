// Package greeting builds a welcome message for a learner.
package greeting

import "strings"

// Message returns a greeting, defaulting to Gopher when name is blank.
func Message(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Gopher"
	}
	return "Hello, " + name + "!"
}
