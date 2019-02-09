package command

// Responder represents Commnad interface
type Responder interface {
	Resolve(string) bool
	Response(string) string
}
