package ref

import "reflect"

// sPrt: is a pointer of struct you want to reset
// replaceFields: is fields you want to reset, if you want to set field to zero value, just use nil represent.
func ResetStruct(sPtr interface{}, replaceFields map[string]interface{}) {

	switch reflect.ValueOf(sPtr).Kind() {
	case reflect.Ptr, reflect.UnsafePointer :
 		// correct, do nothing
	default:
		panic("sPtr must be a pointer")

	}
	keys := StructKeys(sPtr)

	v := reflect.Indirect(reflect.ValueOf(sPtr))
	for idx, k := range keys {
		if val, ok := replaceFields[k]; ok {
			fieldValue := v.Field(idx)

			if fieldValue.IsValid() {
				if !fieldValue.CanSet() {
					panic("can't reset value, field is private, can't be access")
				}
				switch val.(type) {
				case int:
					fieldValue.Set(reflect.ValueOf(val))
				case int64:
					fieldValue.SetInt(val.(int64))
				case float64:
					fieldValue.SetFloat(val.(float64))
				case string:
					fieldValue.SetString(val.(string))
				case nil:
					fieldValue.Set(reflect.Zero(reflect.TypeOf(val)))
				default:
					fieldValue.Set(reflect.ValueOf(val))
				}
			}
		}
	}
}

func StructKeys(s interface{}) []string {
	e := reflect.ValueOf(s)

	switch e.Kind() {
	case reflect.Ptr, reflect.UnsafePointer:
		e = e.Elem()
	case reflect.Struct:
		// NOTE do nothing
	default:
		panic("unexpect interface " + e.Kind().String())
	}

	var keys []string
	for i := 0 ; i < e.NumField() ; i++ {
		keys = append(keys, e.Type().Field(i).Name)
	}
	return keys
}