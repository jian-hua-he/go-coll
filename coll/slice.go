package coll

// Slice method returns a slice of the collection starting at the given index.
func (coll *Collection) Slice(i int) Collection {
	arr := ToInterfaces(*coll)
	return arr[i:len(*coll)]
}

// SliceLen method behavior is quite like Slice method.
// The only difference is pass the slice length to limit the size of return Collection.
func (coll *Collection) SliceLen(i int, j int) Collection {
	arr := ToInterfaces(*coll)
	return arr[i : i+j]
}
