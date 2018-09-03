package coll

// Remove method delete the item from Collection with index position.
func (coll *Collection) Remove(i int) {
	len := len(*coll)
	if len == 0 {
		return
	}

	if i > (len - 1) {
		return
	}

	arr := ToInterfaces(*coll)
	*coll = append(arr[:i], arr[i+1:]...)
}
