package main

import "fmt"

type projects [100]Project
type Project struct {
	nama string
	kategori string
	tingkatKesulitan int
	waktu string
}

func main() {
	menu()
}

func menu(){
	var data projects
	var nProject int
	var pilihan int
	
	for { 
		fmt.Println("----------------------------------")
		fmt.Println("|---Selamat Datang di Aplikasi---|")
		fmt.Println("----------------------------------")
		fmt.Println("1. Proyek User")
		fmt.Println("2. Cari Proyek User")
		fmt.Println("3. Urutkan Proyek User")
		fmt.Println("4. Keluar Dari Aplikasi")
		fmt.Println("----------------------------------")
		
		fmt.Print("pilih :")
		fmt.Scan(&pilihan)
		
		switch pilihan {'
			case 1:
				var pilih int
				
				ProjectUser(&data, &nProject)
				fmt.Scan(&pilih)
				
				switch pilih {
					case 1:
						tambahProject(&data, &nProject)
					case 2:
						tampilkanProject(data, nProject)
					case 3:
						ubahProject(&data, nProject)
					//case 4:
						//hapusProject()
				}
				
			case 2:
				var pilih int
				
				fmt.Println("1. Cari berdasarkan nama")
				fmt.Println("2. Cari berdasarkan kategori")
				fmt.Scan(&pilih)
				
				switch pilih {
					case 1 :
						seqSearch(data, nProject)
					case 2 :
						//binSearch(data, nProject)
				}
				
			//case 3:
				//selSort(data, nProject)
				//insSort(data, nProject)
				
			case 4:
				fmt.Println("Terima kasih")
				break
		
			default :
					fmt.Println("Pilihan tidak valid!")
		}
	}
}

func ProjectUser(p *projects, nProject *int) {
	fmt.Println("1. Tambah Project")
	fmt.Println("2. Tampilkan Semua Project")
	fmt.Println("3. Ubah Project")
	fmt.Println("4. Hapus Project")
	fmt.Print("Pilih :")
}

func tambahProject (p *projects, nProject *int) {
		
	fmt.Print("Nama Project: ")
	fmt.Scan(&p[*nProject].nama)
	fmt.Print("Tingkat kesulitan (1-10): ")
	fmt.Scan(&p[*nProject].tingkatKesulitan)
	
	fmt.Print("kategori: ")
	fmt.Scan(&p[*nProject].kategori)
	fmt.Print("Waktu (dd-mm-yyyy): ")
	fmt.Scan(&p[*nProject].waktu)
	
	*nProject = *nProject + 1
	
	fmt.Println("Project berhasil ditambahkan!!!")
	fmt.Println()
}

func tampilkanProject (p projects, nProject int) {
	if nProject == 0 {
		fmt.Println("Maaf Data Belum Tersedia")
	}
	
	
	for i := 0; i < nProject; i++ {
		fmt.Println("[Project]", i+1)
		fmt.Println()
		fmt.Println("Nama Project :", p[i].nama)
		fmt.Println("kategori", p[i].kategori)
		fmt.Println("Tingkat kesulitan (1-10):", p[i].tingkatKesulitan )
		fmt.Println("Waktu (dd-mm-yyyy)", p[i].waktu)
		fmt.Println()
	} 
}

func ubahProject(p *projects, nProject int) {
	tampilkanProject(*p, nProject)
	fmt.Print("Pilih nomor yang ingin diubah :")
		
	var ubah int
		fmt.Scan(&ubah)
		ubah--
		
		if ubah >= 0 && ubah < nProject {
			fmt.Print("Nama Project baru : ")
			fmt.Scan(&p[ubah].nama)
			fmt.Print("Tingkat kesulitan (1-10): ")
			fmt.Scan(&p[ubah].tingkatKesulitan)
			
			fmt.Print("kategori: ")
			fmt.Scan(&p[ubah].kategori)
			fmt.Print("Waktu (dd-mm-yyyy): ")
			fmt.Scan(&p[ubah].waktu)
			fmt.Println()
			
			fmt.Println("Project berhasil diubah!!")
		} else {
			fmt.Println()
			fmt.Println("Nomor tidak sesuai!!")
			fmt.Println()
		}
}

func seqSearch(p projects, n int) {
	var i, idx int
	var x string
	
	i = 0
	idx = -1
	
	fmt.Print("Masukan nama proyek yang ingin dicari: ")
	fmt.Scan(&x)
	
	for i < n && idx == -1{
		if p[i].nama == x {
			idx = i
		}
		i++
	}
	
	if idx != -1 {
		fmt.Prinf("[%d] Nama: &s | Kategori: &s | Tingkat kesulitan: $d | Waktu: &s", idx, p[idx].nama, p[idx].kategori, p[idx].tingkatKesulitan, p[idx].waktu)
		fmt.Println()
	} else {
		fmt.Println("Proyek tidak ditemukan")
	}
}

func binSearch(p projects, n int) {
	var kanan, kiri, tengah int
	var idx int
	var x string
	
	idx = -1
	kiri = 0
	kanan = n - 1
	
	fmt.Print("Masukan kategori proyek yang ingin dicari: ")
	fmt.Scan(&x)
	
	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan)/ 2
		
		if x > tengah {
			kiri = tengah + 1
		} else if x < tengah {
			kanan = tengah - 1
		} else {
			idx = tengah
		}
	}
	
	if idx != -1 {
		fmt.Prinf("[%d] Nama: &s | Kategori: &s | Tingkat kesulitan: $d | Waktu: &s", idx, p[idx].nama, p[idx].kategori, p[idx].tingkatKesulitan, p[idx].waktu)
		fmt.Println()
	} else {
		fmt.Println("Proyek tidak ditemukan")
	}
}

func selSort() {
	
}

func insSort() {
	
}