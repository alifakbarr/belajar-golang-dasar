package main

import (
	"fmt"
	"errors"
)

func Pembagi(nilai int, pembagi int) (int, error){
	if pembagi == 0{
		return 0, errors.New("pembagi tidak boleh 0")
	}else{
		result := nilai/pembagi
		return result,nil
	}
}


func main(){
	hasil, err := Pembagi(100,0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	}else{
		fmt.Println("Error", err.Error())
	}
}