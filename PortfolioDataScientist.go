package main

import "fmt"

//NMAX adalah konstanta yang mendefinisikan ukuran maksimum array dari projects.
 //Ini berarti program dapat menyimpan hingga 100 project.
const NMAX int = 100

//projects adalah tipe data array yang dapat menampung NMAX (100) objek Project.
 //Ini adalah penyimpanan utama untuk menyimpan semua data project.
type projects [NMAX]Project

/*
  type Project adalah struktur (struct) yang mendefinisikan data untuk sebuah project.
   Setiap project akan memiliki atribut-atribut berikut:
   - nama: Nama project (string).
   - kategori: Kategori project (string).
   - tingkatKesulitan: Tingkat kesulitan project (integer),skala 1-10.
   - waktu: Waktu atau tanggal project (string), diharapkan dalam format seperti "dd-mm-yyyy".
*/
type Project struct {
	nama             string
	kategori         string
	tingkatKesulitan int
	waktu            string
}

	//main adalah fungsi utama program dan merupakan 
	//titik masuk eksekusi ketika program dijalankan.
func main() {
	menu()
}

/*
   menu adalah fungsi inti yang menampilkan menu utama aplikasi kepada pengguna.
   Ini juga mengelola alur program berdasarkan pilihan pengguna.
*/
func menu() {
	var data projects    // Variabel untuk menyimpan semua data project.
	var nProject int     // Variabel untuk melacak jumlah project yang sudah ada.
	var pilihan int      // Variabel untuk menyimpan pilihan menu pengguna.

	for { // Loop tak terbatas agar menu terus ditampilkan hingga pengguna memilih untuk keluar.
		fmt.Println("--------------------------------------------")
		fmt.Println("|--------Selamat Datang di Aplikasi--------|")
		fmt.Println("--------------------------------------------")
		fmt.Println("| 1. Project User                          |")
		fmt.Println("| 2. Cari Project User                     |")
		fmt.Println("| 3. Urutkan Project User                  |")
		fmt.Println("| 4. Keluar Dari Aplikasi                  |")
		fmt.Println("--------------------------------------------")

		fmt.Print(" pilih: ")
		fmt.Scan(&pilihan) // Membaca pilihan menu dari pengguna.
		fmt.Println()

		switch pilihan {
		case 1: // Opsi untuk manajemen project (Tambah, Tampilkan, Ubah, Hapus).
			var pilih int

			// Memanggil ProjectUser untuk menampilkan sub-menu manajemen project
			// dan mendapatkan pilihan pengguna dari sub-menu tersebut.
			pilih = ProjectUser(data, nProject)

			switch pilih {
			case 1:
				tambahProject(&data, &nProject) // Memanggil fungsi untuk menambah project.
			case 2:
				tampilkanProject(data, nProject) // Memanggil fungsi untuk menampilkan semua project.
			case 3:
				ubahProject(&data, nProject) // Memanggil fungsi untuk mengubah detail project.
			case 4:
				hapusProject(&data, &nProject) // Memanggil fungsi untuk menghapus project.
			}

		case 2: // Opsi untuk mencari project.
			var pilih int

			fmt.Println("-----------------------------------------------")
			fmt.Println("| 1. -> Cari berdasarkan nama                 |")
			fmt.Println("| 2. -> Cari berdasarkan tingkat kesulitan    |")
			fmt.Println("| 3. -> Kembali                               |")
			fmt.Println("-----------------------------------------------")

			fmt.Print(" Pilih: ")
			fmt.Scan(&pilih) // Membaca pilihan pencarian dari pengguna.
			fmt.Println()

			switch pilih {
			case 1:
				seqSearch(data, nProject) // Mencari project berdasarkan nama (Sequential Search).
			case 2:
				binSearch(data, nProject) // Mencari project berdasarkan tingkat kesulitan (Binary Search).
				// Catatan: binSearch mengasumsikan data sudah terurut berdasarkan tingkat kesulitan.
			}

		case 3: // Opsi untuk mengurutkan project.
			var pilih int

			fmt.Println("-----------------------------------------------")
			fmt.Println("| 1. -> Urutkan berdasarkan tingkat kesulitan |")
			fmt.Println("| 2. -> Urutkan berdasarkan waktu             |")
			fmt.Println("| 3. -> Kembali                               |")
			fmt.Println("-----------------------------------------------")

			fmt.Print(" Pilih: ")
			fmt.Scan(&pilih) // Membaca pilihan pengurutan dari pengguna.
			fmt.Println()

			switch pilih {
			case 1:
				selSort(&data, nProject)         // Mengurutkan project berdasarkan tingkat kesulitan (Selection Sort).
				tampilkanProject(data, nProject) // Menampilkan project setelah diurutkan.
			case 2:
				insSort(&data, nProject)         // Mengurutkan project berdasarkan waktu (Insertion Sort).
				tampilkanProject(data, nProject) // Menampilkan project setelah diurutkan.
			}

		default: // Opsi jika pilihan tidak ada dalam menu.
			fmt.Println("~ SELAMAT TINGGAL ~")
		}

		if pilihan == 4 { // Keluar dari loop jika pengguna memilih opsi '4'.
			break
		}
	}
}

/*
   ProjectUser adalah fungsi yang menampilkan sub-menu untuk operasi manajemen project.
   function ini dibuat agar pengguna dapat memilih antara menambah, menampilkan, mengubah, atau menghapus project.
   Mengembalikan pilihan pengguna dari sub-menu ini.
*/
func ProjectUser(p projects, nProject int) int {
	var pilih int

	fmt.Println("--------------------------------------------")
	fmt.Println("| 1. -> Tambah Project                     |")
	fmt.Println("| 2. -> Tampilkan Semua Project            |")
	fmt.Println("| 3. -> Ubah Project                       |")
	fmt.Println("| 4. -> Hapus Project                      |")
	fmt.Println("| 5. -> Kembali                            |")
	fmt.Println("--------------------------------------------")

	fmt.Print(" Pilih: ")
	fmt.Scan(&pilih) // Membaca pilihan sub-menu dari pengguna.
	fmt.Println()

	return pilih // Mengembalikan pilihan yang diambil pengguna.
}

/*
   tambahProject, function ini berfungsi untuk menambahkan project baru ke dalam daftar.
   Parameter:
   - p: Pointer ke array projects, memungkinkan modifikasi data asli.
   - nProject: Pointer ke jumlah project saat ini, memungkinkan pembaruan jumlah project.
*/
func tambahProject(p *projects, nProject *int) {

	fmt.Print("Nama Project: ")
	fmt.Scan(&p[*nProject].nama) // Meminta dan menyimpan nama project.
	fmt.Print("Tingkat Kesulitan (1-10): ")
	fmt.Scan(&p[*nProject].tingkatKesulitan) // Meminta dan menyimpan tingkat kesulitan.

	fmt.Print("Kategori: ")
	fmt.Scan(&p[*nProject].kategori) // Meminta dan menyimpan kategori.
	fmt.Print("Waktu (dd-mm-yyyy): ")
	fmt.Scan(&p[*nProject].waktu) // Meminta dan menyimpan waktu.

	*nProject = *nProject + 1 // Menambah jumlah project setelah penambahan.

	fmt.Println("---------------------------------")
	fmt.Println("|Project berhasil ditambahkan!!!|")
	fmt.Println("---------------------------------")
	fmt.Println()
}

/*
   tampilkanProject, function ini berfungsi untuk menampilkan detail semua project yang ada.
   Parameter:
   - p: Array projects yang berisi data project.
   - nProject: Jumlah project yang ada saat ini.
*/
func tampilkanProject(p projects, nProject int) {
	if nProject == 0 { // Memeriksa apakah tidak ada project yang tersedia.
		fmt.Println("Maaf Data Belum Tersedia")
	}

	for i := 0; i < nProject; i++ { // Melakukan perulangan untuk setiap project.
		fmt.Println("|------------------------------------------|")
		fmt.Println("|[Project]", i+1, "|", "                            |")
		fmt.Println("|------------------------------------------|")
		fmt.Println(" Nama Project: ", p[i].nama)               // Menampilkan nama project.
		fmt.Println(" kategori: ", p[i].kategori)               // Menampilkan kategori project.
		fmt.Println(" Tingkat kesulitan (1-10): ", p[i].tingkatKesulitan) // Menampilkan tingkat kesulitan.
		fmt.Println(" Waktu (dd-mm-yyyy): ", p[i].waktu)         // Menampilkan waktu project.
	}
	fmt.Println("|------------------------------------------|")
	fmt.Println()
}

/*
   ubahProject, function ini berfungsi untuk memodifikasi detail project yang sudah ada.
   Parameter:
   - p: Pointer ke array projects, memungkinkan modifikasi data asli.
   - nProject: Jumlah project yang ada saat ini.
*/
func ubahProject(p *projects, nProject int) {
	tampilkanProject(*p, nProject) // Menampilkan daftar project agar pengguna bisa memilih.
	fmt.Print("Pilih nomor yang ingin diubah :")

	var ubah int
	fmt.Scan(&ubah) // Membaca nomor project yang ingin diubah.
	ubah--          // Mengurangi 1 karena indeks array dimulai dari 0.

	if ubah >= 0 && ubah < nProject { // Memeriksa apakah nomor yang dimasukkan valid.
		fmt.Print("Nama Project baru : ")
		fmt.Scan(&p[ubah].nama) // Meminta dan menyimpan nama project baru.
		fmt.Print("Tingkat kesulitan (1-10): ")
		fmt.Scan(&p[ubah].tingkatKesulitan) // Meminta dan menyimpan tingkat kesulitan baru.

		fmt.Print("kategori: ")
		fmt.Scan(&p[ubah].kategori) // Meminta dan menyimpan kategori baru.
		fmt.Print("Waktu (dd-mm-yyyy): ")
		fmt.Scan(&p[ubah].waktu) // Meminta dan menyimpan waktu baru.
		fmt.Println()

		fmt.Println("Project berhasil diubah!!")
	} else {
		fmt.Println()
		fmt.Println("Nomor tidak sesuai!!") // Pesan kesalahan jika nomor tidak valid.
		fmt.Println()
	}
}

/*
   hapusProject, function ini berfungsi untuk menghapus project dari daftar.
   Parameter:
   - p: Pointer ke array projects, memungkinkan modifikasi data asli.
   - nProject: Pointer ke jumlah project saat ini, memungkinkan pembaruan jumlah project.
*/
func hapusProject(p *projects, nProject *int) {
	tampilkanProject(*p, *nProject) // Menampilkan daftar project agar pengguna bisa memilih.
	fmt.Print("Pilih nomor project yang ingin dihapus: ")
	var idx int
	fmt.Scan(&idx) // Membaca nomor project yang ingin dihapus.
	idx--          // Mengurangi 1 karena indeks array dimulai dari 0.

	if idx >= 0 && idx < *nProject { // Memeriksa apakah indeks yang dimasukkan valid.
		// Menggeser elemen-elemen setelah project yang dihapus ke kiri
		for i := idx; i < *nProject-1; i++ {
			p[i] = p[i+1]
		}
		*nProject-- // Mengurangi jumlah project setelah penghapusan.
		fmt.Println("Project berhasil dihapus.")
	} else {
		fmt.Println("Indeks tidak valid.") // Pesan kesalahan jika indeks tidak valid.
	}
}

/*
   seqSearch adalah function untuk melakukan pencarian sekuensial (linear search)
   berdasarkan nama project.
   Parameter:
   - p: Array projects yang berisi data project.
   - n: Jumlah project yang ada saat ini.
*/
func seqSearch(p projects, n int) {
	var i, idx int // i untuk perulangan, idx untuk menyimpan indeks ditemukan.
	var x string   // x untuk menyimpan nama project yang dicari.

	i = 0
	idx = -1 // Inisialisasi idx dengan -1 (tidak ditemukan).

	fmt.Print("Masukan nama project yang ingin dicari: ")
	fmt.Scan(&x) // Meminta nama project yang akan dicari.
	fmt.Println()

	// Melakukan perulangan selama belum mencapai akhir array dan belum ditemukan.
	for i < n && idx == -1 {
		if p[i].nama == x { // Jika nama project cocok.
			idx = i // Simpan indeks project yang ditemukan.
		}
		i++ // Lanjut ke project berikutnya.
	}

	if idx != -1 { // Jika project ditemukan (idx bukan -1).
		fmt.Println("|------------------------------------------|")
		// Menampilkan detail project yang ditemukan.
		fmt.Println("|[Project]", idx+1, "|", "                                    |")
		fmt.Println("|------------------------------------------|")
		// Gunakan 'idx' untuk menampilkan data project yang ditemukan, bukan 'i'
		// Karena 'i' akan menjadi n jika tidak ditemukan atau 1 lebih dari 'idx' jika ditemukan
		fmt.Println(" Nama Project: ", p[idx].nama)
		fmt.Println(" kategori: ", p[idx].kategori)
		fmt.Println(" Tingkat kesulitan (1-10): ", p[idx].tingkatKesulitan)
		fmt.Println(" Waktu (dd-mm-yyyy): ", p[idx].waktu)
		fmt.Println("|------------------------------------------|")
		fmt.Println()
	} else {
		fmt.Println("Project tidak ditemukan") // Pesan jika project tidak ditemukan.
	}
}

/*
   binSearch adalah function untuk melakukan pencarian biner (binary search)
   berdasarkan tingkat kesulitan project.
   Penting: Fungsi ini mengasumsikan bahwa data project sudah terurut
   berdasarkan tingkat kesulitan secara menaik (ascending) sebelum dipanggil.
   Parameter:
   - p: Array projects yang berisi data project.
   - n: Jumlah project yang ada saat ini.
*/
func binSearch(p projects, n int) {
	var kanan, kiri, tengah, x int // Variabel untuk batas pencarian dan nilai yang dicari.
	var idx int                    // Variabel untuk menyimpan indeks ditemukan.

	idx = -1   // Inisialisasi idx dengan -1 (tidak ditemukan).
	kiri = 0   // Batas kiri array.
	kanan = n - 1 // Batas kanan array.

	fmt.Print("Masukan kesulitan project yang ingin dicari: ")
	fmt.Scan(&x) // Meminta tingkat kesulitan yang akan dicari.
	fmt.Println()

	// Melakukan perulangan selama batas kiri kurang dari atau sama dengan batas kanan
	// dan project belum ditemukan.
	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2 // Menghitung indeks tengah.

		if x > p[tengah].tingkatKesulitan { // Jika nilai yang dicari lebih besar dari nilai di tengah.
			kiri = tengah + 1 // Pindahkan batas kiri ke kanan dari tengah.
		} else if x < p[tengah].tingkatKesulitan { // Jika nilai yang dicari lebih kecil dari nilai di tengah.
			kanan = tengah - 1 // Pindahkan batas kanan ke kiri dari tengah.
		} else { // Jika nilai yang dicari sama dengan nilai di tengah.
			idx = tengah // project ditemukan, simpan indeksnya.
		}
	}

	if idx != -1 { // Jika project ditemukan.
		fmt.Println("|------------------------------------------|")
		fmt.Println(" [Project]", idx+1)
		fmt.Println()
		// Menampilkan detail project yang ditemukan.
		fmt.Println(" Nama Project: ", p[idx].nama)
		fmt.Println(" kategori: ", p[idx].kategori)
		fmt.Println(" Tingkat kesulitan (1-10):", p[idx].tingkatKesulitan)
		fmt.Println(" Waktu (dd-mm-yyyy): ", p[idx].waktu)
		fmt.Println("|------------------------------------------|")
		fmt.Println()
	} else {
		fmt.Println("Project tidak ditemukan") // Pesan jika project tidak ditemukan.
	}
}

/*
   selSort adalah fungsi yang mengimplementasikan algoritma Selection Sort.
   Ini mengurutkan array project berdasarkan tingkat kesulitan secara menaik (ascending).
   Parameter:
   - p: Pointer ke array projects, memungkinkan modifikasi data asli (pengurutan).
   - n: Jumlah project yang ada saat ini.
*/
func selSort(p *projects, n int) {
	var pass, idx, i int // Variabel untuk pass, indeks minimum, dan perulangan.
	var temp Project     // Variabel sementara untuk pertukaran elemen.

	pass = 1 // Pass pertama dimulai dari 1.

	for pass <= n-1 { // Melakukan perulangan sebanyak n-1 pass.
		idx = pass - 1 // Awalnya, indeks elemen minimum adalah elemen pertama dari bagian yang belum terurut.
		i = pass       // Mulai pencarian elemen minimum dari elemen berikutnya.

		for i < n { // Mencari elemen dengan tingkat kesulitan terkecil di sisa array.
			if p[idx].tingkatKesulitan > p[i].tingkatKesulitan {
				idx = i // Jika ditemukan yang lebih kecil, perbarui indeks minimum.
			}
			i++ // Lanjut ke elemen berikutnya.
		}

		// Melakukan pertukaran (swap) antara elemen di posisi 'pass-1'
		// dengan elemen yang memiliki tingkat kesulitan terkecil (diindeks 'idx').
		temp = p[pass-1]
		p[pass-1] = p[idx]
		p[idx] = temp

		pass++ // Lanjut ke pass berikutnya.
	}
}

/*
   insSort adalah function yang mengimplementasikan algoritma Insertion Sort.
   Ini mengurutkan array project berdasarkan waktu secara menaik (ascending).
   Parameter:
   - p: Pointer ke array projects, memungkinkan modifikasi data asli (pengurutan).
   - n: Jumlah project yang ada saat ini.
*/
func insSort(p *projects, n int) {
	var temp Project
	for i := 1; i < n; i++ { // Memulai dari elemen kedua (indeks 1).
		j := i - 1     // Indeks elemen sebelumnya.
		temp = p[i]   // Menyimpan elemen saat ini untuk disisipkan.

		// Geser elemen-elemen yang lebih besar dari temp ke kanan
		// untuk membuat ruang bagi temp pada posisi yang benar.ss
		for j >= 0 && p[j].waktu > temp.waktu {
			p[j+1] = p[j] // Geser elemen ke kanan.
			j--           // Pindah ke elemen sebelumnya.
		}
		p[j+1] = temp // Sisipkan elemen temp pada posisi yang benar.
	}
}