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
	for { 
	fmt.Println("----------------------------------")
	fmt.Println("|---Selamat Datang di Aplikasi---|")
	fmt.Println("----------------------------------")
	fmt.Println("1. Project User")
	fmt.Println("2. Cari Project User")
	fmt.Println("3. Urutkan Project User")
	fmt.Println("4. Keluar Dari Aplikasi")
	fmt.Print("pilih :")
	var pilihan int
	fmt.Scan(&pilihan)
		switch pilihan {'
			case 1:
				ProjectUser(&data, &nProject)
				var pilih int
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
			//case 2:
				//cariProject
			//case 3:
			//case 4:
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