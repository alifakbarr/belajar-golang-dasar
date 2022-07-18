package main

import "fmt"

func main(){
	var buah = []string{"apel","pisang"} //array
	fmt.Println(buah[0])
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
	
	var buahs = append(buah,"pepaya")
	fmt.Println(buah)
	fmt.Println(buahs)
	
}