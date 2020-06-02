package ref

import "testing"

func TestIsNotZero(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - int true",
			args: args{
				v: 0,
			},
			want: true,
		},
		{
			name: "success - int false",
			args: args{
				v: 100,
			},
			want: false,
		},
		{
			name: "success - string true",
			args: args{
				v: "",
			},
			want: true,
		},
		{
			name: "success - string false",
			args: args{
				v: "test",
			},
			want: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotZero(tt.args.v); got != tt.want {
				t.Errorf("IsNotZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
