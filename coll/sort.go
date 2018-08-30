package coll

// SortFunc
type SortFunc func(a interface{}, b interface{}) bool

// Sort
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
