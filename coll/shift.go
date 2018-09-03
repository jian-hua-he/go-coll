package coll

// Shift mehtod return the first element of Collection
// and change the Collection
func (coll *Collection) Shift() interface{} {
	len := len(*coll)
	if len == 0 {
		return nil
	}

	arr := ToInterfaces(*coll)
	*coll = arr[1:len]

	return arr[0]
}
