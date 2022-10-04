// Taken from https://socketloop.com/tutorials/golang-how-to-calculate-the-distance-between-two-coordinates-using-haversine-formula
package calculators

import (
	"math"

	"github.com/joeowen22/golang-todd-packers-travels/internal/models"
)

const radius = 6371 // Earth's mean radius in kilometers
const kmToMi = 0.6213711922

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Distance(origin *models.Coordinate, destination *models.Coordinate) float64 {
	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.Latitude))*
			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c

	return d * kmToMi
}
