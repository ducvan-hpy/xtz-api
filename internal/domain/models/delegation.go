package models

import "time"

type Delegation struct {
	Amount    int
	Block     string
	Delegator string
	Timestamp time.Time
}
