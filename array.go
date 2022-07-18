package main

import "fmt"

func main(){
	// var names [2]string

	// names[0] = "halo"
	// names[1] = "kalian"

	// fmt.Println(names[0],names[1])

	var buah = [2]string{"apel","pisang"}
	var angka = [...]int{1,2,3,4,5,6}

	fmt.Println(buah)

	for i := 0; i <= 1; i++ {
		fmt.Println(buah[i])
	}
	
	for _, angka := range angka{
		fmt.Println(angka)
	}
}

