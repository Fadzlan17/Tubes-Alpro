package main
import (
	"fmt"
	"strings"
)

func inputString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputFloat(prompt string) float64 {
	var value float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%f", &value)
	
		reader.ReadString('\n')
		if err == nil && value >= 0 {
			break
		}
		fmt.Println("  [!] Input tidak valid. Masukkan angka positif.")
	}
	return value
}

func inputInt(prompt string) int {
	var value int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%d", &value)
		reader.ReadString('\n')
		if err == nil {
			break
		}
		fmt.Println("  [!] Input tidak valid. Masukkan bilangan bulat.")
	}
	return value
}

func printSeparator() {
	fmt.Println(strings.Repeat("=", 75))
}

func printLine() {
	fmt.Println(strings.Repeat("-", 75))
}

func pause() {
	fmt.Println("\n  Tekan Enter untuk kembali ke menu...")
	reader.ReadString('\n')
}

func formatRupiah(amount float64) string {

	intPart := int64(amount)
	decPart := int64((amount - float64(intPart)) * 100)


	str := fmt.Sprintf("%d", intPart)
	result := ""
	for i, ch := range str {
	
		posFromRight := len(str) - i - 1
		if posFromRight > 0 && posFromRight%3 == 0 {
			result += "."
		}
		result += string(ch)
	}

	return fmt.Sprintf("Rp %s,%02d", result, decPart)
}