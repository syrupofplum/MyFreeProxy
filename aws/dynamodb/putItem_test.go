package dynamodb

import (
	"fmt"
	"testing"
)

func TestPutItem(t *testing.T) {
	accessor := Accessor{}
	accessor.Open()
	itemDict := make(map[string]interface{})
	itemDict["id"] = 1234

	err := PutItem(&accessor, "test", itemDict)
	fmt.Printf("PutItem rst, err: %v\n", err)
}

