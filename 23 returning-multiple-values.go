package main

import "fmt"

func namaPanjang()(string, string,int){	
	return "Alif", "Akbar",21
}

func main(){
	firstName, lastName,_ := namaPanjang()

	fmt.Println(firstName, lastName)
}