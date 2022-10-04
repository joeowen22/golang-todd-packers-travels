package service

import (
	"strconv"

	"github.com/joeowen22/golang-todd-packers-travels/internal/calculators"
	"github.com/joeowen22/golang-todd-packers-travels/internal/converter"
	"github.com/joeowen22/golang-todd-packers-travels/internal/models"
)

func GetSummary(travel *models.Travel) models.Summary {
	from := converter.ConvertToCoordinate(travel.From)
	to := converter.ConvertToCoordinate(travel.To)

	distance := calculators.Distance(&from, &to)

	speed := calculators.Speed(distance, converter.ConvertToHour(travel.TimeTaken))

	price, err := strconv.ParseFloat(travel.PriceOfFuelPerMile, 64)
	if err != nil {
		panic("Can't convert cost")
	}

	cost := distance * price

	journey := &models.Journey{
		From: from,
		To:   to,
	}
	summary := new(models.Summary)
	summary.Journey = *journey
	summary.Speed = speed
	summary.Cost = cost
	summary.Distance = distance
	return *summary
}
