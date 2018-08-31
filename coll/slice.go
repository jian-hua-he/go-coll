package coll

// Slice
func (coll *Collection) Slice(i int) Collection {
	arr := ToInterfaces(*coll)
	return arr[i:len(*coll)]
}

// SliceLen
func (coll *Collection) SliceLen(i int, j int) Collection {
	arr := ToInterfaces(*coll)
	return arr[i : i+j]
}
