package coll

import (
	"reflect"
)

// CollectionMap 返回做過Group By 後的新type
type CollectionMap map[string]interface{}

// GroupBy method to build different groups with key
func (coll *Collection) GroupBy(groupKey string) []CollectionMap {

	// 紀錄有曾出現的groupby key
	var groupDataList []string

	// group by 之後的資料
	var returnMap []CollectionMap

	for _, collItem := range *coll {
		// reflect interface 型態的資料
		val := reflect.Indirect(reflect.ValueOf(collItem))

		// 每個collItem 有幾個field
		fieldLen := val.NumField()

		for i := 0; i < fieldLen; i++ {
			// 如果field是要做group by 的key
			if val.Type().Field(i).Name == groupKey {
				// 取得指定key的value
				groupByValue := val.FieldByName(groupKey)

				// 不在groupby key的紀錄(新的group處理)
				if StringInSlice(groupByValue.Interface().(string), groupDataList) == false {
					// 先記錄在一個slice
					groupDataList = append(groupDataList, groupByValue.Interface().(string))

					// 建立一個新Data，裡面包含groupKey和Object
					// groupKey 是群組的key
					// Object 是原本的資料
					newColl := make(CollectionMap, 0)
					newColl["groupKey"] = groupByValue.Interface().(string)

					// group 的資料用slice串接起來
					var arr []interface{}
					arr = append(arr, collItem)

					newColl["object"] = arr
					returnMap = append(returnMap, newColl)

				} else {
					// 在groupby key的紀錄
					for oKey, oValue := range returnMap {
						// 找到已經有group by 的群組，串接在最後面
						if oValue["groupKey"] == groupByValue.Interface().(string) {

							// group 的資料用slice串接起來
							var arr []interface{}
							arr = append(arr, collItem)

							returnMap[oKey]["object"] = append(returnMap[oKey]["object"].([]interface{}), arr)
						}
					}
				}
			}
		}

	}

	return returnMap
}

// StringInSlice 類似in_array()的功能
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
