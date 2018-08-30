package coll

import "reflect"

// Collection
type Collection []interface{}

// BoolFunc
type BoolFunc func(v interface{}) bool

// BoolFuncWithIndex
type BoolFuncWithIndex func(v interface{}, i int) bool

// ToInterfaces convert Collection to []interface{}
func ToInterfaces(coll Collection) []interface{} {
	return []interface{}(coll)
}

// Collect convert any type of slice to Collection type
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

// Len return Collection length
func (coll *Collection) Len() int {
	return len(*coll)
}

// Push append an element to Collection
func (coll *Collection) Push(v interface{}) {
	*coll = append(*coll, v)
}

// Get return an element in Collection index
func (coll *Collection) Get(i int) interface{} {
	return ToInterfaces(*coll)[i]
}
