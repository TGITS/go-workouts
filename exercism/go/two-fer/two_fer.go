// Package twofer provides an implementation of "Two-fer" or "2-fer" which is short for "two for one. One for you and one for me.""
package twofer

import (
	"fmt"
)

// ShareWith returns a string with the message "One for X, one for me." where "X" is the value given in parameter if it is not the empty string or otherwise "me" .
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
