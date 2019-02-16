package entity

// Message represents Chat Message.
type Message struct {
	From     string
	To       []string
	Content  string
	Channel  string
	Category string
}
