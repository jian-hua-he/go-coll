package coll

func (coll *Collection) Splice(i int) Collection {
	arr := ToInterfaces(*coll)
	return arr[i:]
}

func (coll *Collection) SpliceLen(i int, len int) Collection {
	arr := ToInterfaces(*coll)
	return arr[i : i+len]
}
