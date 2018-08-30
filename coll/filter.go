package coll

// Filter
func (coll *Collection) Filter(f BoolFunc) Collection {
	results := make(Collection, 0)
	for _, v := range *coll {
		if f(v) {
			results = append(results, v)
		}
	}

	return results
}

// FilterWithIndex
func (coll *Collection) FilterWithIndex(f BoolFuncWithIndex) Collection {
	results := make(Collection, 0)
	for i, v := range *coll {
		if f(v, i) {
			results = append(results, v)
		}
	}

	return results
}
