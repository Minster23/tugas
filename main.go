package main

import (
	"fmt"
)

type Pasien struct {
	nama, asal, tingkat  string
	tanggalLahir, antrian int
}

func tambahPasien(l []Pasien) *Pasien {
	var P Pasien
	fmt.Println("\n--- MASUKKAN DATA PASIEN ---")
	fmt.Print("Nama          : ")
	fmt.Scan(&P.nama)

	fmt.Print("Asal          : ")
	fmt.Scan(&P.asal)

	fmt.Print("Tanggal lahir : ")
	fmt.Scan(&P.tanggalLahir)

	fmt.Print("Tingkat penyakit (ringan/sedang/berat): ")
	fmt.Scan(&P.tingkat)

	P.antrian = len(l) + 1

	isValidTingkat := P.tingkat == "ringan" || P.tingkat == "sedang" || P.tingkat == "berat"

	if P == (Pasien{}) || !isValidTingkat {
		fmt.Println("\n[!] Data tidak terpenuhi atau salah, ulangi lagi.")
		P = Pasien{}
		tambahPasien(l)
	}
	fmt.Printf("\n[v] Pasien '%s' berhasil ditambahkan!\n", P.nama)
	return &P
}

func cariPasien(listPasien []Pasien) {
	var Value string
	fmt.Println("\n============= CARI PASIEN =============")
	fmt.Print("Masukkan keyword (Nama/Tingkat): ")
	fmt.Scan(&Value)

  for _, v := range listPasien {
		if Value == v.nama || Value == v.tingkat {
			fmt.Println("---------------------------------------")
			fmt.Printf("Antrian : %d\n", v.antrian)
			fmt.Printf("Nama    : %s\n", v.nama)
		} else {
			fmt.Println("[x] Pasien tidak ada")
		}
	}
	fmt.Println("=======================================")
}

func urutanPasien(l []Pasien) {
	var berat []Pasien
	var sedang []Pasien
	var biasa []Pasien

	for _, v := range l {
		switch v.tingkat {
		case "berat":
			berat = append(berat, v)
		case "sedang":
			sedang = append(sedang, v)
		default:
			biasa = append(biasa, v)
		}
	}

	var m []Pasien
	m = append(m, berat...)
	m = append(m, sedang...)
	m = append(m, biasa...)

	fmt.Println("\n+===============================================+")
	fmt.Println("|                 DAFTAR PASIEN                 |")
	fmt.Println("+===============================================+")
	fmt.Printf("| %-15s | %-9s | %-15s |\n", "Nama", "Antrian", "Tingkat")
	fmt.Println("+-----------------+-----------+-----------------+")
	for _, v := range m {
		fmt.Printf("| %-15s | %-9d | %-15s |\n", v.nama, v.antrian, v.tingkat)
	}
	fmt.Println("+===============================================+")
}

func Menu() {
	fmt.Println("\n+=====================================+")
	fmt.Println("|        SISTEM ANTRIAN KLINIK        |")
	fmt.Println("+=====================================+")
	fmt.Println("| 1 - Tambah Pasien                   |")
	fmt.Println("| 2 - Cari Pasien                     |")
	fmt.Println("| 3 - List Pasien                     |")
	fmt.Println("| 4 - Keluar                          |")
	fmt.Println("+=====================================+")
}

func main() {
	var listPasien []Pasien

	var Control int
	run := true
	for run {
		Menu()
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scan(&Control)
		switch Control {
		case 1:
			listPasien = append(listPasien, *tambahPasien(listPasien))
		case 2:
			cariPasien(listPasien)
		case 3:
			urutanPasien(listPasien)
		case 4:
      run = false
		}
	}
}
