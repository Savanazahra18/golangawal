package main

import "fmt"

// Struct Mahasiswa
type Mahasiswa struct {
	NPM   int
	Nama  string
	Kelas string
}

// Fungsi yang bekerja dengan struct Mahasiswa
func GreetMahasiswa(mhs Mahasiswa) string {
	return fmt.Sprintf("Halo, %s! (NPM: %d) dari kelas %s.", mhs.Nama, mhs.NPM, mhs.Kelas)
}

func main() {
	// Membuat instance dari struct Mahasiswa
	mahasiswa := Mahasiswa{
		NPM:   714230042,
		Nama:  "Savana Zahra Humaira",
		Kelas: "TI-2A",
	}

	// Memanggil fungsi dan mencetak hasil
	message := GreetMahasiswa(mahasiswa)
	fmt.Println(message)
}

