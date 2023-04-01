package main

import "testing"

func Test_miniHeight(t *testing.T) {
	type args struct {
		node *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := miniHeight(tt.args.node); got != tt.want {
				t.Errorf("miniHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
