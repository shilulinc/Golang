package main

import (
	"fmt"
)

func main() {
	// var i  int = 10
	// i := 10
	// fmt.Println("i的内存地址", &i)

	var i = 999
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)
}
