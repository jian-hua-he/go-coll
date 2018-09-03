package coll

// Pop return a interface of the last Collection
// and change origin Collection.
func (coll *Collection) Pop() interface{} {
	len := len(*coll)
	if len == 0 {
		return nil
	}

	arr := ToInterfaces(*coll)
	*coll = arr[:len-1]

	return arr[len-1]
}
