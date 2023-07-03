package model

import "time"

type Delegation struct {
	ID        int
	Amount    int
	Block     string
	Delegator string
	Timestamp time.Time
}

// DelegationsByYear is used to save delegations efficiently.
type DelegationsByYearToSave map[int][]Delegation
