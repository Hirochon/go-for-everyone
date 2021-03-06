package try_reflects

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func TryReflects() {
	p := &Point{
		X: 10,
		Y: 5,
	}
	rv := reflect.ValueOf(p).Elem()
	fmt.Println("=== Phase 1 ===")
	fmt.Printf("rv.Type = %v\n", rv.Type())
	fmt.Printf("rv.Kind = %v\n", rv.Kind())
	fmt.Printf("rv.Interface = %v\n", rv.Interface())

	fmt.Println("=== Phase 2 ===")
	xv := rv.Field(0)
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after SetInt) = %d\n", xv.Int())
}
