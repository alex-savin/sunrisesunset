package sunrisesunset

import (
	"testing"
	"time"
)

func TestGetSunriseSunset(t *testing.T) {

	date := time.Now()

	// Test invalid parameters

	// Table tests
	var invalidParameters = []struct {
		latitude      float64
		longitude     float64
		utcOffset     float64
		date          time.Time
		expectedError string
	}{
		{-95.0, -46.704082, -3.0, date, "Latitude invalid"},
		{100.0, -46.704082, -3.0, date, "Latitude invalid"},
		{-23.545570, -185.0, -3.0, date, "Longitude invalid"},
		{-23.545570, 190.0, -3.0, date, "Longitude invalid"},
		{-23.545570, -46.704082, -15.0, date, "UTC offset invalid"},
		{-23.545570, -46.704082, 18.0, date, "UTC offset invalid"},
		{-23.545570, -46.704082, -3.0, time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC), "Date invalid"},
		{-23.545570, -46.704082, -3.0, time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC), "Date invalid"},
	}

	// Test with all values in the table
	for _, pair := range invalidParameters {
		_, _, err := GetSunriseSunset(pair.latitude, pair.longitude, pair.utcOffset, pair.date)
		if err == nil {
			t.Error(
				"Expect an error",
			)
		}

		p := Parameters{
			Latitude:  pair.latitude,
			Longitude: pair.longitude,
			UtcOffset: pair.utcOffset,
			Date:      pair.date,
		}

		_, _, err = p.GetSunriseSunset()
		if err == nil {
			t.Error(
				"Expect an error",
			)
		}
	}

	// Test with valid values
	tz, _ := time.LoadLocation("Local")
	date = time.Date(int(2019), time.Month(9), int(16), int(0), int(0), int(0), int(0), tz)
	testSunrise := time.Date(int(2019), time.Month(9), int(16), int(02), int(02), int(53), int(0), tz)
	testSunset := time.Date(int(2019), time.Month(9), int(16), int(14), int(00), int(56), int(0), tz)

	// Table tests
	var tTests = []struct {
		latitude  float64
		longitude float64
		utcOffset float64
		date      time.Time
		sunrise   time.Time
		sunset    time.Time
	}{
		{-23.545570, -46.704082, -7.0, date, testSunrise, testSunset}, // Sao Paulo - Brazil-ish
	}

	// Test with all values in the table
	for _, pair := range tTests {
		sunrise, sunset, err := GetSunriseSunset(pair.latitude, pair.longitude, pair.utcOffset, pair.date)

		if err != nil {
			t.Error(
				"Expect: nil",
				"Received: ", err,
			)
		}
		if !sunrise.Equal(pair.sunrise) {
			t.Error(
				"Expected: ", pair.sunrise,
				"Received: ", sunrise,
			)
		}
		if !sunset.Equal(pair.sunset) {
			t.Error(
				"Expected: ", pair.sunset,
				"Received: ", sunset,
			)
		}

		p := Parameters{
			Latitude:  pair.latitude,
			Longitude: pair.longitude,
			UtcOffset: pair.utcOffset,
			Date:      pair.date,
		}

		sunrise, sunset, err = p.GetSunriseSunset()

		if err != nil {
			t.Error(
				"Expect: nil",
				"Received: ", err,
			)
		}
		if !sunrise.Equal(pair.sunrise) {
			t.Error(
				"Expected: ", pair.sunrise,
				"Received: ", sunrise,
			)
		}
		if !sunset.Equal(pair.sunset) {
			t.Error(
				"Expected: ", pair.sunset,
				"Received: ", sunset,
			)
		}
	}
}

func TestDifferentLength(t *testing.T) {
	var slice1 []float64
	var slice2 []float64

	slice1 = append(slice1, 16.0)
	slice2 = append(slice2, 32.0)
	slice2 = append(slice2, 64.0)

	// Table tests
	var tTests = []struct {
		result []float64
	}{
		{calcSunEqCtr(slice1, slice2)},
		{calcSunTrueLong(slice1, slice2)},
		{calcSunAppLong(slice1, slice2)},
		{calcObliqCorr(slice1, slice2)},
		{calcSunDeclination(slice1, slice2)},
		{calcEquationOfTime(slice1, slice2, slice2, slice2)},
		{calcEquationOfTime(slice2, slice2, slice2, slice1)},
	}

	// Test with all values in the table
	for _, pair := range tTests {
		if len(pair.result) != 0 {
			t.Error("Expected: length == 0")
		}
	}
}

func TestMinIndex(t *testing.T) {
	var slice []float64

	result := minIndex(slice)
	if result != -1 {
		t.Error("Expected: minIndex == -1")
	}

	slice = append(slice, 2.10)
	slice = append(slice, 1.00)
	slice = append(slice, 0.99)
	slice = append(slice, 1.50)

	result = minIndex(slice)
	if result != 2 {
		t.Error("Expected: minIndex == 2")
	}
}

func TestRound(t *testing.T) {
	// Table tests
	var tTests = []struct {
		parameter float64
		result    int
	}{
		{-0.002, 0},
		{-0.510, -1},
		{-4.290, -4},
		{0.490, 0},
		{0.510, 1},
		{10.280, 10},
	}

	// Test with all values in the table
	for _, pair := range tTests {
		result := round(pair.parameter)
		if result != pair.result {
			t.Error(
				"Expected: ", pair.result,
				"Received: ", result,
			)
		}
	}
}
