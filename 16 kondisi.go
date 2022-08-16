package main

import "fmt"

func main(){
	nilai := 8

	if(nilai == 10){
		fmt.Println("lulus")
	}else if(nilai >= 8){
		fmt.Println("memuaskan")
	}else{
		fmt.Println("tidak lulus")
	}
	
	// variable temporary pada if -else
	
	if hasil := nilai / 2; hasil == 4 {
		fmt.Println("perfect ", hasil)
	}  else {
		fmt.Println("salah")
	}
}