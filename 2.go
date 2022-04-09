package main

import "fmt"

func main() {
	var tanggal, tahun, bulan1, tanggal2, bulan2, tahun2 int
	var hari, bulan, hari2 string

	fmt.Scan(&hari, &tanggal, &bulan, &tahun)
	bulan1 = angkaBulan(bulan)
	tanggal2 = 0
	bulan2 = 0
	tahun2 = 0
	hari2 = ""
	pengambilan(tanggal, bulan1, tahun, hari, &tanggal2, &bulan2, &tahun2, &hari2)
	fmt.Println(hari2, tanggal2, bulanAngka(bulan2), tahun2)

}

func kabisat(tahun int) bool {
	return tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
}

func angkaBulan(bulan string) int {
	var hasil int
	if bulan == "januari" {
		hasil = 1
	} else if bulan == "februari" {
		hasil = 2
	} else if bulan == "maret" {
		hasil = 3
	} else if bulan == "april" {
		hasil = 4
	} else if bulan == "mei" {
		hasil = 5
	} else if bulan == "juni" {
		hasil = 6
	} else if bulan == "juli" {
		hasil = 7
	} else if bulan == "agustus" {
		hasil = 8
	} else if bulan == "september" {
		hasil = 9
	} else if bulan == "oktober" {
		hasil = 10
	} else if bulan == "november" {
		hasil = 11
	} else if bulan == "desember" {
		hasil = 12
	}
	return hasil
}

func bulanAngka(angka int) string {
	var hasil string
	if angka == 1 {
		hasil = "januari"
	} else if angka == 2 {
		hasil = "februari"
	} else if angka == 3 {
		hasil = "maret"
	} else if angka == 4 {
		hasil = "april"
	} else if angka == 5 {
		hasil = "mei"
	} else if angka == 6 {
		hasil = "juni"
	} else if angka == 7 {
		hasil = "juli"
	} else if angka == 8 {
		hasil = "agustus"
	} else if angka == 9 {
		hasil = "september"
	} else if angka == 10 {
		hasil = "oktober"
	} else if angka == 11 {
		hasil = "november"
	} else if angka == 12 {
		hasil = "desember"
	}
	return hasil
}

func jumlahHari(bulan, tahun int) int {
	var status bool
	var jumlah_hari int
	status = kabisat(tahun)
	if status && bulan == 2 {
		jumlah_hari = 29
	} else if bulan == 2 {
		jumlah_hari = 28
	} else if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		jumlah_hari = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		jumlah_hari = 30
	}
	return jumlah_hari
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	if hari1 == "senin" {
		*hari2 = "rabu"
		*durasi = 2
	} else if hari1 == "selasa" {
		*hari2 = "kamis"
		*durasi = 2
	} else if hari1 == "rabu" {
		*hari2 = "jumat"
		*durasi = 2
	} else if hari1 == "kamis" {
		*hari2 = "senin"
		*durasi = 4
	} else if hari1 == "jumat" {
		*hari2 = "selasa"
		*durasi = 4
	}
}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var durasi, jumlah_hari int
	var hari_pengambilan string
	durasi = 0
	hari_pengambilan = ""
	mencariDurasi(hari1, &hari_pengambilan, &durasi)
	jumlah_hari = jumlahHari(bulan1, tahun1)
	*tanggal2 = tanggal1 + durasi
	*bulan2 = bulan1
	*tahun2 = tahun1
	if *tanggal2 > jumlah_hari {
		*tanggal2 = *tanggal2 - jumlah_hari
		*bulan2 = *bulan2 + 1
	}

	if *bulan2 > 12 {
		*bulan2 = *bulan2 - 12
		*tahun2 = tahun1 + 1
	}
	*hari2 = hari_pengambilan

}
