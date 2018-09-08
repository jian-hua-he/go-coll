package coll

// Splice
func (coll *Collection) Splice(i int) Collection {
	arr := ToInterfaces(*coll)
	r := make(Collection, 0)
	r = append(r, arr[i:]...)

	p := make(Collection, 0)
	p = append(p, arr[:i]...)
	*coll = p

	return r
}

// SpliceLen
func (coll *Collection) SpliceLen(i int, l int) Collection {
	arr := ToInterfaces(*coll)
	r := make(Collection, 0)
	r = append(r, arr[i:i+l]...)

	p := make(Collection, 0)
	p = append(p, arr[:i]...)
	p = append(p, arr[i+l:]...)
	*coll = p

	return r
}
