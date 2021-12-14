package utils

import (
	"errors"
	"reflect"
)

func BeanCopy(src interface{}, dst interface{}) {
	if src == nil || dst == nil {
		panic(errors.New("src or dst can not be nil"))
	}
	srcRv := reflect.ValueOf(src)
	dstRv := reflect.ValueOf(dst)
	srcRt := reflect.TypeOf(src).Elem()
	dstRt := reflect.TypeOf(dst).Elem()
	//拷贝的源对象和目标对象必须是指针
	if srcRv.Kind() != reflect.Ptr || dstRv.Kind() != reflect.Ptr || srcRv.Elem().Kind() != reflect.Struct || dstRv.Elem().Kind() != reflect.Struct {
		panic(errors.New("src or dst is not struct ptr"))
	}
	for i := 0; i < srcRv.Elem().NumField(); i++ {
		for j := 0; j < dstRv.Elem().NumField(); j++ {
			if srcRt.Field(i).Name == dstRt.Field(j).Name && srcRv.Elem().Field(i).Kind() == dstRv.Elem().Field(j).Kind() {
				dstRv.Elem().Field(j).Set(srcRv.Elem().Field(i))
			}
		}
	}
}
