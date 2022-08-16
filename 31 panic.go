package main

import "fmt"

func endApp(){
	fmt.Println("App selesai")
}

func runApp(error bool){
	defer endApp()
	
	if error{
		panic("Aplikasi error")
	}

	fmt.Println("APlikasi berjalan")

}

func main(){
	// panic function adalah function yang bisa kita gunakan untuk menghentikan program
	// panic function bisa dipanggil ketika terjadi error saat program kita berjalan
	// saat panic function dipanggil, program akan berhenti, namun defer function tetap dieksekusi
	
	runApp(true)
}