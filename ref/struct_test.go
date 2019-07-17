package ref

import (
	"gitlab.3ag.xyz/backend/common/testutil"
	"reflect"
	"testing"
	"time"
)

func TestResetStruct(t *testing.T) {
	type args struct {
		sptr  interface{}
		replaceFields map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		panic bool
		want interface{}
	}{
		{
			name: "case 1",
			args: args{
				sptr: &struct{
					A int
					B int64
					C float32
					D float64
				}{
					A: 123,
					B: 123123123123123,
					C: .0,
					D: .3,
				},
				replaceFields: map[string]interface{}{
					"A": 0,
				},
			},
			want: &struct{
					A int
					B int64
					C float32
					D float64
				}{
					A: 0,
					B: 123123123123123,
					C: .0,
					D: .3,
				},
		},
		{
			name: "case 2 - can't access private field",
			args: args{
				sptr: &struct{
					a int
					B int64
					C float32
					D float64
				}{
					a: 123,
					B: 123123123123123,
					C: .0,
					D: .3,
				},
				replaceFields: map[string]interface{}{
					"a": 0,
				},
			},
			panic: true,
		},
		{
			name: "case 3 - not primitive data",
			args: args{
				sptr: &struct{
					A time.Time
				}{
					A: time.Now(),
				},
				replaceFields: map[string]interface{}{
					"A": time.Date(1900, 03, 14, 0, 0, 0, 0, time.Local),
				},
			},
			want:  &struct{
					A time.Time
				}{
					A: time.Date(1900, 03, 14, 0, 0, 0, 0, time.Local),
				},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				testutil.AssertPanic(t, func() {
					ResetStruct(tt.args.sptr, tt.args.replaceFields)
				})
			} else {
				ResetStruct(tt.args.sptr, tt.args.replaceFields)
				if !reflect.DeepEqual(tt.args.sptr, tt.want) {
					t.Errorf("StructKeys() = %v, want %v", tt.args.sptr, tt.want)
				}
			}
		})
	}
}

func TestStructKeys(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{
				struct {
					a int
					b int
				}{},
			},
			want: []string{"a", "b"},
		},
		{
			name: "case 2",
			args: args{
				&struct {
					a int
					b int
				}{},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructKeys(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}