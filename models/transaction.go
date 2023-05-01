package models

import "time"

type TransactionUnmarshalled struct {
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	Date      string  `json:"date"`
	Frequency string  `json:"frequency"`
	Type      string  `json:"type"`
}

type Transaction struct {
	Name      string
	Amount    float64
	Category  string
	Date      time.Time
	Frequency string
	Type      string
}

func (t *Transaction) IsPast() bool {
	return t.Date.Before(time.Now())
}

func (t *Transaction) IsRecurrent() bool {
	return t.Frequency != "Onetime"
}
