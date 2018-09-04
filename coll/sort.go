package coll

// SortFunc given two parameters and return bool.
// Inside SortFunc determind the way to swap elements in Collection.
type SortFunc func(a interface{}, b interface{}) bool

// Sort given a SortFunc callback and return new sorted Collection.
func (coll *Collection) Sort(f SortFunc) {
	l := len(*coll)
	arr := ToInterfaces(*coll)

	for j := 0; j < l-1; j++ {
		for i := j; i < l-1; i++ {
			if f(arr[i], arr[i+1]) {
				temp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}

	*coll = arr
}
