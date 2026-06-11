package main
import "fmt"

func seedData() {
	incomes = []Income{
		{
			Name:            "Freelance Design",
			Category:        "Freelance",
			Amount:          2_500_000,
			Date:            "01-05-2025",
			OperationalCost: 500_000,
		},
		{
			Name:            "YouTube Ads",
			Category:        "Passive",
			Amount:          1_500_000,
			Date:            "05-05-2025",
			OperationalCost: 100_000,
		},
		{
			Name:            "Affiliate Marketing",
			Category:        "Passive",
			Amount:          2_000_000,
			Date:            "10-05-2025",
			OperationalCost: 300_000,
		},
		{
			Name:            "Online Shop",
			Category:        "Business",
			Amount:          4_000_000,
			Date:            "12-05-2025",
			OperationalCost: 1_500_000,
		},
		{
			Name:            "Kursus Online",
			Category:        "Passive",
			Amount:          3_200_000,
			Date:            "15-05-2025",
			OperationalCost: 200_000,
		},
		{
			Name:            "Saham Dividen",
			Category:        "Investment",
			Amount:          800_000,
			Date:            "20-05-2025",
			OperationalCost: 50_000,
		},
		{
			Name:            "Jasa Copywriting",
			Category:        "Freelance",
			Amount:          1_800_000,
			Date:            "22-05-2025",
			OperationalCost: 0,
		},
	}
}


func showMenu() {
	fmt.Println()
	printSeparator()
	fmt.Println("       APLIKASI PELACAK SIDE HUSTLE & PASSIVE INCOME")
	printSeparator()
	fmt.Println("  [CRUD]")
	fmt.Println("  1. Tambah Data Penghasilan")
	fmt.Println("  2. Tampilkan Semua Data")
	fmt.Println("  3. Update Data Penghasilan")
	fmt.Println("  4. Hapus Data Penghasilan")
	printLine()
	fmt.Println("  [PENCARIAN]")
	fmt.Println("  5. Sequential Search (berdasarkan Kategori)")
	fmt.Println("  6. Binary Search     (berdasarkan Nama)")
	printLine()
	fmt.Println("  [PENGURUTAN]")
	fmt.Println("  7. Selection Sort (urutkan berdasarkan Penghasilan)")
	fmt.Println("  8. Insertion Sort (urutkan berdasarkan Kategori)")
	printLine()
	fmt.Println("  [LAPORAN]")
	fmt.Println("  9. Laporan Bulanan & Rekomendasi")
	printLine()
	fmt.Println("  0. Keluar")
	printSeparator()
}

func main() {
	fmt.Println()
	printSeparator()
	fmt.Println("  Selamat datang di Aplikasi Pelacak Side Hustle!")
	fmt.Println("  Data awal berhasil dimuat. Silakan pilih menu.")
	printSeparator()

	// ini buat Isi data awal agar program bisa langsung diuji
	seedData()

	
	for {
		showMenu()
		pilihan := inputInt("  Pilih menu (0-9): ")

		switch pilihan {
		case 1:
			addIncome()

		case 2:
			showAll()

		case 3:
			updateIncome()

		case 4:
			deleteIncome()

		case 5:
			sequentialSearchByCategory()

		case 6:
			binarySearchByName()

		case 7:
			selectionSortByAmount()

		case 8:
			insertionSortByCategory()

		case 9:
			monthlyReport()

		case 0:
			printSeparator()
			fmt.Println("  Terima kasih telah menggunakan Aplikasi Pelacak Side Hustle!")
			fmt.Println("  Semoga side hustle Anda semakin berkembang! 🚀")
			printSeparator()
			return

		default:
			fmt.Println("  [!] Pilihan tidak valid. Masukkan angka 0 sampai 9.")
		}
		pause()
	}
}