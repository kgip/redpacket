package utils

import (
	"errors"
	"reflect"
)

func BeanCopy(src interface{}, dst interface{}, ignoreFields ...string) {
	if src == nil || dst == nil {
		panic(errors.New("src or dst can not be nil"))
	}
	srcRv := reflect.ValueOf(src)
	dstRv := reflect.ValueOf(dst)
	dstRt := reflect.TypeOf(dst).Elem()
	//拷贝的源对象和目标对象必须是指针
	if srcRv.Kind() != reflect.Ptr || dstRv.Kind() != reflect.Ptr || srcRv.Elem().Kind() != reflect.Struct || dstRv.Elem().Kind() != reflect.Struct {
		panic(errors.New("src or dst is not struct ptr"))
	}
	for i := 0; i < dstRv.Elem().NumField(); i++ {
		if !contains(ignoreFields, dstRt.Field(i).Name) {
			if v := srcRv.Elem().FieldByName(dstRt.Field(i).Name); !v.IsZero() {
				dstRv.Elem().Field(i).Set(v)
			}
		}
	}
}

func contains(list []string, item string) bool {
	for _, field := range list {
		if field == item {
			return true
		}
	}
	return false
}
