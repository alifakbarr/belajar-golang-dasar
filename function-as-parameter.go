package main

import "fmt"

type Filter func(string) string
func sayHelloWithFilter(name string, filter Filter)  {

	 fmt.Println("hello", filter(name))
}

// func sayHelloWithFilter(name string, filter func(string) string)  {

// 	 fmt.Println("hello", filter(name))
// }

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}else{
		return name
	}
}

func main()  {
	sayHelloWithFilter("anjing",spamFilter)
}