package v1

import "fmt"

type employee struct {
	firstName string
	lastName string
	totalLeaves int
	leavesToken int
}

func New(firstName string, lastName string, totalLeaves int, leavesToken int) employee {
	e := employee{
		firstName: firstName,
		lastName:lastName,
		totalLeaves: totalLeaves,
		leavesToken:leavesToken,
	}

	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leveas remaining.", e.firstName, e.lastName, (e.totalLeaves - e.leavesToken))
}