package coll

import "reflect"

type Collection []interface{}

func ToInterfaces(coll Collection) []interface{} {
	return []interface{}(coll)
}

func Collect(slice interface{}) Collection {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("Collect: Parameter should be reflect.Slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return Collection(ret)
}

func Push(coll Collection, v interface{}) Collection {
	return append(coll, v)
}

func Pop(coll Collection) (interface{}, Collection) {
	len := len(coll)
	if len == 0 {
		return nil, coll
	}

	arr := ToInterfaces(coll)
	return arr[len-1], arr[:len-1]
}

func Shift(coll Collection) (interface{}, Collection) {
	len := len(coll)
	if len == 0 {
		return nil, coll
	}

	arr := ToInterfaces(coll)
	return arr[0], arr[1:len]
}
