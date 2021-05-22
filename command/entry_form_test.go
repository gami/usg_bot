package command_test

import (
	"testing"

	"usg_bot/command"
)

func TestEntryForm_Resolve(t *testing.T) {
	type args struct {
		m string
	}
	tests := []struct {
		name string
		c    *command.EntryForm
		args args
		want bool
	}{
		{
			name: "OK",
			c:    &command.EntryForm{},
			args: args{
				m: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c
			if got := c.Resolve(tt.args.m); got != tt.want {
				t.Errorf("EntryForm.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
