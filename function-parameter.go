package main

import "fmt"

func sayHello(nama string,umur int){
	fmt.Println("halo :",nama,"Umur :",umur)
}

func main(){
	nama := "jojo"
	sayHello(nama,20)
	sayHello("irdho",21)
}