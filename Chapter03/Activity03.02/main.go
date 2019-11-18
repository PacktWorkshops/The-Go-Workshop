package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	goodScore       = 450
	lowScoreRation  = 10
	goodScoreRation = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term not a multiple of 12")
)

func main() {
	// Approved
	score := 500
	income := 1000.0
	loanAmount := 1000.0
	loanTerm := 24.0
	totalInterest, payment, approved, err := checkLoan(score, income, loanAmount, loanTerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Applicant 1")
	fmt.Println("Credit Score    :", score)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Score           :", score)
	fmt.Println("Approved        :", approved)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("")

	// Denied
	score = 350
	income = 1000.0
	loanAmount = 10000.0
	loanTerm = 12.0
	totalInterest, payment, approved, err = checkLoan(score, income, loanAmount, loanTerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Applicant 2")
	fmt.Println("Credit Score    :", score)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Score           :", score)
	fmt.Println("Approved        :", approved)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Total Cost      :", totalInterest)

}

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) (totalInterest float64, payment float64, approved bool, err error) {
	// Good credit score gets a better rate
	interest := 20.0
	if creditScore >= goodScore {
		interest = 15.0
	}
	//validate score
	if creditScore < 1 {
		err = ErrCreditScore
		return
	}
	// validate income
	if income < 1 {
		err = ErrIncome
		return
	}
	// validate loanAmount
	if loanAmount < 1 {
		err = ErrLoanAmount
		return
	}
	//validate term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		err = ErrLoanTerm
		return
	}

	// Convert interest % to a monthly rate we can use
	rate := interest / 100 / 12
	// compute the monthly payment figure
	x := math.Pow(1+rate, loanTerm)
	payment = (loanAmount * x * rate) / (x - 1)
	// Round to a double
	payment = math.Round(payment*100) / 100
	// Total cost of the loan
	totalInterest = (payment * loanTerm) - loanAmount
	totalInterest = math.Round(totalInterest*100) / 100
	// Can they afford the according to our rules?
	if income < payment {
		return
	} else {
		// payment percent of income
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio > goodScoreRation {
			return
		} else if ratio > lowScoreRation {
			return
		}
	}
	// Everything looks good
	approved = true
	return
}
