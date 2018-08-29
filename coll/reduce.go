package coll

type ReduceFunc func(c interface{}, i interface{}) interface{}

func (coll *Collection) Reduce(ReduceFunc) interface{} {
	return nil
}
