package coll

// Splice
func (coll *Collection) Splice(i int) Collection {
	arr := ToInterfaces(*coll)
	*coll = arr[:i]
	return arr[i:]
}

// SpliceLen
func (coll *Collection) SpliceLen(i int, l int) Collection {
	arr := ToInterfaces(*coll)

	r := make([]interface{}, 0)
	r = append(r, arr[:i]...)
	r = append(r, arr[i+l:]...)

	*coll = r

	return arr[i : i+l]
}
