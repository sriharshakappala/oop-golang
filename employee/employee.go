package employee

import "fmt"

// Employee has FirstName, LastName, TotalLeaves, LeavesTaken
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

// New is a constructor method for initializing employee object
func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeaves, leavesTaken}
	return e
}

// LeavesRemaining is a derived method to calculate number of leaves remaining
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
