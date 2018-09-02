package coll

// Filter method filters the Collection using the given BoolFunc,
// keeping only those items that pass a given truth test
func (coll *Collection) Filter(f BoolFunc) Collection {
	results := make(Collection, 0)
	for _, v := range *coll {
		if f(v) {
			results = append(results, v)
		}
	}

	return results
}

// FilterWithIndex method is almost the same thing with Filter function.
// The only difference is FilterWithIndex use BoolFuncWithIndex
// to determine which item should keep
func (coll *Collection) FilterWithIndex(f BoolFuncWithIndex) Collection {
	results := make(Collection, 0)
	for i, v := range *coll {
		if f(v, i) {
			results = append(results, v)
		}
	}

	return results
}
