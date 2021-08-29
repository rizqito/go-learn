package main

import (
	f "fmt"
)

// interface untuk pelajar
type Pelajar interface {
	makan(makanan string)
	kuliah(mata_kuliah string)
}

// interface untuk asistenlab
type Pekerja interface {
	makan(makanan string)
	kuliah(mata_kuliah string)
	gajian()
}

// Manusia pada umumnya
type Manusia struct {
	nama, telepon, handphone, email string
	umur                            int
}

// Manusia berstatus Mahasiswa
type Mahasiswa struct {
	Manusia // embedded struct Manusia (semacam extend di OOP)
	nim     string
}

// Mahasiswa berstatus AsistenLab
type AsistenLab struct {
	Mahasiswa // embedded struct Mahasiswa (semacam extend di OOP)
	gaji      float64
}

// Semua Manusia pasti makan
func (m *Manusia) makan(makanan string) {
	f.Printf("Manusia Makan: %s \n", makanan)
}

// Semua Mahasiswa pasti kuliah
func (m *Mahasiswa) kuliah(mata_kuliah string) {
	f.Printf("Mahasiswa Kuliah Mata kuliah : %s \n", mata_kuliah)
}

// Semua AsistenLab pasti mengajar praktikum
func (as *AsistenLab) mengajar(praktikum string) {
	f.Printf("Asisten Lab mengajar praktikum : %s \n", praktikum)
}

// Semua AsistenLab pasti gajian
func (as *AsistenLab) gajian() {
	f.Printf("Asisten Lab gajian : Rp.%e \n", as.gaji)
}

func latihan1() {
	// implementasi struct

	f.Println("==========Mahasiswa============")
	// mahasiswa
	alex := Mahasiswa{
		Manusia: Manusia{
			nama:      "alex",
			telepon:   "082350790354",
			handphone: "+6282350790354",
			email:     "qito25@gmail.com",
			umur:      22,
		},
		nim: "1743065",
	}

	f.Printf("Nama      : %s\n", alex.nama)
	f.Printf("telepon   : %s\n", alex.telepon)
	f.Printf("handphone : %s\n", alex.handphone)
	f.Printf("email     : %s\n", alex.email)
	f.Printf("umur      : %v\n", alex.umur)
	f.Printf("nim       : %s\n", alex.nim)

	// interface
	var alex_mahasiswa Pelajar
	alex_mahasiswa = &alex
	alex_mahasiswa.makan("Nasi Campur")
	alex_mahasiswa.kuliah("Bahasa Pemrograman Golang")

	// Asisten Lab
	f.Println("==========Asisten Lab============")

	jonas := AsistenLab{
		Mahasiswa: Mahasiswa{
			Manusia: Manusia{
				nama:      "jonas",
				telepon:   "082350790354",
				handphone: "+6282350790354",
				email:     "jonas@gmail.com",
				umur:      23,
			},
			nim: "1743066",
		},
		gaji: 1500000,
	}

	f.Printf("Nama      : %s\n", jonas.nama)
	f.Printf("telepon   : %s\n", jonas.telepon)
	f.Printf("handphone : %s\n", jonas.handphone)
	f.Printf("email     : %s\n", jonas.email)
	f.Printf("umur      : %v\n", jonas.umur)
	f.Printf("nim       : %s\n", jonas.nim)
	f.Printf("gaji      : %v\n", jonas.gaji)

	var jonasAslab Pekerja
	jonasAslab = &jonas
	jonasAslab.makan("Ayam geprek")
	jonasAslab.kuliah("Golang Advance")
	jonasAslab.gajian()

}
