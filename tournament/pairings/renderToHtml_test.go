package pairings

import (
	"fmt"
	"testing"
)

func TestFuncToHTML(t *testing.T) {
	fmt.Println("tests running")
	type args struct {
		p []P
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]P{}}, " "},
	}
	for _, tt := range tests {
		fmt.Println("tests running")
		t.Run(tt.name, func(t *testing.T) {
			if got := ToHTML(tt.args.p); got != tt.want {
				t.Errorf("ToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
