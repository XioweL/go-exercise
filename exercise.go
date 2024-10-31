package main

import "fmt"

type Pegawai struct {
	ID, Umur int
	Nama     string
	Gaji     float64
}

func TambahPegawai(id int, nama string, umur int, gaji float64) Pegawai {
	return Pegawai{
		ID:   id,
		Umur: umur,
		Nama: nama,
		Gaji: gaji,
	}
}

func CariPegawai(listPegawai []Pegawai) {
	fmt.Println("List Pegawai")
	for _, pegawai := range listPegawai {
		fmt.Printf("Data Pegawai: %d, Nama: %s, Umur: %d, Gaji: %.1f \n", pegawai.ID, pegawai.Nama, pegawai.Umur, pegawai.Gaji)
	}
}

func NaikGaji(ListPegawai []Pegawai) {

}
func main() {
	var ListPegawai []Pegawai

	pegawai1 := TambahPegawai(1, "Ferdian", 25, 5000)
	pegawai2 := TambahPegawai(2, "Ucok", 24, 6000)
	pegawai3 := TambahPegawai(3, "Mamat", 22, 4000)

	ListPegawai = append(ListPegawai, pegawai1, pegawai2, pegawai3)

	CariPegawai(ListPegawai)

}
