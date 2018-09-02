package coll

// Contains determines whether the collection contains a given item
func (coll *Collection) Contains(f BoolFunc) bool {
	for _, v := range *coll {
		if f(v) {
			return true
		}
	}

	return false
}
