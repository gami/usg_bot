package command

import (
	"fmt"
	"strings"
	"usg_bot/entity"
)

// Me represents a status about Me
type Me struct {
}

// Resolve is a function to resolve chat message and check this command.
func (m *Me) Resolve(s string) bool {
	if s == "!me" {
		return true
	}
	if strings.HasPrefix(s, "!me ") {
		return true
	}
	return false
}

func (m *Me) Respond(from *entity.User, msg *entity.Message) string {
	return fmt.Sprintf("%v 君のDiscordIDは`%v`なんだって。イカしてるね！", from.Mention(), from.ID)
}
