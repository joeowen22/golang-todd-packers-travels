package converter

import (
	"strconv"
	"strings"
)

func ConvertToHour(time string) float64 {
	timeSplit := strings.Split(time, ":")
	if len(timeSplit) != 3 {
		panic("Invalid format")
	}

	hour, err := strconv.ParseFloat(timeSplit[0], 64)
	if err != nil {
		panic("Unable to parse hour")
	}
	minute, err := strconv.ParseFloat(timeSplit[1], 64)
	if err != nil {
		panic("Unable to parse minute")
	}
	seconds, err := strconv.ParseFloat(timeSplit[2], 64)
	if err != nil {
		panic("Unable to parse seconds")
	}

	return hour + (minute / 60) + (seconds / 60 / 60)
}
