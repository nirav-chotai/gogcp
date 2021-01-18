// Activity 5.02

package functions

import "fmt"

type employee struct {
	ID        int
	FirstName string
	LastName  string
}

type developer struct {
	Individual employee
	HourlyRate int
	WorkWeek   [7]int
}

type weekday int

const (
	sunday weekday = iota // starts with zero
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func saladvance() {
	d := developer{Individual: employee{ID: 1, FirstName: "Nirav", LastName: "Chotai"}, HourlyRate: 10}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()
	d.loggedHours(monday, 8)
	d.loggedHours(tuesday, 10)
	d.loggedHours(wednesday, 10)
	d.loggedHours(thursday, 10)
	d.loggedHours(friday, 6)
	d.loggedHours(saturday, 8)
	d.payDetails()
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (d *developer) loggedHours(day weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *developer) hoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func (d *developer) payDay() (int, bool) {
	if d.hoursWorked() > 40 {
		hoursOver := d.hoursWorked() - 40
		overTime := hoursOver * 2
		regularPay := d.hoursWorked() * d.HourlyRate
		return regularPay + overTime, true
	}
	return d.hoursWorked() * d.HourlyRate, false
}

func (d *developer) payDetails() {
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
		fmt.Printf("\nHours worked this week: %d\n", d.hoursWorked())
		pay, overtime := d.payDay()
		fmt.Println("Pay for this week: $", pay)
		fmt.Println("Is this overtime pay: ", overtime)
		fmt.Println()
	}
}
