package coll

func (coll *Collection) Contains(f BoolFunc) bool {
	for _, v := range *coll {
		if f(v) {
			return true
		}
	}

	return false
}
