package model

import "time"

type Delegation struct {
	ID        int
	Amount    int
	Block     string
	Delegator string
	Timestamp time.Time
}
