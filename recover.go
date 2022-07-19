package main

import "fmt"

func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("Eror karena",message)
	}
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
	// recover adalah function yang bisa kita gunakan untuk menangkap data panic
	// dengan recover proses panic akan terhenti, sehinggaprogram tetap berjalan
	
	runApp(true)
}