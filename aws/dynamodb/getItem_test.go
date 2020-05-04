package dynamodb

import (
	"fmt"
	"testing"
)

func TestGetItem(t *testing.T) {
	accessor := Accessor{}
	accessor.Open()
	itemDict := make(map[string]interface{})
	itemDict["id"] = 1234

	rst, err := GetItem(&accessor, "test", itemDict)
	fmt.Printf("PutItem rst: %v, err: %v\n", rst, err)
}


