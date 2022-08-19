package main

import (
	"fmt"
	"belajar-golang-dasar/database"
)

func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}