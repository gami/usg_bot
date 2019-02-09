package command

import (
	"strings"
)

var hellos = []string{"マンメンミ", "hello", "Hello", "おはよう"}

type Hello struct {
}

func (c *Hello) Resolve(m string) bool {
	for _, hello := range hellos {
		if strings.Contains(m, hello) {
			return true
		}
	}
	return false
}

func (c *Hello) Response(m string) string {
	return "マンメンミ！"
}
