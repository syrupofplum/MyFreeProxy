package dynamodb

import (
	"fmt"
	"testing"
)

func TestListTables(t *testing.T) {
	accessor := Accessor{}
	accessor.Open()
	rst, err := ListTables(&accessor)
	fmt.Printf("ListTables rst, value: %v, err: %v\n", rst, err)
}
