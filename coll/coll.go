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

	arr := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		arr[i] = s.Index(i).Interface()
	}

	return Collection(arr)
}

func (coll *Collection) Len() int {
	return len(*coll)
}

func (coll *Collection) Push(v interface{}) {
	*coll = append(*coll, v)
}

func (coll *Collection) Get(i int) interface{} {
	return ToInterfaces(*coll)[i]
}
