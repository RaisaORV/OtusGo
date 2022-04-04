package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	k := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(k))
}
