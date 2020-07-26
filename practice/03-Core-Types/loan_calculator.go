// Activity 3.02

package main

import (
	"errors"
	"fmt"
)

const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("invalid income")
	ErrLoanAmount  = errors.New("invalid loan amount")
	ErrLoanTerm    = errors.New("invalid loan term")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	// Good credit card with better rate
	interest := 20.0
	if creditScore >= goodScore {
		interest = 15.0
	}

	// validate score
	if creditScore < 1 {
		return ErrCreditScore
	}

	// validate income
	if income < 1 {
		return ErrIncome
	}

	// validate loan amount
	if loanAmount < 1 {
		return ErrLoanAmount
	}

	// validate loan term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	// total cost of the loan
	totalInterest := (payment * loanTerm) - loanAmount

	// rules
	approved := false
	if income > payment {
		// payment percent of income validation
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = true
		}
	}

	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}

func main() {
	// approved case
	fmt.Println("Applicant 1")
	fmt.Println("------------------------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// denied case
	fmt.Println("Applicant 2")
	fmt.Println("---------------------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
