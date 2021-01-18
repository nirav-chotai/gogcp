// Activity 3.02

package coretypes

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
	errCreditScore = errors.New("invalid credit score")
	errIncome      = errors.New("invalid income")
	errLoanAmount  = errors.New("invalid loan amount")
	errLoanTerm    = errors.New("invalid loan term")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	// Good credit card with better rate
	interest := 20.0
	if creditScore >= goodScore {
		interest = 15.0
	}

	// validate score
	if creditScore < 1 {
		return errCreditScore
	}

	// validate income
	if income < 1 {
		return errIncome
	}

	// validate loan amount
	if loanAmount < 1 {
		return errLoanAmount
	}

	// validate loan term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return errLoanTerm
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

func loan() {
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
