package main

import (
	"fmt"
	"reflect"
)

type ABC struct {
}

func main() {
	a := ABC{}
	r := reflect.TypeOf(a)
	fmt.Print(r.Name())
}
