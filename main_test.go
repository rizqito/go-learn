// untuk unit testing
package main

import (
	"belajar-golang/helper"
	"testing"
)

var (
	persegi      helper.Persegi = helper.Persegi{Sisi: 4} // objek persegi dari folder helper/interface.go
	luasExpected float64        = 16
)

func TestLuasPersegi(t *testing.T) {
	t.Logf("Luas Persegi : %v", persegi.Luas())

	if persegi.Luas() != luasExpected {
		t.Errorf("salah! seharusnya : %v", luasExpected)
	}
}
