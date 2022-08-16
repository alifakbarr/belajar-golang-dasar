package main

import "fmt"

type Customer struct{
	nama, alamat string
	umur int
}

func main(){
	// struct adalah sebuah template data yang digunakanmenggabungkan nol atau lebih tipedata lainnya dalam kesatuan
	// struct basanya reperesentasi data dalam program yang kita buat
	// data distruct dismpan dalam field
	// ssederhananya struct adalahkumpulan field


	// cara mekakai struct 1
	var eko Customer
	eko.nama = "eko"
	eko.alamat ="indonesia"
	eko.umur = 30

	fmt.Println(eko)

	// cara mekakai struct 2
	joko := Customer{
		nama: "joko",
		alamat: "jakarta",
		umur: 30,
	}

	fmt.Println(joko)

	// cara mekakai struct 3
	budi := Customer{"budi","cirebon",20}

	fmt.Println(budi)
}