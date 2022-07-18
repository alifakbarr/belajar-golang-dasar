package main

import "fmt"

func getName(name string)(string){
	return "heloo" + name
}

func main(){
	namaku := getName
	fmt.Println(namaku("Irdho"))
}

