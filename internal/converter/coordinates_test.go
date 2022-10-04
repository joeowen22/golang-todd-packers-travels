package converter

import (
	"testing"
)

func TestConvertValidScenario(t *testing.T) {
	validString := "-3.3685709, 50.8076140"
	coordinate := ConvertToCoordinate(validString)

	if coordinate.Latitude != -3.3685709 && coordinate.Longitude != 50.8076140 {
		t.Errorf("Expected: [%v], [%v] but was actually [%v] [%v]", -3.3685709, 50.8076140, coordinate.Latitude, coordinate.Longitude)
	}
}

func TestConvertInvalidScenario(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	validString := "lat, 50.8076140"
	ConvertToCoordinate(validString)
}
