package tutil

import (
	"testing"
)

func TestRandNum(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "testRand",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandNum(); got == tt.want {
				t.Errorf("RandNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "testUrl",
			args: args{
				"https://github.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckURL(tt.args.url); got != tt.want {
				t.Errorf("CheckURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
