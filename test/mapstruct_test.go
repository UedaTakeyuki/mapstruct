package mapstruct

import (
	"log"
	"testing"
	"time"

	cp "github.com/UedaTakeyuki/compare"

	"local.packages/mapstruct"
)

// type definition
type PointType struct {
	User        interface{} `json:""`
	Type        byte        `json:""`
	CreatedTime int64       `json:""`
	Detail      DetailType  `json:",omitempty"`
}

type DetailType struct {
	Reason string `json:""`
}

func Test_01(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	var err error
	point := map[string]interface{}{"User": "taro",
		"Type":        1,
		"CreatedTime": time.Now().Unix(),
		"Detail":      map[string]interface{}{"Reason": "gift"}}
	log.Println("point", point, point["Detail"])

	// test ToStruct
	var point2 *PointType
	point2, err = mapstruct.ToStruct[PointType](point)
	log.Println("point2", point2)
	cp.Compare(t, err, nil)
	cp.Compare(t, point2.User, point["User"])
	cp.Compare(t, point2.Type, uint8(point["Type"].(int)))
	cp.Compare(t, point2.CreatedTime, point["CreatedTime"])
	cp.Compare(t, point2.Detail.Reason, point["Detail"].(map[string]interface{})["Reason"])

	// test ToMap
	var point3 map[string]interface{}
	point3, err = mapstruct.ToMap[PointType](point2)
	log.Println("point3", point3)
	cp.Compare(t, err, nil)
	cp.Compare(t, point2.User, point3["User"])
	cp.Compare(t, point2.Type, uint8(point3["Type"].(float64)))
	cp.Compare(t, point2.CreatedTime, int64(point3["CreatedTime"].(float64)))
	cp.Compare(t, point2.Detail.Reason, point3["Detail"].(map[string]interface{})["Reason"])
}

func BenchmarkToStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapstruct.ToStruct[PointType](map[string]interface{}{"User": "taro",
			"Type":        1,
			"CreatedTime": time.Now().Unix(),
			"Detail":      map[string]interface{}{"Reason": "gift"}})
	}
}

func BenchmarkToMap(b *testing.B) {
	point := new(PointType)
	for i := 0; i < b.N; i++ {
		mapstruct.ToMap[PointType](point)
	}
}
