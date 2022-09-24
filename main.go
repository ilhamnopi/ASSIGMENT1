package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	type Person struct {
		NoAbsen   int
		Nama      string
		Alamat   string
		Pekerjaan string
		Alasan    string
	}

	var persons []Person = []Person{
		{NoAbsen: 1, Nama: "Ilham Nopi Hendri", Alamat: "Solok", Pekerjaan: "Mahasiswa", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 2, Nama: "Hendri Muda", Alamat: "Medan", Pekerjaan: "Web Developer", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 3, Nama: "Robert", Alamat: "Jakarta Selatan", Pekerjaan: "Frontend Developer", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 4, Nama: "Joko", Alamat: "Bekasi", Pekerjaan: "Backend Deveeloper", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 5, Nama: "Budi", Alamat: "Solo", Pekerjaan: "Software Enginer", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 6, Nama: "Eko", Alamat: "Tebet", Pekerjaan: "Web Devloper", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 7, Nama: "herry", Alamat: "Medan", Pekerjaan: "Web Developer", Alasan: "-"},
		{NoAbsen: 8, Nama: "Jerry", Alamat: "Jakarta Selatan", Pekerjaan: "Frontend Developer", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 9, Nama: "Edly", Alamat: "Bekasi", Pekerjaan: "Backend Deveeloper", Alasan: "meningkatkan pengetahuan"},
		{NoAbsen: 10, Nama: "Narman", Alamat: "Solo", Pekerjaan: "Software Enginer", Alasan: "meningkatkan pengetahuan"},
	}
	
	args := os.Args[1]
	NoAbsenArg, _ := strconv.Atoi(args) //Konversikan String ke int
 for _, person := range persons{
		if person.NoAbsen == NoAbsenArg{
			fmt.Println("Menampilkan person dengan absen ke-", NoAbsenArg)
				fmt.Println("Nomor Absen : ", NoAbsenArg)
				fmt.Println("Nama",person.Nama)
				fmt.Println("Alamat", person.Alamat)
				fmt.Println("Pekerjaan", person.Pekerjaan)
				fmt.Println("Alasan", person.Alasan)
		}
	}
}