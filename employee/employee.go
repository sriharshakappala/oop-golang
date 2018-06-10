package employee

import "fmt"

// Employee has FirstName, LastName, TotalLeaves, LeavesTaken
type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// LeavesRemaining is a derived method to calculate number of leaves remaining
func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
