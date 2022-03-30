package gtools

import (
	"reflect"
)

//SliceUNIQ Slice elements to remove duplicates
func SliceUNIQ(data interface{}) interface{} {
	tOf := reflect.TypeOf(data)
	if tOf.Kind() != reflect.Slice {
		return data
	}
	if tOf.Elem().Comparable() == false {
		return data
	}

	vOf := reflect.ValueOf(data)
	if vOf.Len() == 0 {
		return data
	}
	if vOf.Index(0).CanInterface() == false {
		return data
	}
	var table = make(map[interface{}]struct{})
	sliceOfV := reflect.MakeSlice(tOf, 0, 0)
	for i := 0; i < vOf.Len(); i++ {
		tmp := vOf.Index(i).Interface()
		if _, ok := table[tmp]; ok {
			continue
		}
		table[tmp] = struct{}{}
		sliceOfV = reflect.Append(sliceOfV, reflect.ValueOf(tmp))
	}
	return sliceOfV.Interface()
}

//SliceExists  ele exist data
func SliceExists(data interface{}, ele interface{}) bool {
	tOf := reflect.TypeOf(data)
	if tOf.Kind() != reflect.Slice {
		return false
	}

	vOf := reflect.ValueOf(data)
	if vOf.Len() == 0 {
		return false
	}

	if vOf.Index(0).CanInterface() == false {
		return false
	}

	for i := 0; i < vOf.Len(); i++ {
		tmp := vOf.Index(i).Interface()
		if reflect.DeepEqual(tmp, ele) {
			return true
		}
	}

	return false
}
