package main
import (
	"bufio"
	"os"
)

type Income struct {
	Name            string  
	Category        string  
	Amount          float64 
	Date            string  
	OperationalCost float64 
}

var incomes []Income
var reader = bufio.NewReader(os.Stdin)