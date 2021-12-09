package utils

import (
	"errors"
)

func BeanCopy(src interface{}, dst interface{}) {
	if src == nil || dst == nil {
		panic(errors.New("src or dst can not be nil"))
	}
	//src_rv := reflect.TypeOf(src)
	//dst_rv := reflect.TypeOf(dst)

}
