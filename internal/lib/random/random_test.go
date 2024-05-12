package random

import (
	"testing"
	"unicode/utf8"
)

func TestNewRandomString(t *testing.T) {
	type args struct {
		length int
	}
	type want struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Length 1",
			args: args{
				length: 1,
			},
			want: want{
				length: 1,
			},
		},
		{
			name: "Length 10",
			args: args{
				length: 10,
			},
			want: want{
				length: 10,
			},
		},
		{
			name: "Length 100",
			args: args{
				length: 100,
			},
			want: want{
				length: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomString(tt.args.length); utf8.RuneCountInString(got) != tt.want.length {
				t.Errorf("NewRandomString() = %v, want %v", got, tt.want.length)
			}
		})
	}
}
