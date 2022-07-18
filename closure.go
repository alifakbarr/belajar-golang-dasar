package main

import "fmt"

func main(){
	name := "budi"

	hello := func(){
		name := "eko"
		fmt.Println(name)
	}

	hello()
	fmt.Println(name)
}