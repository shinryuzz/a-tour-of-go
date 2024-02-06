package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// i が T を保持していない場合、この文は panic を引き起こします。
	f = i.(float64) // panic 
	fmt.Println(f)
}
