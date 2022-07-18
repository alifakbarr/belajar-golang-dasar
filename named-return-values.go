package main

import "fmt"

func getFullName()(firstName, lastName string){
	firstName = "Alif"
	lastName = "Akbar"

	return
}

func main(){
	firstName, lastName := getFullName()

	fmt.Println(firstName,lastName)
}