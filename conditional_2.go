package main

import "fmt"

func main() {
	var name string
	var umur int

	// Mencetak pesan untuk meminta input nama
	fmt.Print("Masukan nama Anda: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error: Input nama tidak valid")
		return
	}

	// Mencetak pesan untuk meminta input umur
	fmt.Print("Masukan Umur Anda: ")
	_, err = fmt.Scanln(&umur)

	switch {
	case err != nil:
		fmt.Println("Error: Input umur tidak valid")
	case umur < 0:
		fmt.Println("Umur Tidak boleh kurang dari 0")
	case umur > 100:
		fmt.Println("Umur Tidak boleh lebih dari 100")
	case umur > 18:
		fmt.Printf("Silahkan Masuk. %s\n", name)
	case umur <= 18:
		fmt.Printf("Dilarang masuk, minimal umur 19")
	default:
		fmt.Println("Silahkan Masuk")

	}
}
