package src

import (
	"fmt"
	"go-reflect/people"
	"unsafe"
)

func GetPrivate(s people.People, key string) any {
	// r := reflect.ValueOf(s)
	// if r.Kind() == reflect.Pointer {
	// 	r = r.Elem()
	// }
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Offsetof(s.Name))
	fmt.Println(unsafe.Offsetof(s.Nick))
	user := people.People{Name: "jock"}
	user.SetLike("篮球")
	fmt.Printf("%#v\n", *(*any)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + 32)))
	// fmt.Printf("%#v\n", *(*any)(unsafe.Pointer(uintptr(unsafe.Pointer(&user)) + 16)))

	return nil
}
