package main

import (
	"fmt"
	"time"
)

var no = 0

type proyek struct {
	nama     string
	klien    string
	deadline time.Time
	bayaran  int
	status   string
	sisa     int
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
			editdata()
		case 3:
			cari()
		case 4:
			if no <= 0 {
				fmt.Printf("Tidak ada data\n\n")
				return
				tunggu()
				menu(&n)
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
	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Println("1. Tambah Proyek")
	fmt.Println("2. Perbarui Proyek")
	fmt.Println("3. Cari Proyek")
	fmt.Println("4. Tampilkan Proyek")
	fmt.Println("5. Tutup App")
	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Print("Silahkan pilih opsi: ")
	fmt.Scan(n)
	clear()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func tambahproyek() {
	var inptgl string
	fmt.Print("Nama Proyek: ")
	fmt.Scan(&daftar[no].nama)

	fmt.Print("Nama Klien: ")
	fmt.Scan(&daftar[no].klien)

	fmt.Print("Jumlah Bayaran: ")
	fmt.Scan(&daftar[no].bayaran)

	fmt.Print("Status Pengerjaan(Selesai/Pending): ")
	fmt.Scan(&daftar[no].status)

	fmt.Print("Deadline(DD-MM-YYYY): ")
	fmt.Scan(&inptgl)

	format := "02-01-2006"
	t, err := time.Parse(format, inptgl)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Gunakan format DD-MM-YYYY.")
		return
	}
	daftar[no].deadline = t
	tunggu()
	sisawaktu()
	clear()

}

func cari() {
	var n int
	fmt.Println("Cari berdasarkan apa?")
	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Println("1. Nama Klien")
	fmt.Println("2. Nama Proyek")
	fmt.Println("--------------------------------------------------------------------------------------------------")
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
	var key string
	fmt.Print("Masukkan Nama Klien: ")
	fmt.Scan(&key)

	for i := 0; i < no-1; i++ {
		for j := i + 1; j < no; j++ {
			if daftar[i].klien > daftar[j].klien {
				daftar[i], daftar[j] = daftar[j], daftar[i]
			}
		}
	}

	low, high := 0, no-1
	found := false
	head()
	for low <= high && !found {
		mid := (low + high) / 2
		if daftar[mid].klien == key {
			i := mid
			for i >= 0 && daftar[i].klien == key {
				i--
			}
			i++
			size := 1
			for i < no && daftar[i].klien == key {
				fmt.Printf("%-2d | %-24s | %-24s | %-10d | %-8s | %s (sisa %d hari lagi)\n", size, daftar[i].nama, daftar[i].klien, daftar[i].bayaran, daftar[i].status, daftar[i].deadline.Format("02-01-2006"), daftar[i].sisa)
				i++
				size++
			}
			found = true
		} else if daftar[mid].klien < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
	tunggu()
	clear()
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
			fmt.Printf("%-2d | %-24s | %-24s | %-10d | %-8s | %s (%d hari lagi)\n", size, daftar[i].nama, daftar[i].klien, daftar[i].bayaran, daftar[i].status, daftar[i].deadline.Format("02-01-2006"), daftar[i].sisa)
			size++
		}
		i++
	}
	if size <= 0 {
		fmt.Println("Data tidak ditemukan")
	}
	tunggu()
}

func tampil() {
	var i, n int
	head()
	for i < no {
		fmt.Printf("%-2d | %-24s | %-24s | %-10d | %-8s | %s (%d hari lagi)\n", i+1, daftar[i].nama, daftar[i].klien, daftar[i].bayaran, daftar[i].status, daftar[i].deadline.Format("02-01-2006"), daftar[i].sisa)
		i++
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("1. Back to menu")
	fmt.Println("2. Urutkan")
	fmt.Println("3. Hapus")
	fmt.Println("4. Edit")
	fmt.Print("Opsi: ")
	fmt.Scan(&n)
	switch n {
	case 1:
		clear()
		return
	case 2:
		fmt.Println("----------------------------------------------------------------------------------------------------")
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
	case 4:
		editdata()
	default:
		return
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
	tunggu()
	tampil()
}

func sortdeadline() {
	var opsi int
	fmt.Println("Urutkan berdasarkan deadline:")
	fmt.Println("1. Dari terlama ke tercepat")
	fmt.Println("2. Dari tercepat ke terlama")
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
	clear()
	fmt.Println("Data berhasil diurutkan berdasarkan deadline.")
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
	for i := 1; i < no; i++ {
		temp := daftar[i]
		j := i - 1
		if opsi == 1 {
			for j >= 0 && daftar[j].bayaran < temp.bayaran {
				daftar[j+1] = daftar[j]
				j--
			}
		} else {
			for j >= 0 && daftar[j].bayaran > temp.bayaran {
				daftar[j+1] = daftar[j]
				j--
			}
		}
		daftar[j+1] = temp
	}
	clear()
	fmt.Println("Data berhasil diurutkan berdasarkan gaji.")
	tampil()
}

func head() {
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-2s | %-24s | %-24s | %-10s | %-8s | %s\n", "No", "Nama Proyek", "Nama Klien", "Bayaran", "Status", "Deadline")
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}

func sisawaktu() {
	var selisih time.Duration
	for i := 0; i <= no; i++ {
		selisih = daftar[i].deadline.Sub(time.Now())
		daftar[i].sisa = int(selisih.Hours()/24) + 1
	}
}

func editdata() {
	var a int
	var inptgl string
	fmt.Print("Input nomor data yang ingin diubah: ")
	fmt.Scan(&a)
	if a < 1 || a > no {
		fmt.Println("Nomor proyek tidak valid.")
		return
	}

	fmt.Print("Ubah nama proyek: ")
	fmt.Scan(&daftar[a-1].nama)
	fmt.Print("Ubah nama klien: ")
	fmt.Scan(&daftar[a-1].klien)
	fmt.Print("Ubah gaji: ")
	fmt.Scan(&daftar[a-1].bayaran)
	fmt.Print("Ubah status: ")
	fmt.Scan(&daftar[a-1].status)
	fmt.Print("Ubah deadline(DD-MM-YYYY): ")
	fmt.Scan(&inptgl)
	fmt.Println("Data berhasil diubah.")
	format := "02-01-2006"
	t, err := time.Parse(format, inptgl)
	if err != nil {
		fmt.Println("Format tanggal tidak valid. Gunakan format DD-MM-YYYY.")
		return
	}
	daftar[a-1].deadline = t
	sisawaktu()
	tampil()
}

func tunggu() {
	fmt.Print("Ketik 'ok' untuk melanjutkan: ")
	var s string
	fmt.Scan(&s)
	for s != "ok" {
		fmt.Print("Silahkan ketik 'ok' untuk melanjutkan: ")
		fmt.Scan(&s)
	}
}
