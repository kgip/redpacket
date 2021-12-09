package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBeanCopy(t *testing.T) {
	a := 1
	var b interface{} = &a
	brt := reflect.TypeOf(b)
	fmt.Println(brt.Kind())
	if brt.Kind() == reflect.Ptr {
		print("ptr")
	}
}
