package main

import (
	"fmt"
)

func printData(i interface{}) {
	var n, ok = i.(name)
	if ok {
		fmt.Println(n)
	}

	var a, ok1 = i.(age)
	if ok1 {
		fmt.Println(a)
	}

	var w, ok2 = i.(weight)
	if ok2 {
		fmt.Println(w)
	}
}

type name string
type age int
type weight float32

func main() {
	var n name = "Sandra"
	var a age = 19
	var w weight = 55.5

	printData(n)
	printData(a)
	printData(w)
}
