// Activity 5.01

package main

import "fmt"

type Employee struct {
	ID        int
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
	Sunday Weekday = iota // starts at zero
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total
}

func main() {
	d := Developer{Individual: Employee{ID: 1, FirstName: "Nirav", LastName: "Chotai"}, HourlyRate: 10}
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday: ", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", d.WorkWeek[Tuesday])
	fmt.Printf("Hours worked for this week: %d", d.HoursWorked())
}
