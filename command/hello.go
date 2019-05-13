package command

import (
	"strings"
)

var hellos = []string{"マンメンミ", "hello", "Hello", "おはよう"}

// Hello represents a command which respond a greeting.
type Hello struct {
}

// Resolve is a function to resolve chat message and check this command.
func (c *Hello) Resolve(m string) bool {
	for _, hello := range hellos {
		if strings.Contains(m, hello) {
			return true
		}
	}
	return false
}

// Response is a function to respond comamnd result.
func (c *Hello) Response(m string) string {
	return "マンメンミ！"
}
