package main
import (
	"fmt"
	"strings"
)


func sequentialSearchByCategory() {
	printSeparator()
	fmt.Println("  PENCARIAN BERDASARKAN KATEGORI (Sequential Search)")
	printSeparator()

	if len(incomes) == 0 {
		fmt.Println("  [!] Data masih kosong.")
		return
	}


	fmt.Println("  Kategori yang tersedia: Freelance, Passive, Business, Investment")
	category := inputString("  Masukkan kategori yang dicari: ")

	if category == "" {
		fmt.Println("  [!] Kategori tidak boleh kosong.")
		return
	}


	found := false
	count := 0

	printLine()
	fmt.Printf("  Mencari kategori: \"%s\"\n", category)
	printLine()


	for i := 0; i < len(incomes); i++ {
	
		if strings.EqualFold(incomes[i].Category, category) {
		
			if !found {
				fmt.Printf("  %-3s  %-22s  %-12s  %-14s  %-14s\n",
					"No", "Nama Sumber", "Kategori", "Penghasilan", "Tanggal")
				printLine()
				found = true
			}
		
			fmt.Printf("  %-3d  %-22s  %-12s  %-14s  %-14s\n",
				i+1,
				truncateString(incomes[i].Name, 22),
				incomes[i].Category,
				formatRupiah(incomes[i].Amount),
				incomes[i].Date,
			)
			count++
		}
	}

	printLine()
	if !found {
		fmt.Printf("  Kategori \"%s\" tidak ditemukan dalam data.\n", category)
	} else {
		fmt.Printf("  Ditemukan %d data dengan kategori \"%s\".\n", count, category)
	}
}


func binarySearchByName() {
	printSeparator()
	fmt.Println("  PENCARIAN BERDASARKAN NAMA (Binary Search)")
	printSeparator()

	if len(incomes) == 0 {
		fmt.Println("  [!] Data masih kosong.")
		return
	}

	target := strings.ToLower(inputString("  Masukkan nama yang dicari: "))
	if target == "" {
		fmt.Println("  [!] Nama tidak boleh kosong.")
		return
	}

	sorted := make([]Income, len(incomes))
	copy(sorted, incomes)


	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && strings.ToLower(sorted[j].Name) > strings.ToLower(key.Name) {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}

	printLine()
	fmt.Printf("  Data diurutkan A-Z, mencari: \"%s\"\n", target)
	printLine()


	low := 0
	high := len(sorted) - 1
	langkah := 1

	for low <= high {
		mid := (low + high) / 2
		current := strings.ToLower(sorted[mid].Name)

		fmt.Printf("  Langkah %d: Memeriksa indeks %d → \"%s\"\n",
			langkah, mid, sorted[mid].Name)
		langkah++

		if current == target {
		
			printLine()
			fmt.Println(" Data ditemukan!")
			printLine()
			fmt.Printf("  Nama Sumber  : %s\n", sorted[mid].Name)
			fmt.Printf("  Kategori     : %s\n", sorted[mid].Category)
			fmt.Printf("  Tanggal      : %s\n", sorted[mid].Date)
			fmt.Printf("  Penghasilan  : %s\n", formatRupiah(sorted[mid].Amount))
			fmt.Printf("  Biaya Op.    : %s\n", formatRupiah(sorted[mid].OperationalCost))
			fmt.Printf("  Profit       : %s\n", formatRupiah(sorted[mid].Amount-sorted[mid].OperationalCost))
			printLine()
			return

		} else if current < target {
		
			low = mid + 1
			fmt.Printf("           → \"%s\" < target, cari di kanan\n", sorted[mid].Name)
		} else {
		
			high = mid - 1
			fmt.Printf("           → \"%s\" > target, cari di kiri\n", sorted[mid].Name)
		}
	}


	printLine()
	fmt.Printf("  Data dengan nama \"%s\" tidak ditemukan.\n", target)
	fmt.Println("  Pastikan ejaan nama sudah benar (pencarian tidak membedakan huruf besar/kecil).")
}