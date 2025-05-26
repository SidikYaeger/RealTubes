package main

import (
	"fmt"
)

var no = 0

type proyek struct {
	nama    string
	klien   string
	bayaran int
	status  string
	sisa    int
}

var daftar [1000]proyek

func main() {
	var n int
	fmt.Println("App Manajemen dan Tracking Kegiatan Freelance(ANJAY)")
	fmt.Println("Made By Rahman n Grandito")
	fmt.Print("\n\n")
	for n != 6 {
		menu(&n)
		clear()
		switch n {
		case 1:
			tambahproyek()
			no++
		case 2:
			if no <= 0 {
				fmt.Printf("tidak ada data\n\n")
				return
			}
			updatestatus()
		case 3:
			cari()
		case 4:
			if no <= 0 {
				fmt.Printf("tidak ada data\n\n")
				return
			}
			tampil()
		case 5:
			return
		}
	}

}

func menu(n *int) {
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃  ███╗   ███╗███████╗███╗   ██╗██╗   ██╗ ┃")
	fmt.Println("┃  ████╗ ████║██╔════╝████╗  ██║██║   ██║ ┃")
	fmt.Println("┃  ██╔████╔██║█████╗  ██╔██╗ ██║██║   ██║ ┃")
	fmt.Println("┃  ██║╚██╔╝██║██╔══╝  ██║╚██╗██║██║   ██║ ┃")
	fmt.Println("┃  ██║ ╚═╝ ██║███████╗██║ ╚████║████████║ ┃")
	fmt.Println("┃  ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝ ╚══════╝ ┃")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("1. Tambah Proyek")
	fmt.Println("2. Perbarui/edit Proyek")
	fmt.Println("3. Cari Proyek")
	fmt.Println("4. Tampilkan Proyek")
	fmt.Println("5. Tutup App")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Print("Silahkan Pilih Opsi: ")
	fmt.Scan(n)
	clear()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func tambahproyek() {

	fmt.Print("Nama Proyek: ")
	fmt.Scan(&daftar[no].nama)

	fmt.Print("Nama Klien: ")
	fmt.Scan(&daftar[no].klien)

	fmt.Print("Jumlah Bayaran: ")
	fmt.Scan(&daftar[no].bayaran)

	fmt.Print("Status: ")
	fmt.Scan(&daftar[no].status)

	fmt.Print("Deadline(sisa hari): ")
	fmt.Scan(&daftar[no].sisa)

	clear()
}

func cari() {
	var n int
	fmt.Println("Cari pakai apa?")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("1. Nama Klien")
	fmt.Println("2. Nama Proyek")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Print("Opsi:")
	fmt.Scan(&n)
	clear()
	switch n {
	case 1:
		cariklien()
	case 2:
		cariproyek()
	}
}

func cariklien() {
	var i, size int = 0, 1
	var key string
	fmt.Print("Masukkan Nama Klien: ")
	fmt.Scan(&key)

	head()
	for i < no {
		if daftar[i].klien == key {
			fmt.Printf("%-2d | %-15s | %-15s | %-10d | %-8s | %d hari lagi\n", size, daftar[i].nama, daftar[i].klien, daftar[i].bayaran, daftar[i].status, daftar[i].sisa)
			size++
		}
		i++
	}

}

func cariproyek() {
	var i, size int = 0, 1
	var key string
	fmt.Print("Masukkan Nama proyek: ")
	fmt.Scan(&key)
	clear()
	head()
	for i < no {
		if daftar[i].nama == key {
			fmt.Printf("%-2d | %-15s | %-15s | %-10d | %-8s | %d hari lagi\n", size, daftar[i].nama, daftar[i].klien, daftar[i].bayaran, daftar[i].status, daftar[i].sisa)
			size++
		}
		i++
	}

}

func tampil() {
	var i, n int
	head()
	for i < no {
		fmt.Printf("%-2d | %-15s | %-15s | %-10d | %-8s | %d hari lagi\n", i+1, daftar[i].nama, daftar[i].klien, daftar[i].bayaran, daftar[i].status, daftar[i].sisa)
		i++
	}
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("1. Back to menu")
	fmt.Println("2. sort")
	fmt.Println("3. Hapus")
	fmt.Print("Opsi: ")
	fmt.Scan(&n)
	switch n {
	case 1:
		clear()
		return
	case 2:
		fmt.Println("--------------------------------------------------------------------------")
		fmt.Println("1. Deadline ")
		fmt.Println("2. Gaji")
		fmt.Println("3. Back to menu")
		fmt.Print("Opsi: ")
		fmt.Scan(&n)
		clear()
		switch n {
		case 1:
			sortdeadline()
		case 2:
			sortgaji()
		default:
			return
		}
	case 3:
		hapus()
	}

}

func hapus() {
	var i int
	fmt.Print("Masukkan nomor proyek yang ingin dihapus: ")
	fmt.Scan(&i)
	for j := i - 1; j < no-1; j++ {
		daftar[j] = daftar[j+1]
	}

	no--
	fmt.Println("Proyek berhasil dihapus.")
	tampil()
}

func sortdeadline() {
	var opsi int
	fmt.Println("Urutkan berdasarkan deadline:")
	fmt.Println("1. Dari terbesar ke terkecil")
	fmt.Println("2. Dari terkecil ke terbesar")
	fmt.Print("Opsi: ")
	fmt.Scan(&opsi)
	if opsi != 1 && opsi != 2 {
		fmt.Println("Invalid")
		return
	}
	clear()
	for i := 0; i < no-1; i++ {
		temp := i
		for j := i + 1; j < no; j++ {
			if (opsi == 1 && daftar[j].sisa > daftar[temp].sisa) || (opsi == 2 && daftar[j].sisa < daftar[temp].sisa) {
				temp = j
			}
		}
		if temp != i {
			daftar[i], daftar[temp] = daftar[temp], daftar[i]
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan gaji.")
	tampil()
}

func sortgaji() {
	var opsi int
	fmt.Println("Urutkan berdasarkan gaji:")
	fmt.Println("1. Dari terbesar ke terkecil")
	fmt.Println("2. Dari terkecil ke terbesar")
	fmt.Print("Opsi: ")
	fmt.Scan(&opsi)
	if opsi != 1 && opsi != 2 {
		fmt.Println("Invalid")
		return
	}
	clear()
	for i := 0; i < no-1; i++ {
		temp := i
		for j := i + 1; j < no; j++ {
			if (opsi == 1 && daftar[j].bayaran > daftar[temp].bayaran) || (opsi == 2 && daftar[j].bayaran < daftar[temp].bayaran) {
				temp = j
			}
		}
		if temp != i {
			daftar[i], daftar[temp] = daftar[temp], daftar[i]
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan gaji.")
	tampil()
}

func head() {
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("No | Nama Proyek     | Nama Klien      | Gaji       | Status   | Deadline")
	fmt.Println("--------------------------------------------------------------------------")
}

func updatestatus() {
	var a int
	fmt.Print("Input nomor data yang ingin diubah: ")
	fmt.Scan(&a)
	fmt.Print("Ubah nama proyek: ")
	fmt.Scan(&daftar[a-1].nama)
	fmt.Print("Ubah nama klien: ")
	fmt.Scan(&daftar[a-1].klien)
	fmt.Print("Ubah gaji: ")
	fmt.Scan(&daftar[a-1].bayaran)
	fmt.Print("Ubah status: ")
	fmt.Scan(&daftar[a-1].status)
	fmt.Print("Ubah deadline(sisa hari): ")
	fmt.Scan(&daftar[a-1].sisa)
	fmt.Println("Data berhasil diubah.")
	tampil()
}
