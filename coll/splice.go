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
	r := make(Collection, len(arr[i:i+l]))
	copy(r, arr[i:i+l])

	p1 := make(Collection, len(arr[:i]))
	copy(p1, arr[:i])
	p2 := make(Collection, len(arr[i+l:]))
	copy(p2, arr[i+l:])

	p := make(Collection, 0, len(p1)+len(p2))
	p = append(p, p1...)
	p = append(p, p2...)
	*coll = p

	return r
}
