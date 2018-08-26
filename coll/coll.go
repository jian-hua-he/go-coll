package coll

import "reflect"

type Collection []interface{}

type BoolFunc func(v interface{}) bool

type BoolFuncWithIndex func(v interface{}, i int) bool

// Convert Collection to []interface{}
func ToInterfaces(coll Collection) []interface{} {
	return []interface{}(coll)
}

// Convert any type of slice to Collection type
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

// Return Collection length
func (coll *Collection) Len() int {
	return len(*coll)
}

// Append an element to Collection
func (coll *Collection) Push(v interface{}) {
	*coll = append(*coll, v)
}

// Return an element in Collection index
func (coll *Collection) Get(i int) interface{} {
	return ToInterfaces(*coll)[i]
}
