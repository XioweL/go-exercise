package main

import "fmt"

func main() {
	people := []string{"Walt", "Jesse", "Skyler", "Saul"}

	fmt.Println("Panjang Slice people", (len(people)))

	people = append(people, "Hank", "Marrie")
	fmt.Println("Menambahkan", (people[4]), "dan", (people[5]))
	fmt.Println("Panjang Slice people", (len(people)))
	fmt.Println(people)
}
