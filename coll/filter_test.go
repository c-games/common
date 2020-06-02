package coll

import (
	"reflect"
	"testing"
)

func Test_checkUserFieldsUpdatable(t *testing.T) {
	type args struct {
		updatableKeys []string
		data          map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "success",
			args: args{
				updatableKeys: []string{"a", "b", "c"},
				data: map[string]interface{}{
					"a": 123,
					"b": 456,
					"c": 879,
					"d": 10,
					"e": 11,
				},
			},
			want: map[string]interface{}{
				"a": 123,
				"b": 456,
				"c": 879,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterMapByKey(tt.args.updatableKeys, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterMapByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}