package main


import (
	"fmt"
	"strings"
)

func monthlyReport() {
	printSeparator()
	fmt.Println("  LAPORAN BULANAN SIDE HUSTLE & PASSIVE INCOME")
	printSeparator()

	if len(incomes) == 0 {
		fmt.Println("  [!] Belum ada data untuk dilaporkan.")
		return
	}

				totalIncome := calculateTotalIncome()
	totalCost := calculateTotalCost()
	totalProfit := calculateTotalProfit()
	marginPersen := 0.0
	if totalIncome > 0 {
		marginPersen = (totalProfit / totalIncome) * 100
	}

	fmt.Println()
	fmt.Println("  [ RINGKASAN KESELURUHAN ]")
	printLine()
	fmt.Printf("  Total Penghasilan Kotor  : %s\n", formatRupiah(totalIncome))
	fmt.Printf("  Total Biaya Operasional  : %s\n", formatRupiah(totalCost))
	fmt.Printf("  Total Laba Bersih        : %s\n", formatRupiah(totalProfit))
	fmt.Printf("  Margin Keuntungan        : %.1f%%\n", marginPersen)
	fmt.Printf("  Jumlah Sumber Penghasilan: %d entri\n", len(incomes))

				fmt.Println()
	fmt.Println("  [ ANALISIS PER KATEGORI ]")
	printLine()

		categories := getUniqueCategories()

	fmt.Printf("  %-15s  %-6s  %-14s  %-14s  %-12s\n",
		"Kategori", "Jumlah", "Total Income", "Total Biaya", "Total Profit")
	printLine()

	for _, cat := range categories {
		catIncome, catCost, catProfit, catCount := summarizeByCategory(cat)
		fmt.Printf("  %-15s  %-6d  %-14s  %-14s  %-12s\n",
			cat, catCount,
			formatRupiah(catIncome),
			formatRupiah(catCost),
			formatRupiah(catProfit),
		)
	}

				bestIdx, worstIdx := findBestAndWorst()

	fmt.Println()
	fmt.Println("  [ SUMBER PENGHASILAN TERBAIK & TERBURUK ]")
	printLine()

	bestProfit := incomes[bestIdx].Amount - incomes[bestIdx].OperationalCost
	worstProfit := incomes[worstIdx].Amount - incomes[worstIdx].OperationalCost

	fmt.Printf("  🏆 Terbaik  : %-22s → Profit %s\n",
		incomes[bestIdx].Name, formatRupiah(bestProfit))
	fmt.Printf("  📉 Terburuk : %-22s → Profit %s\n",
		incomes[worstIdx].Name, formatRupiah(worstProfit))

				fmt.Println()
	fmt.Println("  [ REKOMENDASI OPTIMASI PEMASUKAN ]")
	printLine()

	generateRecommendations(totalProfit, totalIncome, marginPersen)

	printSeparator()
}

func getUniqueCategories() []string {
	var categories []string

	for _, inc := range incomes {
		found := false
				for _, cat := range categories {
			if strings.EqualFold(cat, inc.Category) {
				found = true
				break
			}
		}
				if !found {
			categories = append(categories, inc.Category)
		}
	}

	return categories
}

func summarizeByCategory(category string) (income, cost, profit float64, count int) {
	for _, inc := range incomes {
		if strings.EqualFold(inc.Category, category) {
			income += inc.Amount
			cost += inc.OperationalCost
			profit += inc.Amount - inc.OperationalCost
			count++
		}
	}
	return
}

func findBestAndWorst() (bestIdx, worstIdx int) {
	bestIdx = 0
	worstIdx = 0

	bestProfit := incomes[0].Amount - incomes[0].OperationalCost
	worstProfit := incomes[0].Amount - incomes[0].OperationalCost

	for i := 1; i < len(incomes); i++ {
		profit := incomes[i].Amount - incomes[i].OperationalCost

		if profit > bestProfit {
			bestProfit = profit
			bestIdx = i
		}

		if profit < worstProfit {
			worstProfit = profit
			worstIdx = i
		}
	}

	return
}

func generateRecommendations(totalProfit, totalIncome, margin float64) {
		if totalProfit >= 10_000_000 {
		fmt.Println(" Penghasilan Anda sudah sangat baik!")
		fmt.Println(" → Pertimbangkan untuk berinvestasi kembali ke bisnis yang paling menguntungkan.")
		fmt.Println(" → Diversifikasi ke instrumen investasi (reksa dana, saham) untuk pertumbuhan pasif.")

	} else if totalProfit >= 5_000_000 {
		fmt.Println("Performa keuangan Anda cukup baik.")
		fmt.Println("→ Pertahankan strategi yang ada dan tingkatkan kapasitas produksi.")
		fmt.Println("→ Coba tambah 1-2 sumber penghasilan baru di kategori yang sudah terbukti.")

	} else if totalProfit >= 0 {
		fmt.Println("Profit masih positif, tapi ada ruang untuk ditingkatkan.")
		fmt.Println("→ Fokus pada sumber penghasilan dengan margin tertinggi.")
		fmt.Println("→ Kurangi atau eliminasi sumber penghasilan yang profit-nya kecil.")

	} else {
		fmt.Println("Total profit negatif! Segera lakukan evaluasi.")
		fmt.Println("→ Identifikasi sumber penghasilan yang paling banyak merugikan.")
		fmt.Println("→ Kurangi biaya operasional yang tidak perlu.")
		fmt.Println("→ Pertimbangkan menghentikan sumber penghasilan yang merugi.")
	}

		fmt.Println()
	if margin >= 70 {
		fmt.Println("Margin keuntungan sangat tinggi (>70%) — bisnis Anda sangat efisien!")
	} else if margin >= 50 {
		fmt.Println("Margin keuntungan baik (50-70%) — terus optimalkan pengeluaran.")
	} else if margin >= 30 {
		fmt.Println("Margin keuntungan cukup (30-50%) — ada potensi untuk meningkatkan efisiensi.")
	} else if margin > 0 {
		fmt.Println("Margin keuntungan rendah (<30%) — perlu evaluasi serius pada biaya operasional.")
	}

		for _, inc := range incomes {
		if inc.Amount < inc.OperationalCost {
			fmt.Printf("\n Perhatian: \"%s\" sedang merugi %s!\n",
				inc.Name, formatRupiah(inc.OperationalCost-inc.Amount))
		}
	}
}
