package main

import "oop/employee"

func main() {
	e := employee.Employee{
		FirstName:   "Sri Harsha",
		LastName:    "Kappala",
		TotalLeaves: 25,
		LeavesTaken: 6,
	}
	e.LeavesRemaining()
}
