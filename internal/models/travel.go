package models

type Travel struct {
	From               string `json:"from"`
	To                 string `json:"to"`
	PriceOfFuelPerMile string `json:"priceOfFuelPerMile"`
	TimeTaken          string `json:"timeTaken"`
}
