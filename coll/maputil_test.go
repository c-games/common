package coll

import (
	"reflect"
	"testing"
)

func TestGetKepsFromMap(t *testing.T) {
	type args struct {
		ori_map map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "pass case",
			args: args{
				ori_map: map[string]interface{}{
					"first": 123,
					"third": "test",
					"second": 1000.0,
				},
			},
			// Note: because result has been sorted by sort.Strings
			want: []string{"first", "second", "third"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetKeysFromMap(tt.args.ori_map); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeysFromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}