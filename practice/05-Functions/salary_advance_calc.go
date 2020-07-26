// Activity 5.02

package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota // starts with zero
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	d := Developer{Individual: Employee{Id: 1, FirstName: "Nirav", LastName: "Chotai"}, HourlyRate: 10}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()
	d.LoggedHours(Monday, 8)
	d.LoggedHours(Tuesday, 10)
	d.LoggedHours(Wednesday, 10)
	d.LoggedHours(Thursday, 10)
	d.LoggedHours(Friday, 6)
	d.LoggedHours(Saturday, 8)
	d.PayDetails()
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (d *Developer) LoggedHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func (d *Developer) PayDay() (int, bool) {
	if d.HoursWorked() > 40 {
		hoursOver := d.HoursWorked() - 40
		overTime := hoursOver * 2
		regularPay := d.HoursWorked() * d.HourlyRate
		return regularPay + overTime, true
	}
	return d.HoursWorked() * d.HourlyRate, false
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
		fmt.Printf("\nHours worked this week: %d\n", d.HoursWorked())
		pay, overtime := d.PayDay()
		fmt.Println("Pay for this week: $", pay)
		fmt.Println("Is this overtime pay: ", overtime)
		fmt.Println()
	}
}
