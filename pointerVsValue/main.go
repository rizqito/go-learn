package main

import "fmt"

type Manusia struct {
	Nama         string
	JenisKelamin string
}

func (h Manusia) GantiJenisKelamin(JenisKelaminBaru string) {
	h.JenisKelamin = JenisKelaminBaru
}

func (h *Manusia) OperasiJenisKelamin(JenisKelaminBaru string) {
	h.JenisKelamin = JenisKelaminBaru
}

func main() {
	handsome := "Saya ganteng"
	//* get pointer
	pointerHandsome := &handsome
	//* ubah data via pointer
	*pointerHandsome = "saya cantik"
	//* get value pointer
	ganteng := *pointerHandsome
	fmt.Println(ganteng)

	//* method receiver
	andi := Manusia{Nama: "Andi", JenisKelamin: "Laki-laki"}
	fmt.Printf("%s mempunyai jenis kelamin : %s\n", andi.Nama, andi.JenisKelamin)

	andi.OperasiJenisKelamin("Perempuan")
	fmt.Printf("sekarang %s operasi jenis kelamin menjadi : %s\n", andi.Nama, andi.JenisKelamin)

	andi.GantiJenisKelamin("Laki-laki")
	fmt.Printf("%s mau jenis kelamin : %s\n", andi.Nama, andi.JenisKelamin)
}
