package main

import (
	"fmt"
)

const MAX = 100
const MAX_USER = 50

type Akun struct {
	Username string
	Password string
}

type Destinasi struct {
	ID     string
	Nama   string
	Lokasi string
	Rating float64
}

var data [MAX]Destinasi
var n int = 5

var akun [MAX_USER]Akun
var jumlahAkun int = 0

func registrasi() {
	if jumlahAkun >= MAX_USER {
		fmt.Println("Maksimal akun tercapai.")
		return
	}
	var u, p string
	fmt.Println("== REGISTRASI AKUN ==")
	fmt.Print("Username: ")
	fmt.Scan(&u)
	fmt.Print("Password: ")
	fmt.Scan(&p)

	for i := 0; i < jumlahAkun; i++ {
		if akun[i].Username == u {
			fmt.Println("Username sudah digunakan.")
			return
		}
	}

	akun[jumlahAkun] = Akun{u, p}
	jumlahAkun++
	fmt.Println("Akun berhasil dibuat!\n")
}

func login() bool {
	var u, p string
	fmt.Println("== LOGIN ==")
	fmt.Print("Username: ")
	fmt.Scan(&u)
	fmt.Print("Password: ")
	fmt.Scan(&p)

	for i := 0; i < jumlahAkun; i++ {
		if akun[i].Username == u && akun[i].Password == p {
			fmt.Println("Login berhasil!\n")
			return true
		}
	}

	fmt.Println("Username atau password salah.\n")
	return false
}

func menuAwal() bool {
	var pilih int
	for {
		fmt.Println("==== SELAMAT DATANG ====")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			registrasi()
		case 2:
			if login() {
				return true
			}
		case 0:
			return false
		default:
			fmt.Println("Pilihan tidak valid.\n")
		}
	}
}

func create(d Destinasi) {
	if n >= MAX {
		fmt.Println("Data penuh.")
		return
	}
	data[n] = d
	n++
	fmt.Println("Data berhasil ditambahkan!")
}

func read() {
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	fmt.Println("\nData Destinasi:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID: %s | Nama: %s | Lokasi: %s | Rating: %.2f\n", i+1, data[i].ID, data[i].Nama, data[i].Lokasi, data[i].Rating)
	}
}

func update(id string) {
	for i := 0; i < n; i++ {
		if data[i].ID == id {
			fmt.Print("Nama baru: ")
			fmt.Scan(&data[i].Nama)
			fmt.Print("Lokasi baru: ")
			fmt.Scan(&data[i].Lokasi)
			fmt.Print("Rating baru: ")
			fmt.Scan(&data[i].Rating)
			fmt.Println("Data berhasil diupdate.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func deleteData(id string) {
	for i := 0; i < n; i++ {
		if data[i].ID == id {
			for j := i; j < n-1; j++ {
				data[j] = data[j+1]
			}
			n--
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func search(id string) int {
	for i := 0; i < n; i++ {
		if data[i].ID == id {
			return i
		}
	}
	return -1
}

func sortRating() {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j].Rating < data[j+1].Rating {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan rating secara descending.")
}

func tampilkanRatingEkstrem() {
	if n == 0 {
		fmt.Println("Data kosong.")
		return
	}
	max := data[0]
	min := data[0]
	for i := 1; i < n; i++ {
		if data[i].Rating > max.Rating {
			max = data[i]
		}
		if data[i].Rating < min.Rating {
			min = data[i]
		}
	}
	fmt.Printf("Rating Tertinggi: %s (%.2f)\n", max.Nama, max.Rating)
	fmt.Printf("Rating Terendah: %s (%.2f)\n", min.Nama, min.Rating)
}

func hitungRataRata() {
	if n == 0 {
		fmt.Println("Data kosong.")
		return
	}
	var total float64
	for i := 0; i < n; i++ {
		total += data[i].Rating
	}
	rata := total / float64(n)
	fmt.Printf("Rata-rata rating semua destinasi: %.2f\n", rata)
}

func init() {
	data[0] = Destinasi{"D001", "Tangkuban Perahu", "Bandung", 3.0}
	data[1] = Destinasi{"D002", "Taman Nasional Gunung Ciremai", "Kuningan", 4.6}
	data[2] = Destinasi{"D003", "Taman Nasional Gunung Merbabu", "Magelang", 5.0}
	data[3] = Destinasi{"D004", "Gunung Gede", "Bogor", 4.9}
	data[4] = Destinasi{"D005", "Gunung Papandayan", "Garut", 4.2}
}

func main() {
	if !menuAwal() {
		fmt.Println("Program keluar.")
		return
	}

	var pilihan int
	var id string

	for {
		fmt.Println("\n===== MENU UTAMA =====")
		fmt.Println("1. Tambah destinasi")
		fmt.Println("2. Tampilkan semua destinasi")
		fmt.Println("3. Edit destinasi")
		fmt.Println("4. Hapus destinasi")
		fmt.Println("5. Cari berdasarkan ID")
		fmt.Println("6. Tampilkan rating tertinggi dan terendah")
		fmt.Println("7. Urutkan berdasarkan rating")
		fmt.Println("8. Hitung rata-rata rating")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var d Destinasi
			fmt.Print("ID: ")
			fmt.Scan(&d.ID)
			fmt.Print("Nama: ")
			fmt.Scan(&d.Nama)
			fmt.Print("Lokasi: ")
			fmt.Scan(&d.Lokasi)
			fmt.Print("Rating (0.0 - 5.0): ")
			fmt.Scan(&d.Rating)
			if d.Rating < 0.0 || d.Rating > 5.0 {
				fmt.Println("Rating tidak valid.")
				continue
			}
			create(d)
		case 2:
			read()
		case 3:
			fmt.Print("Masukkan ID destinasi yang ingin diedit: ")
			fmt.Scan(&id)
			update(id)
		case 4:
			fmt.Print("Masukkan ID destinasi yang ingin dihapus: ")
			fmt.Scan(&id)
			deleteData(id)
		case 5:
			fmt.Print("Masukkan ID yang dicari: ")
			fmt.Scan(&id)
			index := search(id)
			if index != -1 {
				fmt.Println("Data ditemukan:")
				fmt.Printf("ID: %s | Nama: %s | Lokasi: %s | Rating: %.2f\n", data[index].ID, data[index].Nama, data[index].Lokasi, data[index].Rating)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case 6:
			tampilkanRatingEkstrem()
		case 7:
			sortRating()
		case 8:
			hitungRataRata()
		case 0:
			fmt.Println("Terima kasih. Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
