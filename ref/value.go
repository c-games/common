package ref

import "reflect"

func IsNotZero(v interface{}) bool {
	zv := reflect.Zero(reflect.TypeOf(v)).Interface()

	return reflect.DeepEqual(v, zv)
}
