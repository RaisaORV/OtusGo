package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	var k string = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(k))
}
