package main

import "fmt"

type employee interface {
	salary() int
}

type permananet struct {
	empID    int
	basicPay int
	pf       int
}

type contract struct {
	empID    int
	basicPay int
}

func (p permananet) salary() int {
	return p.basicPay + p.pf
}

func (c contract) salary() int {
	return c.basicPay
}

func totalExpense(e []employee) {
	expense := 0
	for _, emp := range e {
		expense += emp.salary()
	}
	fmt.Printf("Total expense per month %d", expense)
}

func main() {
	pemp1 := permananet{1, 5000, 250}
	pemp2 := permananet{1, 5000, 250}
	pemp3 := permananet{1, 5000, 250}
	cemp1 := contract{4, 3500}
	cemp2 := contract{5, 2500}
	employees := []employee{pemp1, pemp2, pemp3, cemp1, cemp2}
	totalExpense(employees)
}
