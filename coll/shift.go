package coll

func (coll *Collection) Shift() interface{} {
	len := len(*coll)
	if len == 0 {
		return nil
	}

	arr := ToInterfaces(*coll)
	*coll = arr[1:len]

	return arr[0]
}
