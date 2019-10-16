package main
import "fmt"
var budgetCategories = make(map[int]string)
var payeeToCategory = make(map[string]int)
func init() {
	fmt.Println("Initializing our budgetCategories")
	budgetCategories[1] = "Car Insurance"
	budgetCategories[2] = "Mortgage"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[7] = "Groceries"
	budgetCategories[8] = "Car Payment"
}
func init() {
	fmt.Println("Assign our Payees to categories")
	payeeToCategory["Nationwide"] = 1
	payeeToCategory["BBT Loan"] = 2
	payeeToCategory["First Energy Electric"] = 3
	payeeToCategory["Ameriprise Financial"] = 4
	payeeToCategory["Walt Disney World"] = 5
	payeeToCategory["ALDI"] = 7
	payeeToCategory["Martins"] = 7
	payeeToCategory["Wal Mart"] = 7
	payeeToCategory["Chevy Loan"] = 8
}
func main() {
	fmt.Println("In main, printing payee to category")
	for k, v := range payeeToCategory {
		fmt.Printf("Payee: %s, Category: %s\n", k, budgetCategories[v])
	}
}
