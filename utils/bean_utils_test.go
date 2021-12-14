package utils

import (
	"fmt"
	"testing"
)

type User struct {
	Name *string
	Age  int
}

func (user *User) String() string {
	return fmt.Sprintf("{%s,%d}", *user.Name, user.Age)
}

func TestBeanCopy(t *testing.T) {
	str := "aaa"
	//str2 := "bbb"
	var a = &User{
		Name: &str,
		Age:  19,
	}
	var b = &User{}
	BeanCopy(a, b)
	fmt.Println(a, b)
	//now := time.Now().Format("2006-01-02 15:04:05")
}
