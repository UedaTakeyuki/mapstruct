# mapstruct
Converting mutually between map and struct.

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
refer [mapstruct_test.go](test/mapstruct.go)

## requirement
Go version 1.18 or later (using generics)
