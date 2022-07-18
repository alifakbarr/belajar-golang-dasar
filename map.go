package main

import "fmt"

func main(){
	// var chicken =map[string]int{}
	
	// chicken["januari"] = 50
	// fmt.Println(chicken["januari"]);
	
	
	var chicken =map[string]int{"januari":50, "februari":10}
	
	for key, val := range chicken{
		fmt.Println(key,":",val)
	}

}




