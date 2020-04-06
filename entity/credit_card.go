package entity

import "time"

type CreditCard struct {
	Number     string
	Expiration time.Time
}
