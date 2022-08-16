package main

import "fmt"

func sumAll(numbers ...int)(int){
	total := 0
	for _,number := range numbers{
		total += number
	}

	return total
}

func main(){
	fmt.Println(sumAll(10,10,10,10))

	slice := []int{2,3,4,5}
	fmt.Println(sumAll(slice...))
}