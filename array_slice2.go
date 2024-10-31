package main

import "fmt"

func main() {
	dataSlice := make([]map[string]string, 0, 3) // Panjangnya 2 (len) Capasity 5 (cap)

	// Menambahkan Data Ke Dalam Array Dengan Append
	dataSlice = append(dataSlice, map[string]string{"name": "Hank", "gender": "M"})
	dataSlice = append(dataSlice, map[string]string{"name": "Heisenberg", "gender": "M"})
	dataSlice = append(dataSlice, map[string]string{"name": "Skyler", "gender": "F"})

	// Menggunakan for range untuk mengubah data array
	for _, person := range dataSlice { // for index, value := range collection { KODENYA SEPERTI INI
		if person["gender"] == "F" {
			person["name"] = "Mrs " + person["name"]
		} else if person["gender"] == "M" {
			person["name"] = "Mr " + person["name"]
		}
	}
	fmt.Println(dataSlice)
	// Cetak hasilnya dengan urutan yang spesifik
	for _, person := range dataSlice {
		fmt.Printf("map[name:%s | gender:%s]\n", person["name"], person["gender"])
	}
}
