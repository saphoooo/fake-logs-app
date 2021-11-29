package main

import (
	"testing"
)

func Test_randomize(t *testing.T) {
	type args struct {
		elements []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"foo", args{
			elements: []string{"foo"},
		}, "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomize(tt.args.elements); got != tt.want {
				t.Errorf("randomize() = %v, want %v", got, tt.want)
			}
		})
	}
}
