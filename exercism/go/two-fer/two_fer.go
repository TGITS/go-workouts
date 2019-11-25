// The package twofer is an implementation of "Two-fer"
// "Two-fer" or "2-fer" is short for two for one. One for you and one for me.
package twofer

/*
The function ShareWith, given a string representing a name, returns a string with the message "One for X, one for me." where "X" is the given name.
However, if the variable name is missing, this is an empty string or a blank string, it returns the string:
"One for you, one for me."
*/
import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
