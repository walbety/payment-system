package canonical

import "time"

type Product struct {
	Id    string       `json:"id"`
	Name  string       `json:"name"`
	Specs ProductSpecs `json:"specs"`
}

type ProductSpecs struct {
	Model      string       `json:"model"`
	Category   CategoryType `json:"category"`
	LaunchDate time.Time    `json:"launch"`
}

type CategoryType int

const (
	Unspecified      CategoryType = 0
	Consumable       CategoryType = 1
	Electrodomestics CategoryType = 2
	Utilitaries      CategoryType = 3
)
