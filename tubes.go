package main

import "fmt"

const NMAX int = 1000
const NMAXX int = 10000

type RelasiP struct {
	namaPF, namaPW, fasilitas, wahana [NMAXX]string
	nf, nw                            int
}

type pariwisata struct {
	nama  [NMAX]string
	n     int
	jarak [NMAX]float64
	harga [NMAX]float64
}

func TambahData(P *pariwisata, RP *RelasiP) {
	/* Fitur mode admin untuk menambah data pada array*/
	var w, fas, namaP string
	var m bool
	fmt.Println("Nama Tempat?")
	fmt.Println("(Penggunaan spasi diganti dengan _)")
	fmt.Scan(&namaP)
	P.nama[P.n] = namaP
	fmt.Println("Jarak? (m)")
	fmt.Scan(&P.jarak[P.n])
	fmt.Println("Harga? (Rp)")
	fmt.Scan(&P.harga[P.n])
	fmt.Println("Input fasilitas-fasilitas yang tersedia")
	fmt.Println("(Spasi diganti dengan _ dan akhiri input fasilitas dengan ---)")
	for !m {
		fmt.Scan(&fas)
		m = fas == "---"
		if !m {
			RP.namaPF[RP.nf] = namaP
			RP.fasilitas[RP.nf] = fas
			RP.nf++
		}
	}
	fmt.Println("Input wahana-wahana yang tersedia")
	fmt.Println("(Spasi diganti dengan _ dan akhiri input wahana dengan ---)")
	m = false
	for !m {
		fmt.Scan(&w)
		m = w == "---"
		if !m {
			RP.namaPW[RP.nw] = namaP
			RP.wahana[RP.nw] = w
			RP.nw++
		}
	}
	P.n++
	fmt.Print("\033[H\033[2J")
	fmt.Println("Input Pariwisata Berhasil!")
}

func AdminEdit(P *pariwisata, RP *RelasiP) {
	/* Fitur mode admin untuk meng-edit data pada array P dan RP */
	var idx, opsi int
	var namaP, temp string
	var fasw string
	for i := 0; i < P.n; i++ {
		fmt.Print(i, ".", P.nama[i])
		fmt.Println("")
	}
	fmt.Println("Pilih data pariwisata yang akan diubah sesuai indeksnya")
	fmt.Println("Input ")
	fmt.Scan(&idx)
	namaP = P.nama[idx]
	temp = namaP
	if idx >= P.n {
		fmt.Println("Indeks Tidak Ditemukan!")
	} else {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Pilih data dari ", namaP, " yang akan diubah")
		fmt.Println("1. Data Pariwisata")
		fmt.Println("2. Data Fasilitas")
		fmt.Println("3. Data Wahana")
		fmt.Println("4. Batal Edit")
		fmt.Print("Input?")
		fmt.Scan(&opsi)
		if opsi == 4 {
			fmt.Print("\033[H\033[2J")
		} else if opsi == 3 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Daftar wahana dari ", namaP)
			for i := 0; i < RP.nw; i++ {
				if RP.namaPW[i] == namaP {
					fmt.Println(i, ".", RP.wahana[i])
				}
			}
			fmt.Print("Input wahana ", namaP, " yang akan diganti sesuai indeksnya: ")
			fmt.Scan(&idx)
			fmt.Println("")
			fmt.Println("Input wahana: (isi --- untuk membatalkan)")
			fmt.Scan(&fasw)
			if fasw != "---" && RP.namaPW[idx] == namaP {
				RP.wahana[idx] = fasw
				fmt.Println("Edit wahana berhasil!")
			} else {
				fmt.Println("Edit wahana dibatalkan / indeks Salah")
			}
		} else if opsi == 2 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Daftar fasilitas dari ", namaP)
			for i := 0; i < RP.nf; i++ {
				if RP.namaPF[i] == namaP {
					fmt.Println(i, ".", RP.fasilitas[i])
				}
			}
			fmt.Print("Input fasilitas ", namaP, " yang akan diganti sesuai indeksnya: ")
			fmt.Scan(&idx)
			fmt.Println("")
			fmt.Println("Input fasilitas: (isi --- untuk membatalkan)")
			fmt.Scan(&fasw)
			if fasw != "---" && RP.namaPF[idx] == namaP {
				RP.fasilitas[idx] = fasw
				fmt.Println("Edit fasilitas berhasil!")
			} else {
				fmt.Println("Edit fasilitas dibatalkan / indeks salah")
			}

		} else if opsi == 1 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Nama Tempat?")
			fmt.Println("(Penggunaan spasi diganti dengan _)")
			fmt.Scan(&namaP)
			P.nama[idx] = namaP
			fmt.Println("Jarak? (m)")
			fmt.Scan(&P.jarak[idx])
			fmt.Println("Harga? (Rp)")
			fmt.Scan(&P.harga[idx])
			for i := 0; i < RP.nf; i++ {
				if RP.namaPF[i] == temp {
					RP.namaPF[i] = namaP
				}
			}
			for i := 0; i < RP.nw; i++ {
				if RP.namaPW[i] == temp {
					RP.namaPW[i] = namaP
				}
			}
			fmt.Print("\033[H\033[2J")
			fmt.Println("Edit data pariwisata berhasil!")

		}
	}
}

func AdminDel(P *pariwisata, RP *RelasiP) {
	/*Penghapusan data pariwisata*/
	var idx, opsi, j int
	for i := 0; i < P.n; i++ {
		fmt.Print(i, ".", P.nama[i])
		fmt.Println("")
	}
	fmt.Println("Pilih data pariwisata yang akan dihapus sesuai indeksnya")
	fmt.Println("Input ")
	fmt.Scan(&idx)
	if idx >= P.n {
		fmt.Println("Indeks Tidak Ditemukan!")
	} else {
		fmt.Println("Anda Yakin? (1=Ya ,2=Tidak)")
		fmt.Scan(&opsi)
		if opsi == 1 {
			j = idx
			for j < P.n {
				P.nama[j] = P.nama[j+1]
				P.harga[j] = P.harga[j+1]
				P.jarak[j] = P.jarak[j+1]
				j++
			}
			if idx < P.n {
				P.n--
			}
			fmt.Println("Penghapusan data berhasil!")
		} else {
			fmt.Println("Penghapusan data dibatalkan")
		}
	}
}

func Admin(exit bool, P *pariwisata, RP *RelasiP) {
	/* Menu admin untuk mengakses dan memanipulasi data*/
	var opsi int
	for exit == false {
		fmt.Println("______________________ADMIN__________________________")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data ")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Main Menu               ")
		fmt.Println("=====================================================")
		fmt.Print("Input?")
		fmt.Scan(&opsi)
		fmt.Print("\033[H\033[2J")
		if opsi == 4 {
			exit = true
			MainMenu(false, P, RP)
		} else if opsi == 3 {
			AdminDel(P, RP)
		} else if opsi == 2 {
			AdminEdit(P, RP)
		} else if opsi == 1 {
			TambahData(P, RP)
		}
	}

}
func DaftarP(exit bool, P pariwisata, RP RelasiP) {
	/*Prosedur untuk menampilkan menu daftar pariwisata secara terurut berdasarkan data atau tidak*/
	var P2 pariwisata
	var opsi int
	for exit == false {
		fmt.Println("__________________#Daftar Pariwisata#__________________")
		fmt.Println("1. Urut Sesuai Jarak")
		fmt.Println("2. Urut Sesuai Harga")
		fmt.Println("3. Urut Sesuai Nama ")
		fmt.Println("4. Daftar Tidak Terurut")
		fmt.Println("5. Main Menu")
		fmt.Println("=======================================================")
		fmt.Print("Input?")
		fmt.Scan(&opsi)
		fmt.Print("\033[H\033[2J")
		P2 = P
		if opsi == 1 {
			DaftarPUrutJarak(&P2)
			TampilArr(P2, RP)
		} else if opsi == 2 {
			DaftarUrutHarga(&P2)
			TampilArr(P2, RP)
		} else if opsi == 4 {
			TampilArr(P2, RP)
		} else if opsi == 5 {
			exit = true
		} else if opsi == 3 {
			DaftarPUrutNama(&P2)
			TampilArr(P2, RP)
		}
	}
}

func DaftarPUrutNama(P *pariwisata) {
	var pass, idx int
	var temp pariwisata
	for pass = 1; pass < P.n; pass++ {
		idx = pass - 1
		for i := pass; i < P.n; i++ {
			if P.nama[idx] > P.nama[i] {
				idx = i
			}
		}
		temp.nama[idx] = P.nama[idx]
		temp.harga[idx] = P.harga[idx]
		temp.jarak[idx] = P.jarak[idx]
		P.nama[idx] = P.nama[pass-1]
		P.harga[idx] = P.harga[pass-1]
		P.jarak[idx] = P.jarak[pass-1]
		P.nama[pass-1] = temp.nama[idx]
		P.harga[pass-1] = temp.harga[idx]
		P.jarak[pass-1] = temp.jarak[idx]
	}
}

func DaftarPUrutJarak(P *pariwisata) {
	var pass, idx int
	var temp pariwisata
	for pass = 1; pass < P.n; pass++ {
		idx = pass - 1
		for i := pass; i < P.n; i++ {
			if P.jarak[idx] > P.jarak[i] {
				idx = i
			}
		}
		temp.nama[idx] = P.nama[idx]
		temp.harga[idx] = P.harga[idx]
		temp.jarak[idx] = P.jarak[idx]
		P.nama[idx] = P.nama[pass-1]
		P.harga[idx] = P.harga[pass-1]
		P.jarak[idx] = P.jarak[pass-1]
		P.nama[pass-1] = temp.nama[idx]
		P.harga[pass-1] = temp.harga[idx]
		P.jarak[pass-1] = temp.jarak[idx]
	}
}

func DaftarUrutHarga(P *pariwisata) {
	var i, j int
	for i = 1; i < P.n; i++ {
		key1 := P.harga[i]
		key2 := P.nama[i]
		key3 := P.jarak[i]
		j = i - 1
		for j >= 0 && key1 < P.harga[j] {
			P.nama[j+1] = P.nama[j]
			P.harga[j+1] = P.harga[j]
			P.jarak[j+1] = P.jarak[j]
			j--
		}
		P.harga[j+1] = key1
		P.nama[j+1] = key2
		P.jarak[j+1] = key3
	}
}

func TampilArr(P pariwisata, RP RelasiP) {
	/*Prosedur untuk menampilkan array pariwisata dan relasinya*/
	var nama string
	for i := 0; i < P.n; i++ {
		nama = P.nama[i]
		idxnf := 0
		idxnw := 0
		for k := 0; k < RP.nf; k++ {
			if RP.namaPF[k] == nama {
				idxnf++
			}
		}
		for j := 0; j < RP.nw; j++ {
			if RP.namaPW[j] == nama {
				idxnw++
			}
		}
		fmt.Println("    ")
		fmt.Println("========================", i+1, "===========================")
		fmt.Println("Nama:", P.nama[i])
		fmt.Println("Jarak:", P.jarak[i], "m")
		fmt.Println("Harga: Rp.", P.harga[i])
		fmt.Print("Fasilitas: ")
		for i := 0; i < RP.nf; i++ {
			if RP.namaPF[i] == nama && idxnf == 1 {
				fmt.Print(RP.fasilitas[i])
			} else if RP.namaPF[i] == nama {
				fmt.Print(RP.fasilitas[i], ",")
				idxnf--
			}
		}
		fmt.Print(".")
		fmt.Println("")
		fmt.Print("Wahana: ")
		for i := 0; i < RP.nw; i++ {
			if RP.namaPW[i] == nama && idxnw == 1 {
				fmt.Print(RP.wahana[i])
			} else if RP.namaPW[i] == nama {
				fmt.Print(RP.wahana[i], ",")
				idxnw--
			}
		}
		fmt.Print(".")
		fmt.Println("")
		fmt.Println("======================================================")
	}

	fmt.Println("    ")
	fmt.Println(" ")
}

func PencarianNFW(P pariwisata, RP RelasiP, s string, opsi int) {
	/*Prosedur pencarian data pada array sesuai kata kunci dan opsi lalu menampilkan semua data terkait s */
	var P2 pariwisata
	var names [NMAX]string
	var nNama int
	var s2, opt string
	var nilai float64
	var left, right, mid int
	if opsi == 1 {
		/*names[0] = s
		nNama++*/
		/*Pencarian nama menggunakan binary search*/
		DaftarPUrutNama(&P)
		right = P.n
		mid = (left + right) / 2
		for left <= right && P.nama[mid] != s {
			if s < P.nama[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			mid = (left + right) / 2
		}
		if P.nama[mid] == s {
			names[0] = P.nama[mid]
			nNama++
		}
	} else if opsi == 2 {
		for i := 0; i < RP.nf; i++ {
			if RP.fasilitas[i] == s {
				names[nNama] = RP.namaPF[i]
				nNama++
			}
		}
	} else if opsi == 3 {
		for i := 0; i < RP.nw; i++ {
			if RP.wahana[i] == s {
				names[nNama] = RP.namaPW[i]
				nNama++
			}
		}
	} else if opsi == 4 {
		fmt.Print("Masukan Harga (Rp):")
		fmt.Scan(&nilai)
		fmt.Println("Pilih Pencarian Sesuai Harga")
		fmt.Println("A = Harga lebih kecil dari Rp.", nilai)
		fmt.Println("B = Harga sama dengan Rp.", nilai)
		fmt.Println("C = Harga lebih besar dari Rp.", nilai)
		fmt.Println("Input?")
		fmt.Scan(&opt)
		if opt == "A" {
			for i := 0; i < P.n; i++ {
				if P.harga[i] < nilai {
					names[nNama] = P.nama[i]
					nNama++
				}
			}

		} else if opt == "B" {
			for i := 0; i < P.n; i++ {
				if P.harga[i] == nilai {
					names[nNama] = P.nama[i]
					nNama++
				}
			}

		} else if opt == "C" {
			for i := 0; i < P.n; i++ {
				if P.harga[i] > nilai {
					names[nNama] = P.nama[i]
					nNama++
				}
			}
		}
	} else if opsi == 5 {
		fmt.Print("Masukan Jarak (m):")
		fmt.Scan(&nilai)
		fmt.Println("Pilih Pencarian Sesuai Jarak")
		fmt.Println("A = Jarak lebih kecil dari", nilai, "m")
		fmt.Println("B = Jarak sama dengan ", nilai, "m")
		fmt.Println("C = Jarak lebih besar dari", nilai, "m")
		fmt.Println("Input?")
		fmt.Scan(&opt)
		if opt == "A" {
			for i := 0; i < P.n; i++ {
				if P.jarak[i] < nilai {
					names[nNama] = P.nama[i]
					nNama++
				}
			}

		} else if opt == "B" {
			for i := 0; i < P.n; i++ {
				if P.jarak[i] == nilai {
					names[nNama] = P.nama[i]
					nNama++
				}
			}

		} else if opt == "C" {
			for i := 0; i < P.n; i++ {
				if P.jarak[i] > nilai {
					names[nNama] = P.nama[i]
					nNama++
				}
			}
		}
	}
	for i := 0; i < nNama; i++ {
		s2 = names[i]
		for i := 0; i < P.n; i++ {
			if s2 == P.nama[i] {
				P2.nama[P2.n] = P.nama[i]
				P2.harga[P2.n] = P.harga[i]
				P2.jarak[P2.n] = P.jarak[i]
				P2.n++
			}
		}
	}
	TampilArr(P2, RP)
	if P2.n == 0 {
		fmt.Println("Pariwisata tidak ditemukan")
	}
}

func PencarianKey(exit bool, P pariwisata, RP RelasiP) {
	/*Prosedur untuk menampilkan menu pencarian kata kunci*/
	var opsi int
	var s string
	for exit == false {
		fmt.Println("______________PENCARIAN KATA KUNCI/NILAI______________")
		fmt.Println("1. Pencarian Sesuai Nama")
		fmt.Println("2. Pencarian Sesuai Fasilitas")
		fmt.Println("3. Pencarian Sesuai Wahana")
		fmt.Println("4. Pencarian Sesuai Harga")
		fmt.Println("5. Pencarian Sesuai Jarak")
		fmt.Println("6. Main Menu")
		fmt.Println("======================================================")
		fmt.Print("Input?")
		fmt.Scan(&opsi)
		fmt.Print("\033[H\033[2J")
		if opsi == 6 {
			exit = true
		} else if opsi == 5 {
			PencarianNFW(P, RP, s, opsi)
		} else if opsi == 4 {
			PencarianNFW(P, RP, s, opsi)
		} else if opsi == 3 {
			fmt.Print("Masukan kata kunci:")
			fmt.Scan(&s)
			fmt.Println("Berikut daftar pariwisata yang memiliki wahana", s)
			PencarianNFW(P, RP, s, 3)
		} else if opsi == 2 {
			fmt.Print("Masukan kata kunci:")
			fmt.Scan(&s)
			fmt.Println("Berikut daftar pariwisata yang memiliki fasilitas", s)
			PencarianNFW(P, RP, s, 2)
		} else if opsi == 1 {
			fmt.Print("Masukan Nama:")
			fmt.Scan(&s)
			PencarianNFW(P, RP, s, 1)
		}

	}
}

func MainMenu(exit bool, P *pariwisata, RP *RelasiP) {
	/* Menu utama aplikasi*/
	var opsi int
	var adm string
	for exit == false {
		fmt.Println("______________________MAIN MENU_______________________")
		fmt.Println("1. Admin")
		fmt.Println("2. Daftar Pariwisata")
		fmt.Println("3. Pencarian Kata Kunci/Nilai")
		fmt.Println("4. Keluar               ")
		fmt.Println("======================================================")
		fmt.Print("Input?")
		fmt.Scan(&opsi)
		fmt.Print("\033[H\033[2J")
		if opsi == 4 {
			exit = true
		} else if opsi == 3 {
			PencarianKey(false, *P, *RP)
		} else if opsi == 2 {
			DaftarP(false, *P, *RP)

		} else if opsi == 1 {
			fmt.Println("Masukan Password Admin (passwordnya admin)")
			fmt.Scan(&adm)
			if adm == "admin" {
				Admin(false, P, RP)
				exit = true
			} else {
				fmt.Println("Password Salah")
			}
		}
	}

}

func main() {
	var P pariwisata
	var RP RelasiP
	fmt.Println("======================================================")
	fmt.Println("||                APLIKASI PARIWISATA               ||")
	fmt.Println("||                 TUGAS BESAR ALPRO                ||")
	fmt.Println("======================================================")
	fmt.Println()
	MainMenu(false, &P, &RP)
}
