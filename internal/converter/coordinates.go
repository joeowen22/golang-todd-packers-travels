package converter

import (
	"strconv"
	"strings"

	"github.com/joeowen22/golang-todd-packers-travels/internal/models"
)

func ConvertToCoordinate(rawCoordinate string) models.Coordinate {
	splitCoordinate := strings.Split(rawCoordinate, ", ")
	if len(splitCoordinate) != 2 {
		panic("Coordinate not in correct format")
	}

	latitude, err := strconv.ParseFloat(splitCoordinate[0], 64)
	if err != nil {
		panic("Unable to parse latitude")
	}

	longitude, err := strconv.ParseFloat(splitCoordinate[1], 64)
	if err != nil {
		panic("Unable to parse latitude")
	}

	coordinate := new(models.Coordinate)
	coordinate.Latitude = latitude
	coordinate.Longitude = longitude

	return *coordinate
}
