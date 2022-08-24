package main

import (
	"fmt"
	"package-inisilation/database"
)

func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}