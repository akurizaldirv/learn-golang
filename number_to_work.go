package main

import (
	"fmt"
	"strings"
)

func main() {
	number := 100
	kalimat := convertNumber(number)
	fmt.Println(kalimat)
}

func convertNumber(number int) string {
	satuan := []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}
	var hasil string

	if number < 12 {
		hasil = satuan[number]
	} else if number < 20 {
		hasil = convertNumber(number-10) + " belas"
	} else if number < 100 {
		hasil = convertNumber(number/10) + " puluh " + convertNumber(number%10)
	} else if number < 200 {
		hasil = "seratus " + convertNumber(number-100)
	} else if number < 1000 {
		hasil = convertNumber(number/100) + " ratus " + convertNumber(number%100)
	} else if number < 2000 {
		hasil = "seribu " + convertNumber(number-1000)
	} else if number < 1000000 {
		hasil = convertNumber(number/1000) + " ribu " + convertNumber(number%1000)
	} else if number < 1000000000 {
		hasil = convertNumber(number/1000000) + " juta " + convertNumber(number%1000000)
	} else if number < 1000000000000 {
		hasil = convertNumber(number/1000000000) + " miliar " + convertNumber(number%1000000000)
	} else if number < 1000000000000000 {
		hasil = convertNumber(number/1000000000000) + " triliun " + convertNumber(number%1000000000000)
	} else {
		hasil = "angka terlalu besar"
	}

	hasil = strings.TrimSpace(hasil)
	return hasil
}
