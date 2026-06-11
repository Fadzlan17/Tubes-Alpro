package main

import "fmt"

func printTableHeader() {
	printSeparator()
	fmt.Printf("  %-3s  %-22s  %-12s  %-14s  %-14s  %-12s\n",
		"No", "Nama Sumber", "Kategori", "Penghasilan", "Biaya Op.", "Profit")
	printLine()
}

func printTableRow(index int, inc Income) {
	profit := inc.Amount - inc.OperationalCost
	fmt.Printf("  %-3d  %-22s  %-12s  %-14s  %-14s  %-12s\n",
		index,
		truncateString(inc.Name, 22),
		truncateString(inc.Category, 12),
		formatRupiah(inc.Amount),
		formatRupiah(inc.OperationalCost),
		formatRupiah(profit),
	)
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-2] + ".."
}

func addIncome() {
	printSeparator()
	fmt.Println("  TAMBAH DATA PENGHASILAN BARU")
	printSeparator()

	name := inputString("  Nama sumber penghasilan : ")
	if name == "" {
		fmt.Println("  [!] Nama tidak boleh kosong.")
		return
	}

	category := inputString("  Kategori (Freelance/Passive/Business/Investment) : ")
	if category == "" {
		fmt.Println("  [!] Kategori tidak boleh kosong.")
		return
	}

	date := inputString("  Tanggal (DD-MM-YYYY)    : ")
	if date == "" {
		date = "01-01-2025"
	}

	amount := inputFloat("  Jumlah penghasilan (Rp) : ")
	cost := inputFloat("  Biaya operasional  (Rp) : ")

	incomes = append(incomes, Income{
		Name:            name,
		Category:        category,
		Amount:          amount,
		Date:            date,
		OperationalCost: cost,
	})

	printLine()
	fmt.Println(" Data berhasil ditambahkan!")
}

func showAll() {
	printSeparator()
	fmt.Println("  DAFTAR SEMUA SUMBER PENGHASILAN")
	printTableHeader()

	if len(incomes) == 0 {
		fmt.Println("  (Belum ada data. Silakan tambah data terlebih dahulu.)")
		printSeparator()
		return
	}

	for i, inc := range incomes {
		printTableRow(i+1, inc)
	}

	printLine()
	fmt.Printf("  %-3s  %-22s  %-12s  %-14s  %-14s  %-12s\n",
		"", "TOTAL", "",
		formatRupiah(calculateTotalIncome()),
		formatRupiah(calculateTotalCost()),
		formatRupiah(calculateTotalProfit()),
	)
	printSeparator()
	fmt.Printf("  Total data: %d entri\n", len(incomes))
}

func updateIncome() {
	showAll()

	if len(incomes) == 0 {
		return
	}

	printLine()
	index := inputInt("  Pilih nomor data yang ingin diubah (0 = batal): ")

	if index == 0 {
		fmt.Println("  Update dibatalkan.")
		return
	}

	if index < 1 || index > len(incomes) {
		fmt.Println("  [!] Nomor tidak valid. Pilih antara 1 sampai", len(incomes))
		return
	}

	i := index - 1

	printSeparator()
	fmt.Printf("  Mengedit data: %s\n", incomes[i].Name)
	fmt.Println("  (Kosongkan dan tekan Enter untuk mempertahankan nilai lama)")
	printLine()

	name := inputString(fmt.Sprintf("  Nama [%s]: ", incomes[i].Name))
	if name != "" {
		incomes[i].Name = name
	}

	category := inputString(fmt.Sprintf("  Kategori [%s]: ", incomes[i].Category))
	if category != "" {
		incomes[i].Category = category
	}

	date := inputString(fmt.Sprintf("  Tanggal [%s]: ", incomes[i].Date))
	if date != "" {
		incomes[i].Date = date
	}

	amount := inputFloat(fmt.Sprintf("  Penghasilan [%.0f] (0 = tidak ubah): ", incomes[i].Amount))
	if amount > 0 {
		incomes[i].Amount = amount
	}

	cost := inputFloat(fmt.Sprintf("  Biaya Op. [%.0f] (0 = tidak ubah): ", incomes[i].OperationalCost))
	if cost > 0 {
		incomes[i].OperationalCost = cost
	}

	printLine()
	fmt.Println(" Data berhasil diperbarui!")
}

func deleteIncome() {
	showAll()

	if len(incomes) == 0 {
		return
	}

	printLine()
	index := inputInt("  Pilih nomor data yang ingin dihapus (0 = batal): ")

	if index == 0 {
		fmt.Println("  Penghapusan dibatalkan.")
		return
	}

	if index < 1 || index > len(incomes) {
		fmt.Println("  [!] Nomor tidak valid.")
		return
	}

	i := index - 1

	fmt.Printf("\n  Data yang akan dihapus: %s (%s)\n", incomes[i].Name, incomes[i].Category)
	konfirmasi := inputString("  Yakin ingin menghapus? (y/n): ")

	if konfirmasi != "y" && konfirmasi != "Y" {
		fmt.Println("  Penghapusan dibatalkan.")
		return
	}

	// ini buat Hapus elemen dari slice pake teknik append sama ngegabung elemen sebelum index dan sesudah index
	incomes = append(incomes[:i], incomes[i+1:]...)

	fmt.Println("  Data berhasil dihapus!")
}

func calculateTotalIncome() float64 {
	total := 0.0
	for _, inc := range incomes {
		total += inc.Amount
	}
	return total
}

func calculateTotalCost() float64 {
	total := 0.0
	for _, inc := range incomes {
		total += inc.OperationalCost
	}
	return total
}

func calculateTotalProfit() float64 {
	total := 0.0
	for _, inc := range incomes {
		total += inc.Amount - inc.OperationalCost
	}
	return total
}
