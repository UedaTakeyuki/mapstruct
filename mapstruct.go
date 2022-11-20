// Converting mutually between map and struct.
//
// require: go 1.18 or later (generics)
//
// © Takeyuki Ueda.
package mapstruct

import "encoding/json"

func ToStruct[T any](d map[string]interface{}) (result *T, err error) {
	var jsonStr []byte
	if jsonStr, err = json.Marshal(d); err != nil {
		return
	} else {
		result = new(T)
		err = json.Unmarshal(jsonStr, result)
	}
	return
}

func ToMap[T any](d *T) (result map[string]interface{}, err error) {
	var jsonStr []byte
	if jsonStr, err = json.Marshal(d); err != nil {
		return
	} else {
		result = map[string]interface{}{}
		err = json.Unmarshal(jsonStr, &result)
	}
	return
}
