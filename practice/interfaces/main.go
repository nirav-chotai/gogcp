package main

import (
	"errors"
	"fmt"
	"os"
)

type employee struct {
	ID        int
	FirstName string
	LastName  string
}

type developer struct {
	Individual        employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type manager struct {
	Individual     employee
	Salary         float64
	CommissionRate float64
}

// Payer interface exported
type Payer interface {
	Pay() (string, float64)
}

func (d developer) FullName() string {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName
}
func (d developer) Pay() (string, float64) {
	return d.FullName(), d.HourlyRate * d.HoursWorkedInYear
}

func (m manager) Pay() (string, float64) {
	fullName := m.Individual.FirstName + " " + m.Individual.LastName
	return fullName, m.Salary + (m.Salary * m.CommissionRate)
}

func payDetails(p Payer) {
	fullName, yearPay := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", fullName, yearPay)
}

func convertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("Invalid rating: " + str)
	}
}

func overallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("Unknown type")
	}
}

func (d developer) reviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := overallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.FullName(), averageRating)
	return nil
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := developer{Individual: employee{ID: 1, FirstName: "Nirav", LastName: "Chotai"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := manager{Individual: employee{ID: 2, FirstName: "Gunja", LastName: "Chotai"}, Salary: 150000, CommissionRate: .07}

	err := d.reviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}
