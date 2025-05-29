package main

import "fmt"

const NMAX int = 100
type projects [NMAX]Project
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
		fmt.Println("--------------------------------------------")
		fmt.Println("|--------Selamat Datang di Aplikasi--------|")
		fmt.Println("--------------------------------------------")
		fmt.Println(" 1. Proyek User")
		fmt.Println(" 2. Cari Proyek User")
		fmt.Println(" 3. Urutkan Proyek User")
		fmt.Println(" 4. Keluar Dari Aplikasi")
		fmt.Println("--------------------------------------------")
		
		fmt.Print(" pilih: ")
		fmt.Scan(&pilihan)
		fmt.Println()
		
		switch pilihan {
			case 1:
				var pilih int
				
				pilih = ProjectUser(data, nProject)
				
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
				
				fmt.Println("--------------------------------------------")
				fmt.Println(" 1. Cari berdasarkan nama")
				fmt.Println(" 2. Cari berdasarkan kategori")
				fmt.Println("--------------------------------------------")
				
				fmt.Print(" Pilih: ")
				fmt.Scan(&pilih)
				fmt.Println()
				
				switch pilih {
					case 1 :
						seqSearch(data, nProject)
					case 2 :
						//binSearch(data, nProject)
				}
				
			case 3:
				var pilih int
				
				fmt.Println("--------------------------------------------")
				fmt.Println(" 1. Urutkan berdasarkan tingkat kesulitan")
				fmt.Println(" 2. Urutkan berdasarkan waktu")
				fmt.Println("--------------------------------------------")
				
				fmt.Print(" Pilih: ")
				fmt.Scan(&pilih)
				fmt.Println()
				
				switch pilih {
					case 1 :
						selSort(&data, nProject)
					case 2 :
						//insSort(data, nProject)
				}
		
			default :
					fmt.Println("Pilihan tidak valid!")
		}
		
		if pilihan == 4 {
			break
		}
	}
}

func ProjectUser(p projects, nProject int) int {
	var pilih int
	
	fmt.Println("--------------------------------------------")
	fmt.Println(" 1. Tambah Project")
	fmt.Println(" 2. Tampilkan Semua Project")
	fmt.Println(" 3. Ubah Project")
	fmt.Println(" 4. Hapus Project")
	fmt.Println("--------------------------------------------")
	
	fmt.Print(" Pilih: ")
	fmt.Scan(&pilih)
	fmt.Println()
	
	return pilih
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
		fmt.Println("|------------------------------------------|")
		fmt.Println(" [Project]", i+1)
		fmt.Println()
		fmt.Println(" Nama Project:" , p[i].nama)
		fmt.Println(" kategori: ", p[i].kategori)
		fmt.Println(" Tingkat kesulitan (1-10): ", p[i].tingkatKesulitan )
		fmt.Println(" Waktu (dd-mm-yyyy): ", p[i].waktu)
	} 
	
	fmt.Println("|------------------------------------------|")
	fmt.Println()
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
	fmt.Println()
	
	for i < n && idx == -1{
		if p[i].nama == x {
			idx = i
		}
		i++
	}
	
	if idx != -1 {
		fmt.Println("|------------------------------------------|")
		fmt.Println(" [Project]", idx + 1)
		fmt.Println()
		fmt.Println(" Nama Project: ", p[idx].nama)
		fmt.Println(" kategori: ", p[idx].kategori)
		fmt.Println(" Tingkat kesulitan (1-10):", p[idx].tingkatKesulitan )
		fmt.Println(" Waktu (dd-mm-yyyy): ", p[idx].waktu)
		fmt.Println("|------------------------------------------|")
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
	fmt.Println()
	
	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan)/ 2
		
		if x > p[tengah].kategori {
			kiri = tengah + 1
		} else if x < p[tengah].kategori {
			kanan = tengah - 1
		} else {
			idx = tengah
		}
	}
	
	if idx != -1 {
		fmt.Println("|------------------------------------------|")
		fmt.Println(" [Project]", idx + 1)
		fmt.Println()
		fmt.Println(" Nama Project: ", p[idx].nama)
		fmt.Println(" kategori: ", p[idx].kategori)
		fmt.Println(" Tingkat kesulitan (1-10):", p[idx].tingkatKesulitan )
		fmt.Println(" Waktu (dd-mm-yyyy): ", p[idx].waktu)
		fmt.Println("|------------------------------------------|")
		fmt.Println()
	} else {
		fmt.Println("Proyek tidak ditemukan")
	}
}

func selSort(p *projects, n int) {
	var pass, idx, i int
	var temp projects
	pass = 1
	
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		
		for i < n {
			if p[idx].tingkatKesulitan < p[i].tingkatKesulitan {
				idx = i
			}
			i++
		}
		
		temp = p[pass - 1]
		p[pass - 1] = p[idx]
		p[idx]= temp
		
		pass++
	}
}

func insSort() {
	
}