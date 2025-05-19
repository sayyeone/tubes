package main

import (
	"fmt"
	"os"
	"strings"
)

// Array statis untuk menyimpan data siswa, maksimal 100 siswa
var data [100]TPerson
var count int = 0 // Variabel untuk menghitung jumlah siswa yang dimasukkan

// Struct untuk data siswa
type TPerson struct {
	id           int
	name         string
	username     string
	password     string
	age          int
	gender       string
	class        string
	schedule     string
	subject      string
	score1       int
	score2       int
	score3       int
	averageScore float64 // Menyimpan nilai rata-rata try out
}

// Fungsi untuk mencetak welcome page dan menu login
func printWelcome() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("   Selamat datang di Program BrainEd")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("Silahkan login sebagai:")
	fmt.Println("1. Admin")
	fmt.Println("2. Siswa")
	fmt.Println(strings.Repeat("=", 40))
}

// Fungsi untuk login sebagai admin atau siswa
func login(userType string) string {
	var username, password string
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	if userType == "admin" && username == "admin" && password == "admin123" {
		return "admin"
	} else if userType == "siswa" {
		// Cek jika username dan password siswa valid
		for _, person := range data[:count] {
			if person.username == username && person.password == password {
				return "siswa"
			}
		}
	}
	return "invalid"
}

// Fungsi untuk menu admin
func adminMenu() {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("     Menu Admin")
		fmt.Println("1. Input Siswa")
		fmt.Println("2. Daftar Siswa")
		fmt.Println("3. Input Nilai Siswa")
		fmt.Println("4. Hapus Siswa")
		fmt.Println("5. Edit Data Siswa")
		fmt.Println("6. Urutkan Siswa Berdasarkan Rata-rata Nilai Try Out")
		fmt.Println("7. Urutkan Siswa Berdasarkan ID")
		fmt.Println("8. Kembali ke Menu Sebelumnya")
		fmt.Println(strings.Repeat("=", 40))

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addSiswa()
		case 2:
			showSiswa()
		case 3:
			inputNilai()
		case 4:
			deleteSiswa()
		case 5:
			editSiswa()
		case 6:
			insertionSortByAvgScore() // Gunakan Insertion Sort untuk mengurutkan berdasarkan nilai rata-rata
			fmt.Println("Siswa diurutkan berdasarkan nilai rata-rata!")
		case 7:
			selectionSortByID() // Pilihan lain untuk pengurutan berdasarkan ID
			fmt.Println("Siswa diurutkan berdasarkan ID!")
		case 8:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}

// Fungsi untuk mengurutkan data berdasarkan ID (ascending) menggunakan Selection Sort
func selectionSortByID() {
	for i := 0; i < count-1; i++ {
		minIndex := i
		for j := i + 1; j < count; j++ {
			if data[j].id < data[minIndex].id {
				minIndex = j
			}
		}
		// Swap data[i] dengan data[minIndex]
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}

// Fungsi untuk menambahkan siswa dan menyimpan ke file txt
func addSiswa() {
	var person TPerson
	fmt.Print("Masukkan ID Siswa: ")
	fmt.Scanln(&person.id)
	fmt.Print("Masukkan Username Siswa: ")
	fmt.Scanln(&person.username)
	fmt.Print("Masukkan Password Siswa: ")
	fmt.Scanln(&person.password)
	fmt.Print("Masukkan Nama Siswa: ")
	fmt.Scanln(&person.name)
	fmt.Print("Masukkan Jenis Kelamin: ")
	fmt.Scanln(&person.gender)
	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&person.age)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&person.class)
	fmt.Print("Masukkan Jadwal Les: ")
	fmt.Scanln(&person.schedule)
	fmt.Print("Masukkan Mata Pelajaran: ")
	fmt.Scanln(&person.subject)
	fmt.Print("Masukkan Nilai Try Out 1: ")
	fmt.Scanln(&person.score1)
	fmt.Print("Masukkan Nilai Try Out 2: ")
	fmt.Scanln(&person.score2)
	fmt.Print("Masukkan Nilai Try Out 3: ")
	fmt.Scanln(&person.score3)

	// Tambahkan ke array data
	data[count] = person
	count++

	// Simpan ke file txt
	saveDataToFile()

	fmt.Println("Siswa berhasil ditambahkan!")
}

// Fungsi untuk menyimpan data siswa ke file txt
func saveDataToFile() {
	// Membuka atau membuat file data.txt
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Menulis data siswa ke file sebagai teks
	for i := 0; i < count; i++ {
		person := data[i]
		fmt.Fprintf(file, "%d|%s|%s|%s|%d|%s|%s|%s|%s|%d|%d|%d|%.2f\n", person.id, person.username, person.password, person.name, person.age, person.gender, person.class, person.schedule, person.subject, person.score1, person.score2, person.score3, person.averageScore)
	}

	fmt.Println("Data berhasil disimpan ke file txt!")
}

// Fungsi untuk membaca data siswa dari file txt
func loadDataFromFile() {
	file, err := os.Open("data.txt") // Membuka file data.txt
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	var temp []TPerson

	// Membaca data baris per baris
	var line string
	for {
		_, err := fmt.Fscanln(file, &line)
		if err != nil {
			break
		}

		// Memecah line berdasarkan delimiter "|"
		parts := strings.Split(line, "|")
		if len(parts) == 13 {
			var person TPerson
			person.id = stringToInt(parts[0])
			person.username = parts[1]
			person.password = parts[2]
			person.name = parts[3]
			person.age = stringToInt(parts[4])
			person.gender = parts[5]
			person.class = parts[6]
			person.schedule = parts[7]
			person.subject = parts[8]
			person.score1 = stringToInt(parts[9])
			person.score2 = stringToInt(parts[10])
			person.score3 = stringToInt(parts[11])
			person.averageScore = stringToFloat(parts[12])

			temp = append(temp, person)
		}
	}

	// Salin data ke dalam array statis
	for i, person := range temp {
		if i >= len(data) {
			break
		}
		data[i] = person
	}

	// Tentukan jumlah data yang terisi
	count = len(temp)

	fmt.Println("Data berhasil dimuat dari file txt!")
}

// Fungsi untuk mengubah string ke integer
func stringToInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

// Fungsi untuk mengubah string ke float
func stringToFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

// Fungsi untuk menampilkan daftar siswa
func showSiswa() {
	if count == 0 {
		fmt.Println("Tidak ada data siswa!")
		return
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%-5s%-20s%-10s%-10s%-15s%-20s\n", "ID", "Username", "Nama", "Umur", "Kelas", "Mata Pelajaran")
	for i := 0; i < count; i++ {
		person := data[i]
		fmt.Printf("%-5d%-20s%-10s%-10d%-15s%-20s\n", person.id, person.username, person.name, person.age, person.class, person.subject)
	}
	fmt.Println(strings.Repeat("=", 40))
}

// / Fungsi untuk menampilkan data siswa berdasarkan username
func showDataSiswa(username string) {
	// Cek username dengan cara yang lebih aman dan konsisten
	for _, person := range data {
		if strings.ToLower(person.username) == strings.ToLower(username) {
			// Tampilkan data siswa jika ditemukan
			fmt.Printf("Nama: %s\nJenis Kelamin: %s\nUmur: %d\nKelas: %s\nMata Pelajaran: %s\nJadwal: %s\n",
				person.name, person.gender, person.age, person.class, person.subject, person.schedule)
			return
		}
	}
	// Jika tidak ditemukan
	fmt.Println("Siswa tidak ditemukan!")
}

// Fungsi untuk menampilkan nilai siswa berdasarkan username
func showNilaiSiswa(username string) {
	// Cek username dengan cara yang lebih aman dan konsisten
	for _, person := range data {
		if strings.ToLower(person.username) == strings.ToLower(username) {
			// Tampilkan nilai try out siswa jika ditemukan
			fmt.Printf("Nilai Try Out 1: %d\nNilai Try Out 2: %d\nNilai Try Out 3: %d\n", person.score1, person.score2, person.score3)
			return
		}
	}
	// Jika tidak ditemukan
	fmt.Println("Siswa tidak ditemukan!")
}

// Fungsi untuk menampilkan peringkat berdasarkan nilai rata-rata
func showRankings() {
	// Menghitung rata-rata nilai try out untuk setiap siswa
	for i := 0; i < count; i++ {
		data[i].averageScore = float64(data[i].score1+data[i].score2+data[i].score3) / 3.0
	}

	// Mengurutkan siswa berdasarkan rata-rata nilai (ascending)
	insertionSortByAvgScore()

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%-20s%-10s%-10s\n", "Nama", "Rata-rata", "Peringkat")
	for rank, person := range data {
		fmt.Printf("%-20s%-10.2f%-10d\n", person.name, person.averageScore, rank+1)
	}
	fmt.Println(strings.Repeat("=", 40))
}

// Insertion Sort untuk mengurutkan siswa berdasarkan rata-rata nilai try out (ascending)
func insertionSortByAvgScore() {
	for i := 1; i < count; i++ {
		key := data[i]
		j := i - 1

		// Geser elemen yang lebih besar ke kanan
		for j >= 0 && data[j].averageScore > key.averageScore {
			data[j+1] = data[j]
			j = j - 1
		}
		// Tempatkan elemen key pada posisi yang benar
		data[j+1] = key
	}
}

// Fungsi untuk mencari jadwal les berdasarkan mata pelajaran
func showScheduleBySubject(username string) {
	var subject string
	fmt.Print("Masukkan Mata Pelajaran yang ingin dicari: ")
	fmt.Scanln(&subject)

	for _, person := range data {
		if person.username == username && person.subject == subject {
			fmt.Printf("Jadwal Les: %s\n", person.schedule)
			return
		}
	}
	fmt.Println("Jadwal les tidak ditemukan!")
}

// Fungsi untuk input nilai try out siswa
func inputNilai() {
	var username string
	fmt.Print("Masukkan Username Siswa: ")
	fmt.Scanln(&username)

	// Mencari siswa berdasarkan username dan input nilai
	for i := 0; i < count; i++ {
		if data[i].username == username {
			fmt.Println("Masukkan Nilai Try Out Baru:")
			fmt.Print("Nilai Try Out 1: ")
			fmt.Scanln(&data[i].score1)
			fmt.Print("Nilai Try Out 2: ")
			fmt.Scanln(&data[i].score2)
			fmt.Print("Nilai Try Out 3: ")
			fmt.Scanln(&data[i].score3)

			// Hitung ulang rata-rata nilai try out
			data[i].averageScore = float64(data[i].score1+data[i].score2+data[i].score3) / 3.0
			saveDataToFile() // Simpan perubahan ke file txt
			fmt.Println("Nilai berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Siswa tidak ditemukan!")
}

// Fungsi untuk menghapus siswa
func deleteSiswa() {
	var id int
	fmt.Print("Masukkan ID Siswa yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i := 0; i < count; i++ {
		if data[i].id == id {
			// Menggeser elemen setelahnya untuk menghapus data
			for j := i; j < count-1; j++ {
				data[j] = data[j+1]
			}
			count-- // Decrement jumlah siswa
			saveDataToFile()
			fmt.Println("Siswa berhasil dihapus!")
			return
		}
	}
	fmt.Println("Siswa tidak ditemukan!")
}

// Fungsi untuk mengedit data siswa berdasarkan ID
func editSiswa() {
	var id int
	fmt.Print("Masukkan ID Siswa yang ingin diedit: ")
	fmt.Scanln(&id)

	for i := 0; i < count; i++ {
		if data[i].id == id {
			// Edit data siswa
			fmt.Print("Masukkan Nama Siswa: ")
			fmt.Scanln(&data[i].name)
			fmt.Print("Masukkan Jenis Kelamin: ")
			fmt.Scanln(&data[i].gender)
			fmt.Print("Masukkan Umur: ")
			fmt.Scanln(&data[i].age)
			fmt.Print("Masukkan Kelas: ")
			fmt.Scanln(&data[i].class)
			fmt.Print("Masukkan Jadwal Les: ")
			fmt.Scanln(&data[i].schedule)
			fmt.Print("Masukkan Mata Pelajaran: ")
			fmt.Scanln(&data[i].subject)
			fmt.Print("Masukkan Nilai Try Out 1: ")
			fmt.Scanln(&data[i].score1)
			fmt.Print("Masukkan Nilai Try Out 2: ")
			fmt.Scanln(&data[i].score2)
			fmt.Print("Masukkan Nilai Try Out 3: ")
			fmt.Scanln(&data[i].score3)

			// Simpan perubahan ke file
			saveDataToFile()
			fmt.Println("Data siswa berhasil diubah!")
			return
		}
	}
	fmt.Println("Siswa tidak ditemukan!")
}

// Fungsi untuk keluar dari program
func exitProgram() {
	fmt.Println("Keluar dari program...")
	os.Exit(0)
}

// Fungsi untuk menu siswa
func siswaMenu(username string) {
	for {
		fmt.Println(strings.Repeat("=", 40))
		fmt.Println("     Menu Siswa")
		fmt.Println("1. Lihat Data Diri")
		fmt.Println("2. Lihat Nilai Try Out")
		fmt.Println("3. Lihat Peringkat Nilai Try Out")
		fmt.Println("4. Cari Jadwal Les")
		fmt.Println("5. Kembali ke Menu Sebelumnya")
		fmt.Println(strings.Repeat("=", 40))

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			showDataSiswa(username)
		case 2:
			showNilaiSiswa(username)
		case 3:
			showRankings()
		case 4:
			showScheduleBySubject(username)
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}

// Main function
func main() {
	// Memuat data dari file saat aplikasi dimulai
	loadDataFromFile()

	var choice int
	for {
		printWelcome()
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			result := login("admin")
			if result == "admin" {
				adminMenu()
			} else {
				fmt.Println("Username atau password salah!")
			}
		case 2:
			result := login("siswa")
			if result == "siswa" {
				var username string
				fmt.Print("Masukkan Username Siswa: ")
				fmt.Scanln(&username)
				siswaMenu(username)
			} else {
				fmt.Println("Username atau password salah!")
			}
		case 0:
			exitProgram()
		default:
			fmt.Println("Pilihan tidak valid, coba lagi!")
		}
	}
}
