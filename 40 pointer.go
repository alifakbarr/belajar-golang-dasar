package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func main(){
	address1 := Address{"subang", "jawa barat", "indonesia"}
	address2 := &address1

	address2.City = "bandung"
	
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := new(Address)
	fmt.Println(address3)

}