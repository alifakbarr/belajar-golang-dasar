package main

import "fmt"

func logging(){
	fmt.Println("Haloo")
}

func runApplication(){
	defer logging()
	fmt.Println("hai")
	
}

func main(){
	// defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
	// defer function akan selalu dieksekusi walaupun terjadi error difunction yang dieksekusi
	runApplication()
}