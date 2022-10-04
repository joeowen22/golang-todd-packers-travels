package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joeowen22/golang-todd-packers-travels/internal/models"
	"github.com/joeowen22/golang-todd-packers-travels/internal/service"
)

const speedLimit float64 = 70.0

func main() {
	data, err := os.ReadFile("./data/todd_packer.json")
	check(err)

	var toddPacker models.Profile
	json.Unmarshal(data, &toddPacker)

	var summaries []models.Summary

	for _, travel := range toddPacker.Travels {
		summaries = append(summaries, service.GetSummary(&travel))
	}

	fmt.Printf("- Total miles travelled: %.2f miles at an average of %.2f  miles an hour\n", totalMiles(summaries), averageSpeed(summaries))
	fmt.Printf("- Average price of fuel per mile (total): £%.2f\n", averagePricePerMile(summaries))
	validExpenses := getValidExpenses(summaries)
	invalidExpenses := getInvalidExpenses(summaries)
	fmt.Printf("- Speed limit broken: %v times - £%.2f reimbursement lost due to speeding\n", len(invalidExpenses), totalCost(invalidExpenses))
	mostExpensive := mostExpensive(summaries)
	fmt.Printf("- Most expensive path: From %v, %v to %v, %v (%.2f miles), costing £%.2f amount (optional)\n", mostExpensive.Journey.From.Latitude, mostExpensive.Journey.From.Longitude, mostExpensive.Journey.To.Latitude, mostExpensive.Journey.To.Longitude, mostExpensive.Distance, mostExpensive.Cost)
	cheapestJourney := cheapestJourney(summaries)
	fmt.Printf("- Cheapest path: From %v, %v to %v, %v (%.2f miles), costing £%.2f amount (optional)\n", cheapestJourney.Journey.From.Latitude, cheapestJourney.Journey.From.Longitude, cheapestJourney.Journey.To.Latitude, cheapestJourney.Journey.To.Longitude, cheapestJourney.Distance, cheapestJourney.Cost)

	fmt.Printf("\nExpense Report (%v Total: £%.2f)\n", len(validExpenses), totalCost(validExpenses))
	for i, summary := range validExpenses {
		fmt.Printf("%v) Travelled from %v, %v to %v, %v doing %.2f miles an hour (miles: %.2f), expense: £%.2f\n", i+1, summary.Journey.From.Latitude, summary.Journey.From.Longitude, summary.Journey.To.Latitude, summary.Journey.To.Longitude, summary.Speed, summary.Distance, summary.Cost)
	}

	fmt.Printf("\nSpeed Limit Violations (%v Total - %v Not Paid)\n", len(invalidExpenses), len(invalidExpenses))
	for i, summary := range invalidExpenses {
		fmt.Printf("%v) Travelled from %v, %v to %v, %v doing %.2f miles an hour (miles: %.2f), %.2f over the speed limit\n", i+1, summary.Journey.From.Latitude, summary.Journey.From.Longitude, summary.Journey.To.Latitude, summary.Journey.To.Longitude, summary.Speed, summary.Distance, summary.Speed-speedLimit)
	}
}

func totalMiles(summaries []models.Summary) float64 {
	var total float64 = 0
	for _, summary := range summaries {
		total += summary.Distance
	}
	return total
}

func totalCost(summaries []models.Summary) float64 {
	var totalCost float64 = 0
	for _, summary := range summaries {
		totalCost += summary.Cost
	}
	return totalCost
}

func averageSpeed(summaries []models.Summary) float64 {
	var totalSpeed float64 = 0
	for _, summary := range summaries {
		totalSpeed += summary.Speed
	}
	return totalSpeed / float64(len(summaries))
}

func averagePricePerMile(summaries []models.Summary) float64 {
	var total float64 = 0
	for _, summary := range summaries {
		total += summary.Cost
	}
	return total / totalMiles(summaries)
}

func mostExpensive(summaries []models.Summary) models.Summary {
	var mostExpensive models.Summary
	for _, summary := range summaries {
		if summary.Cost > mostExpensive.Cost {
			mostExpensive = summary
		}
	}
	return mostExpensive
}

func cheapestJourney(summaries []models.Summary) models.Summary {
	var cheapestJourney models.Summary = summaries[0]
	for _, summary := range summaries {
		if summary.Cost < cheapestJourney.Cost {
			cheapestJourney = summary
		}
	}
	return cheapestJourney
}

func getValidExpenses(summaries []models.Summary) []models.Summary {
	var validSummaries []models.Summary
	for _, summary := range summaries {
		if summary.Speed <= speedLimit {
			validSummaries = append(validSummaries, summary)
		}
	}
	return validSummaries
}

func getInvalidExpenses(summaries []models.Summary) []models.Summary {
	var invalidSummaries []models.Summary
	for _, summary := range summaries {
		if summary.Speed > speedLimit {
			invalidSummaries = append(invalidSummaries, summary)
		}
	}
	return invalidSummaries
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
