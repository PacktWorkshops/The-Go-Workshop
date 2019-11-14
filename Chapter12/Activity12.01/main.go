package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type budgetCategory string

const (
	autoFuel   budgetCategory = "fuel"
	food       budgetCategory = "food"
	mortgage   budgetCategory = "mortgage"
	repairs    budgetCategory = "repairs"
	insurance  budgetCategory = "insurance"
	utilities  budgetCategory = "utilities"
	retirement budgetCategory = "retirement"
)

var (
	ErrInvalidBudgetCategory = errors.New("budget category not found")
)

type transaction struct {
	id       int
	payee    string
	spent    float64
	category budgetCategory
}

func main() {
	bankFile := flag.String("c", "", "location of the bank transaction csv file")
	logFile := flag.String("l", "", "logging of errors")
	flag.Parse()

	if *bankFile == "" {
		fmt.Println("csvFile is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *logFile == "" {
		fmt.Println("logFile is required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := os.Stat(*bankFile)
	if os.IsNotExist((err)) {
		fmt.Println("BankFile does not exist: ", *bankFile)
		os.Exit(1)
	}

	_, err = os.Stat(*logFile)
	if !os.IsNotExist((err)) {
		os.Remove(*logFile)
	}
 
	csvFile, err := os.Open(*bankFile)
	if err != nil {
		fmt.Println("Error opening file: ", *bankFile)
		os.Exit(1)
	}
	trxs := parseBankFile(csvFile, *logFile)
    	fmt.Println()
    	for _, trx := range trxs {
        fmt.Printf("%v\n", trx)
    }

}

func parseBankFile(bankTransactions io.Reader, logFile string) []transaction {
	r := csv.NewReader(bankTransactions)
	trxs := []transaction{}
	header := true

	for {
		trx := transaction{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !header {
			for idx, value := range record {
				switch idx {
				// id
				case 0:
					value = strings.TrimSpace(value)
					trx.id, err = strconv.Atoi(value)
					if err != nil {
						log.Fatal(err)
					}
				// payee
				case 1:
					value = strings.TrimSpace(value)
					trx.payee = value
				// spent
				case 2:
					value = strings.TrimSpace(value)
					trx.spent, err = strconv.ParseFloat(value, 64)
					if err != nil {
						log.Fatal(err)
					}
				// category
				case 3:
					trx.category, err = convertToBudgetCategory(value)
					if err != nil {
						s := strings.Join(record, ", ")
						writeErrorToLog("error converting csv category column - ", err, s, logFile)
					}
				}
			}
			trxs = append(trxs, trx)
		}
		header = false
	}
	return trxs
}

func writeErrorToLog(msg string, err error, data string, logFile string) error {
	msg += "\n" + err.Error() + "\nData: " + data + "\n\n"
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(msg); err != nil {
		return err
	}
	return nil
}

func convertToBudgetCategory(value string) (budgetCategory, error) {
	value = strings.TrimSpace(strings.ToLower(value))
	switch value {
	case "fuel", "gas":
		return autoFuel, nil
	case "food":
		return food, nil
	case "mortgage":
		return mortgage, nil
	case "repairs":
		return repairs, nil
	case "car insurance", "life insurance":
		return insurance, nil
	case "utilities":
		return utilities, nil
	default:
		return "", ErrInvalidBudgetCategory
	}
}
