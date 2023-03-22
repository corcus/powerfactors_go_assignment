package main

import (
	"reflect"
	"testing"
	"time"
)

func TestPeriodParsing_1y(t *testing.T) {
	//arrange
	p := "1y"

	//act
	n, u := parsePeriod(p)

	//assert
	if n != 1 || u != "y" {
		t.Errorf("period %s should be parsed to 1 and y. Instead it was %d and %s", p, n, u)
	}
}

func TestPeriodParsing_15mo(t *testing.T) {
	//arrange
	p := "15mo"

	//act
	n, u := parsePeriod(p)

	//assert
	if n != 15 || u != "mo" {
		t.Errorf("period %s should be parsed to 15 and mo. Instead it was %d and %s", p, n, u)
	}
}

func TestEndOfDay(t *testing.T) {
	//arrange
	date := time.Date(2023, 02, 03, 15, 12, 45, 0, time.UTC)
	expectedDate := time.Date(2023, 02, 04, 00, 00, 00, 0, time.UTC)
	//act

	endOfDayDate := endOfDay(date)

	//assert
	if endOfDayDate.Equal(expectedDate) == false {
		t.Errorf("end of day date : %v should have been equal to expected date : %v", endOfDayDate, expectedDate)
	}
}

//TODO : similar tests for EndOfMonth and EndOfYear

// vanilla test
func TestCalculateForDays_EuropeAthens_1d(t *testing.T) {
	//arrange
	expectedPtlist := []string{
		"20211010T210000Z",
		"20211011T210000Z",
		"20211012T210000Z",
		"20211013T210000Z",
		"20211014T210000Z",
		"20211015T210000Z",
		"20211016T210000Z",
		"20211017T210000Z",
		"20211018T210000Z",
		"20211019T210000Z",
		"20211020T210000Z",
		"20211021T210000Z",
		"20211022T210000Z",
		"20211023T210000Z",
		"20211024T210000Z",
		"20211025T210000Z",
		"20211026T210000Z",
		"20211027T210000Z",
		"20211028T210000Z",
		"20211029T210000Z",
		"20211030T210000Z",
		"20211031T220000Z",
		"20211101T220000Z",
		"20211102T220000Z",
		"20211103T220000Z",
		"20211104T220000Z",
		"20211105T220000Z",
		"20211106T220000Z",
		"20211107T220000Z",
		"20211108T220000Z",
		"20211109T220000Z",
		"20211110T220000Z",
		"20211111T220000Z",
		"20211112T220000Z",
		"20211113T220000Z",
		"20211114T220000Z",
	}
	loc, err := time.LoadLocation("Europe/Athens")
	if err != nil {
		t.Error(err)
	}
	t1 := time.Date(2021, 10, 10, 20, 46, 03, 0, loc)
	t2 := time.Date(2021, 11, 15, 12, 34, 56, 0, loc)

	//act

	ptlist := calculateTimestampsDays(1, t1, t2)

	//assert
	if reflect.DeepEqual(expectedPtlist, ptlist) == false {
		t.Errorf("Expected list %v \n and calculated list %v \n do not match", expectedPtlist, ptlist)
	}
}

// test no results edge case
func TestCalculateForDays_EuropeAthens_NoResults(t *testing.T) {
	//arrange
	expectedPtlist := make([]string, 0)
	loc, err := time.LoadLocation("Europe/Athens")
	if err != nil {
		t.Error(err)
	}

	//No days between t1 and t2
	t1 := time.Date(2021, 10, 10, 10, 46, 03, 0, loc)
	t2 := time.Date(2021, 10, 10, 12, 34, 56, 0, loc)

	//act

	ptlist := calculateTimestampsDays(1, t1, t2)

	//assert
	if reflect.DeepEqual(expectedPtlist, ptlist) == false {
		t.Error("Expected and calculated list do not match")
	}
}

// test different timezone
func TestCalculateForDays_AmericaChicago_1d(t *testing.T) {
	//arrange
	expectedPtlist := []string{
		"20211011T050000Z", "20211012T050000Z", "20211013T050000Z", "20211014T050000Z", "20211015T050000Z",
		"20211016T050000Z", "20211017T050000Z", "20211018T050000Z", "20211019T050000Z", "20211020T050000Z",
		"20211021T050000Z", "20211022T050000Z", "20211023T050000Z", "20211024T050000Z", "20211025T050000Z",
		"20211026T050000Z", "20211027T050000Z", "20211028T050000Z", "20211029T050000Z", "20211030T050000Z",
		"20211031T050000Z", "20211101T050000Z", "20211102T050000Z", "20211103T050000Z", "20211104T050000Z",
		"20211105T050000Z", "20211106T050000Z", "20211107T050000Z", "20211108T060000Z", "20211109T060000Z",
		"20211110T060000Z", "20211111T060000Z", "20211112T060000Z", "20211113T060000Z", "20211114T060000Z",
		"20211115T060000Z",
	}
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		t.Error(err)
	}
	t1 := time.Date(2021, 10, 10, 20, 46, 03, 0, loc)
	t2 := time.Date(2021, 11, 15, 12, 34, 56, 0, loc)

	//act

	ptlist := calculateTimestampsDays(1, t1, t2)

	//assert
	if reflect.DeepEqual(expectedPtlist, ptlist) == false {
		t.Errorf("Expected list %v \n and calculated list %v \n do not match", expectedPtlist, ptlist)
	}
}

// test different interval
func TestCalculateForDays_EuropeAthens_2d(t *testing.T) {
	//arrange
	expectedPtlist := []string{
		"20211010T210000Z",
		"20211012T210000Z",
		"20211014T210000Z",
		"20211016T210000Z",
		"20211018T210000Z",
		"20211020T210000Z",
		"20211022T210000Z",
		"20211024T210000Z",
		"20211026T210000Z",
		"20211028T210000Z",
		"20211030T210000Z",
		"20211101T220000Z",
		"20211103T220000Z",
		"20211105T220000Z",
		"20211107T220000Z",
		"20211109T220000Z",
		"20211111T220000Z",
		"20211113T220000Z",
	}
	loc, err := time.LoadLocation("Europe/Athens")
	if err != nil {
		t.Error(err)
	}
	t1 := time.Date(2021, 10, 10, 20, 46, 03, 0, loc)
	t2 := time.Date(2021, 11, 15, 12, 34, 56, 0, loc)

	//act

	ptlist := calculateTimestampsDays(2, t1, t2)

	//assert
	if reflect.DeepEqual(expectedPtlist, ptlist) == false {
		t.Errorf("Expected list %v \n and calculated list %v \n do not match", expectedPtlist, ptlist)
	}
}

//TODO: similar tests for all different calculations
