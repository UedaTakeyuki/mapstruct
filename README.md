# mapstruct
Converting mutually between map[string]interface{} and struct.

# Functions
## ToStruct
```go
// Convert map[string]interface{} to a specific struct
func ToStruct[T any](d map[string]interface{}) (result *T, err error) 
```
##
```go
// Convert a specific struct to ap[string]interface{} 
func ToMap[T any](d *T) (result map[string]interface{}, err error) 
```

## example
```go
type PointType struct {
	User        interface{} `json:""`
	Type        byte        `json:""`
	CreatedTime int64       `json:""`
	Detail      DetailType  `json:",omitempty"`
}

type DetailType struct {
	Reason string `json:""`
}

var err error
point := map[string]interface{}{"User": "taro",
	"Type":        1,
	"CreatedTime": time.Now().Unix(),
	"Detail":      map[string]interface{}{"Reason": "gift"}}

var point2 *PointType
point2, err = mapstruct.ToStruct[PointType](point)

var point3 map[string]interface{}
point3, err = mapstruct.ToMap[PointType](point2)
```
for more detail, refer [mapstruct_test.go](https://github.com/UedaTakeyuki/mapstruct/blob/main/test/mapstruct_test.go)

## requirement
Go version 1.18 or later (using generics)
