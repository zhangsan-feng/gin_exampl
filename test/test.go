package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	fmt.Println(reflect.TypeOf(new(int)),
		reflect.TypeOf(new([]int)),
		reflect.TypeOf(make([]int, 10)),
		unsafe.Sizeof(struct{}{}),
		unsafe.Sizeof(&struct{}{}),
		unsafe.Sizeof(int32(1)),
		unsafe.Sizeof(int64(1)),
	)
	defer fmt.Println("11111")
	defer fmt.Println("22222")
	// func 2222 1111
}

func c(a int) {

}
