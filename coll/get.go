package coll

func (coll Collection) Get(i int) interface{} {
	return ToInterfaces(coll)[i]
}
