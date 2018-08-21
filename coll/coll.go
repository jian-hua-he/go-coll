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
