package main

import (
	"fmt"
	"time"
)

type Product struct {
	barcodeNumber string
	price         int64
	name          string
}

func (p *Product) information() string {
	return fmt.Sprintf("Informasi produk, nomor  barcode %s, harga: %d, nama: %s.", p.barcodeNumber, p.price, p.name)
}

type SamyangNoodle struct {
	Product
	variant     string
	isFried     bool
	color       string
	expiredDate time.Time
}

func (sn *SamyangNoodle) PrintInformation() {
	fmt.Printf("%s\n", sn.information())
	fmt.Printf("Varian: %s.\n", sn.variant)
	fmt.Printf("isFried: %s.\n", sn.isFried)
	fmt.Printf("color: %s.\n", sn.color)
	expired := sn.expiredDate
	expiredString := fmt.Sprintf("%v.%v.%v", expired.Day(), int(expired.Month()), expired.Year())
	fmt.Printf("Tanggal Kedaluarsa: %s.\n", expiredString)
}

func inheritance() {
	fmt.Println("====inheritance====")
	cheeseSamyang := SamyangNoodle{
		Product: Product{
			barcodeNumber: "1122333",
			price:         20000,
			name:          "Cheese Samyang Noodle",
		},
		variant:     "cheese",
		isFried:     true,
		color:       "Kuning",
		expiredDate: time.Now().Add(12 * 30 * 24 * time.Hour),
	}
	cheeseSamyang.PrintInformation()
}
