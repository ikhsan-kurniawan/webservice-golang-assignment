package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	NoPeserta   string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var listPeserta = []Biodata{
		{NoPeserta: "1", Nama: "Ikhsan Kurniawan", Alamat: "Purwokerto", Pekerjaan: "Mahasiswa", Alasan: "Belajar GoLang"},
		{NoPeserta: "2", Nama: "John Doe", Alamat: "Jakarta", Pekerjaan: "Pengembang", Alasan: "Minat dalam pemrograman"},
		{NoPeserta: "3", Nama: "Jane Doe", Alamat: "Surabaya", Pekerjaan: "Desainer", Alasan: "Ingin mengembangkan keterampilan desain"},
		{NoPeserta: "4", Nama: "Ahmad Ridwan", Alamat: "Bandung", Pekerjaan: "Peneliti", Alasan: "Penelitian dalam bidang ilmu komputer"},
		{NoPeserta: "5", Nama: "Siti Aminah", Alamat: "Yogyakarta", Pekerjaan: "Guru", Alasan: "Mengajar dan berbagi pengetahuan"},
		{NoPeserta: "6", Nama: "Budi Santoso", Alamat: "Semarang", Pekerjaan: "Wirausaha", Alasan: "Membangun bisnis teknologi"},
		{NoPeserta: "7", Nama: "Rina Dewi", Alamat: "Makassar", Pekerjaan: "Dokter", Alasan: "Menerapkan teknologi dalam bidang kesehatan"},
		{NoPeserta: "8", Nama: "Fahmi Rahman", Alamat: "Malang", Pekerjaan: "Mahasiswa", Alasan: "Mendalami pemrograman web"},
		{NoPeserta: "9", Nama: "Citra Indah", Alamat: "Denpasar", Pekerjaan: "Arsitek", Alasan: "Menggabungkan teknologi dalam desain arsitektur"},
		{NoPeserta: "10", Nama: "Dedy Pratama", Alamat: "Medan", Pekerjaan: "Analisis Data", Alasan: "Meneliti pola data untuk pengembangan bisnis"},
		{NoPeserta: "11", Nama: "Eka Putri", Alamat: "Palembang", Pekerjaan: "Pelajar", Alasan: "Belajar pemrograman untuk persiapan masa depan"},
		{NoPeserta: "12", Nama: "Farhan Yusuf", Alamat: "Aceh", Pekerjaan: "Fotografer", Alasan: "Menggabungkan teknologi dalam fotografi"},
		{NoPeserta: "13", Nama: "Gita Wijaya", Alamat: "Balikpapan", Pekerjaan: "Pengusaha", Alasan: "Mengembangkan startup teknologi"},
		{NoPeserta: "14", Nama: "Hendra Gunawan", Alamat: "Banjarmasin", Pekerjaan: "Programmer", Alasan: "Mengasah keterampilan pemrograman"},
		{NoPeserta: "15", Nama: "Ines Amelia", Alamat: "Bogor", Pekerjaan: "Pelajar", Alasan: "Belajar pemrograman untuk hobi"},
	}

	if len(os.Args) == 1 {
		fmt.Println("Masukan nomor peserta")
	} else if len(os.Args) == 2 {
		getPeserta(listPeserta, os.Args[1])
	} else {
		fmt.Println("Nomor peserta tidak valid")
	}

}

func getPeserta(peserta []Biodata, cari string) {
	var adaPeserta bool

	for _, value := range peserta {
		if cari == value.NoPeserta {
			adaPeserta = true
			fmt.Println("Nama:", value.Nama)
			fmt.Println("Alamat:", value.Alamat)
			fmt.Println("Pekerjaan:", value.Pekerjaan)
			fmt.Println("Alasan:", value.Alasan)
			break
		}
	}

	if !adaPeserta {
		fmt.Printf("Nomor peserta %s tidak ditemukan", cari)
	}

}