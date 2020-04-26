package command

import (
	"strings"
	"usg_bot/entity"
)

// EntryForm represents a command for entry to this tournament.
type EntryForm struct {
}

// Resolve is a function to resolve chat message and check this command.
func (c *EntryForm) Resolve(m string) bool {
	if strings.HasPrefix(m, "!entry ") {
		return true
	}
	return false
}

// Response is a function to respond comamnd result.
func (c *EntryForm) Response(from *entity.User, msg *entity.Message) string {
	return ""
}
