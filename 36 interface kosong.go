package main

import "fmt"

func Ups(i int) interface{}{
	if i == 1 {
		return 1
	}else{
		return "not 1"
	}
}

func main(){
	var data interface{} = Ups(1)
	fmt.Println(data)
}