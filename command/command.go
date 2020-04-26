package command

import "usg_bot/entity"

// Responder represents Commnad interface
type Responder interface {
	Resolve(chat string) bool
	Respond(from *entity.User, msg *entity.Message) string
}
