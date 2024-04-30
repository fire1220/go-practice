package src

import (
	"context"
	"fmt"
	"go-unsafe/people"
	"reflect"
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

type ContextTemp struct {
	context.Context
	Key, Val any
}

func ContextKeys(ctx context.Context) {
	r := reflect.ValueOf(ctx)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	// reflect.Value{}
	temp := ContextTemp{}
	tr := reflect.ValueOf(&temp)
	if tr.Kind() == reflect.Ptr {
		tr = tr.Elem()
	}
	x := r.FieldByName("key")
	y := r.FieldByName("val")
	key := *(*any)(unsafe.Pointer(&x))
	val := *(*any)(unsafe.Pointer(&y))
	fmt.Println(key, val)
}

func ContextKeys1(ctx context.Context) {
	r := reflect.ValueOf(ctx)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	// reflect.Value{}
	x := r.Field(0).Interface().(context.Context)
	fmt.Printf("A:%#v\n", x)
	// fmt.Printf("C:%#v\n", unsafe.Pointer(&x))
	fmt.Printf("B:%#v\n", *(*any)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + 16)))
}
