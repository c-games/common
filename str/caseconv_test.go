package str

import (
	"testing"
)

func TestIsUppercase(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				c: 'A',
			},
			want: true,
		},
		{
			args: args{
				c: 'Z',
			},
			want: true,
		},
		{
			args: args{
				c: 'a',
			},
			want: false,
		},
		{
			args: args{
				c: 'z',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUppercase(tt.args.c); got != tt.want {
				t.Errorf("IsUppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPascal2Snake(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "TestForPascal",
			},
			want: "test_for_pascal",
		},
		{
			args: args{
				str: "Test4Pascal",
			},
			want: "test_4_pascal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pascal2Snake(tt.args.str); got != tt.want {
				t.Errorf("Pascal2Snake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				c: '9',
			},
			want: true,
		},
		{
			args: args{
				c: '0',
			},
			want: true,
		},
		{
			args: args{
				c: 'a',
			},
			want: false,
		},
		{
			args: args{
				c: 'Z',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.c); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnake2Pascal(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "test_case",
			},
			want: "TestCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snake2Pascal(tt.args.str); got != tt.want {
				t.Errorf("Snake2Pascal() = %v, want %v", got, tt.want)
			}
		})
	}
}