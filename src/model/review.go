package model

import (
	"time"
)

// Review is the review of wine
type Review struct {
	Score   int
	Color   string
	Aroma   string
	Taste   string
	Reviewd time.Time
	Created time.Time
	Updated time.Time
	Account string
	WineID  int
}
