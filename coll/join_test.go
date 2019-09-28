package coll

import "testing"

func TestJoinString(t *testing.T) {
	type args struct {
		v         []string
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 element",
			args: args{
				v: []string{"1"},
				delimiter: ",",
			},
			want: "1",
		},
		{
			name: "2 element",
			args: args{
				v: []string{"1", "2"},
				delimiter: ",",
			},
			want: "1,2",
		},
		{
			name: "3 element",
			args: args{
				v: []string{"1", "2", "3"},
				delimiter: ",",
			},
			want: "1,2,3",
		},
		{
			name: "2 element",
			args: args{
				v: []string{"agent_id = ?", "create_time >= ? AND create_time <= ?"},
				delimiter: " AND ",
			},
			want: "agent_id = ? AND create_time >= ? AND create_time <= ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinString(tt.args.v, tt.args.delimiter); got != tt.want {
				t.Errorf("JoinString() = %v, want %v", got, tt.want)
			}
		})
	}
}