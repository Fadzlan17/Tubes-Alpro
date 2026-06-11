package main

import (
	"fmt"
	"strings"
)

func selectionSortByAmount() {
	printSeparator()
	fmt.Println("  PENGURUTAN BERDASARKAN PENGHASILAN (Selection Sort)")
	printSeparator()

	n := len(incomes)

	if n == 0 {
		fmt.Println("  [!] Data masih kosong.")
		return
	}

	if n == 1 {
		fmt.Println("  Data hanya 1 elemen, tidak perlu diurutkan.")
		return
	}

	fmt.Printf("  Mengurutkan %d data...\n", n)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if incomes[j].Amount < incomes[minIndex].Amount {
				minIndex = j
			}
		}

	
		if minIndex != i {
			incomes[i], incomes[minIndex] = incomes[minIndex], incomes[i]
		}
	}

	fmt.Println("Data berhasil diurutkan dari penghasilan terkecil ke terbesar!")
	fmt.Println()

	showAll()
}


func insertionSortByCategory() {
	printSeparator()
	fmt.Println("  PENGURUTAN BERDASARKAN KATEGORI (Insertion Sort)")
	printSeparator()

	n := len(incomes)

	if n == 0 {
		fmt.Println("  [!] Data masih kosong.")
		return
	}

	if n == 1 {
		fmt.Println("  Data hanya 1 elemen, tidak perlu diurutkan.")
		return
	}

	fmt.Printf("  Mengurutkan %d data berdasarkan kategori (A-Z)...\n", n)


	for i := 1; i < n; i++ {
		key := incomes[i]
		j := i - 1
	
		for j >= 0 && strings.ToLower(incomes[j].Category) > strings.ToLower(key.Category) {
			incomes[j+1] = incomes[j]
			j--
		}
		incomes[j+1] = key
	}

	fmt.Println(" Data berhasil diurutkan berdasarkan kategori (A-Z)!")
	fmt.Println()
	showAll()
}