package main

import "oop/employee"

func main() {
	e := employee.New("Sri Harsha", "Kappala", 25, 6)
	e.LeavesRemaining()
}
