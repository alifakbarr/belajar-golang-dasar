package main

import "fmt"

func main(){
	x := 0
for i := 0; i < 2; i++ {
for j := 0; j < 2; j++ {
x += i + j
}
}
fmt.Println(x)
}