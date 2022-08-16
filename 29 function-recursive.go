// function recursive = function yang memanggil dirinya sendiri
package main

import "fmt"

func factorial(value int) int{
	if value == 1{
		return 1
	}else{
		return value * factorial(value - 1)
	}
}

func main(){
	faktor := factorial(5)
	fmt.Println(faktor)
	fmt.Println(5*4*3*2*1)
}