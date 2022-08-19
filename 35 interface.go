package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	// interface adalah tipe data abstract, dia tidak memiliki implentasi langsung
	// sebuah interface yang berisikan definisi definisi method
	// biasanya interface digunakan sebagai kontrak

	// implementasi interface
	// setiap data yang yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
	// sehingga kita tidak perlu mengimplentasinya secara manual
	// hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface,kita harus menyebutkan secara eksplisit akan menggunakan interface yang sama
	var eko Person
	eko.Name = "EKo"
	SayHello(eko)

}
