package main

import "fmt"

const NMAX = 100
type projects [NMAX]proyek
type proyek struct {
	nama string
	kategori string
	teknologi [10]string
	jTeknologi int
	kesulitan int
	waktu string
}

func main() {
	var p projects
	var n int
	menu(&p, &n)
}

func menu(projects *proyek, n *int) {
	for {
		fmt.Println("\n=== Aplikasi Manajemen Portofolio Data Science ===")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tampilkan Semua Proyek")
		fmt.Println("3. Ubah Proyek")
		fmt.Println("4. Hapus Proyek")
		fmt.Println("5. Cari Proyek (Sequential Search)")
		fmt.Println("6. Cari Proyek (Binary Search)")
		fmt.Println("7. Urutkan Proyek (Selection Sort)")
		fmt.Println("8. Urutkan Proyek (Insertion Sort)")
		fmt.Println("9. Statistik Proyek per Kategori")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahProyek(projects, n)
		case 2:
			tampilkanProyek(*projects, *n)
		case 3:
			ubahProyek(projects, n)
		case 4:
			hapusProyek(projects, n)
		case 5:
			sequentialSearch(*projects, *n)
		case 6:
			binarySearch(projects, *n)
		case 7:
			selectionSort(projects, *n)
			tampilkanProyek(*projects, *n)
		case 8:
			insertionSort(projects, *n)
			tampilkanProyek(*projects, *n)
		case 9:
			statistikKategori(*projects, *n)
		case 0:
			fmt.Println("Keluar dari aplikasi.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahProyek(projects *proyek, n *int) {
	var p Project
	
	if *n >= NMAX {
		fmt.Println("Kapasitas proyek penuh.")
	}

	fmt.Print("Nama Proyek: ")
	fmt.Scan(&p.nama)
	fmt.Print("Kategori (contoh: machine_learning): ")
	fmt.Scan(&p.kategori)
	fmt.Print("Jumlah Teknologi (maks 10): ")
	fmt.Scan(&p.jTeknologi)

	if p.jTeknologi > 10 {
		fmt.Println("Maksimum 10 teknologi!")
		p.jTeknologi = 10
	}

	for i := 0; i < p.jTeknologi; i++ {
		fmt.Printf("Teknologi ke-%d: ", i+1)
		fmt.Scan(&p.teknologi[i])
	}
	
	fmt.Print("Tingkat Kesulitan (1-5): ")
	fmt.Scan(&p.kesulitan)
	fmt.Print("Tanggal Pembuatan (dd-mm-yyy): ")
	fmt.Scan(&p.waktu)

	projects[*n] = p
	*n++
	fmt.Println("Proyek berhasil ditambahkan!")
}

func tampilkanProyek(projects Project, n int) {
	fmt.Println("\n=== Daftar Proyek ===")
	
	for i := 0; i < count; i++ {
		p := projects[i]
		fmt.Printf("[%d] %s | %s | %d | %s | Teknologi: ", i+1, p.Name, p.Category, p.Difficulty, p.Date)
		
		for j := 0; j < p.TechCount; j++ {
			fmt.Print(p.Technologies[j], " ")
		}
		fmt.Println()
	}
}

func ubahProyek(projects *Project, n *int) {
	var idx int
	
	tampilkanProyek(*projects, *n)
	
	fmt.Print("Pilih nomor proyek yang ingin diubah: ")
	fmt.Scan(&idx)
	
	if idx < 1 || idx > *n {
		fmt.Println("Indeks tidak valid.")
	}
	
	for i := idx - 1; i < *n-1; i++ {
		projects[i] = projects[i+1]
	}
	*n--
	
	tambahProyek(projects, count)
}

func hapusProyek() {
	var idx int
	
	tampilkanProyek(*projects, *count)
	
	fmt.Print("Pilih nomor proyek yang ingin dihapus: ")
	fmt.Scan(&idx)
	
	if idx < 1 || idx > *count {
		fmt.Println("Indeks tidak valid.")
	}
	
	for i := idx - 1; i < *count-1; i++ {
		projects[i] = projects[i+1]
	}
	
	*count--
	fmt.Println("Proyek berhasil dihapus!")
} 

func sequentialSearch() {
	
}

func binarySearch() {
	
}

func selectionSort() {
	
}

func insertionSort() {
	
}

func statistikKategori() {
	
}