package main

import (
	"fmt"
)

func main() {
	var firstName, lastName string

	// Mencetak pesan untuk meminta input nama
	fmt.Print("Masukan nama: ")

	// Membaca input dari pengguna
	fmt.Scanln(&firstName, &lastName)

	// Menampilkan pesan
	fmt.Printf("hello %s %s\n", firstName, lastName)
}
