package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var name string
	fmt.Print("Masukan nama Anda: ")
	fmt.Scanln(&name)

	// rand.Intn(100) menghasilkan angka dari 0 hingga 99, jadi +1 untuk 1 hingga 100
	randomNumber := rand.Intn(100) + 1

	//fmt.Println("Random number:", randomNumber)

	// Switch Case Jika Mendapatkan Angka 1-100
	switch {
	case randomNumber > 80:
		fmt.Printf("Selamat Anda sangat beruntung %s,Anda Mendapatkan Angka %d", name, randomNumber)
	case randomNumber <= 80 && randomNumber > 60:
		fmt.Printf("Selamat Anda beruntung %s,Anda Mendapatkan Angka %d", name, randomNumber)
	case randomNumber <= 60 && randomNumber > 40:
		fmt.Printf("Mohon maaf %s Anda kurang beruntung, Anda Mendapatkan Angka %d", name, randomNumber)
	case randomNumber <= 40:
		fmt.Printf("Mohon maaf %s Anda sial, Anda Mendapatkan Angka %d", name, randomNumber)
	}
}
