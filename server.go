package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type httpServer struct {
	address              string
	port                 int
	validator            requestValidator
	validateAndParseFunc func()
}

func newHttpServer(address string, port int, validator requestValidator) httpServer {
	return httpServer{
		address:   address,
		port:      port,
		validator: validator,
	}
}

func (s httpServer) getAddressAndPort() string {
	return fmt.Sprintf("%s:%d", s.address, s.port)
}

func (s httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestModel, err := s.validator.validateAndParse(r.URL.Path, r.URL.Query())
	if err != nil {
		response := badRequestResponse{
			Status: "error",
			Error:  err.Error(),
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "")
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, string(responseBytes))
		return
	}

	response := handleRequest(requestModel)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(responseBytes))
}

func handleRequest(r *request) SuccessResponse {
	number, unit := parsePeriod(r.period)
	t1InLocation := r.t1.In(r.location)
	t2InLocation := r.t2.In(r.location)
	var ptList []string

	//It is assumed that t1 and t2 cannot be part of the ptlist.
	switch unit {
	case hours:
		ptList = calculateTimestampsHours(number, t1InLocation, t2InLocation)
		break
	case days:
		ptList = calculateTimestampsDays(number, t1InLocation, t2InLocation)
		break
	case months:
		ptList = calculateTimestampsMonths(number, t1InLocation, t2InLocation)
		break
	case years:
		ptList = calculateTimestampsYears(number, t1InLocation, t2InLocation)
		break
	}

	return SuccessResponse(ptList)
}

func parsePeriod(period string) (number uint8, unit periodUnit) {
	i := 0
	for ; i < len(period); i++ {
		digit := period[i]
		if digit < '0' || digit > '9' {
			break
		}
		number = number*10 + digit - '0' //there are no checks for overflows. It is assumed that supported periods are sensible values.
	}
	return number, periodUnit(period[i:]) //since the period has passed validation in an earlier stage it is safe to cast like this
}

// endOfDay returns the time.Time at the end of the day
// like this example for 2023-01-02T13:35:50
// add one day 2023-01-03T13:35:50
// subtract 13 hours, 35 minutes and 50 seconds 2023-01-03T00:00:00
func endOfDay(date time.Time) time.Time {
	subtractDuration := time.Duration(date.Hour())*time.Hour + time.Duration(date.Minute())*time.Minute + time.Duration(date.Second())*time.Second
	return date.AddDate(0, 0, 1).Add(-subtractDuration)
}

// endOfMonth returns the time.Time at the end of the month
// by using the time package convention that 2023-02-28 = 2023-03-00
// and endOfDay
func endOfMonth(date time.Time) time.Time {
	return endOfDay(date.AddDate(0, 1, -date.Day()))
}

// endOfYear returns the time.Time at the end of the year
// by using the time package convention that 2023-12-31 = 2024-01-00
func endOfYear(date time.Time) time.Time {
	return date.AddDate(1, -int(date.Month()-1), -(date.Day() - 1)).Add(time.Duration(-date.Hour()) * time.Hour)
}

func calculateTimestampsHours(step uint8, t1InLocation time.Time, t2InLocation time.Time) []string {
	ptList := make([]string, 0)
	nextTimestamp := t1InLocation.Add(time.Duration(step) * time.Hour).Truncate(1 * time.Hour) //add and truncate sets the invocation timestamp at the beginning of the hour
	for nextTimestamp.Before(t2InLocation) {
		ptList = append(ptList, nextTimestamp.UTC().Format(validTimeFormat))
		nextTimestamp = nextTimestamp.Add(time.Duration(step) * time.Hour).Truncate(1 * time.Hour)
	}
	return ptList
}

func calculateTimestampsDays(step uint8, t1InLocation time.Time, t2InLocation time.Time) []string {
	ptList := make([]string, 0)
	nextTimestamp := endOfDay(t1InLocation)
	for nextTimestamp.Before(t2InLocation) {
		ptList = append(ptList, nextTimestamp.UTC().Format(validTimeFormat))
		nextTimestamp = nextTimestamp.AddDate(0, 0, int(step))
	}
	return ptList
}

func calculateTimestampsMonths(step uint8, t1InLocation time.Time, t2InLocation time.Time) []string {
	ptList := make([]string, 0)
	nextTimestamp := endOfMonth(t1InLocation)
	for nextTimestamp.Before(t2InLocation) {
		ptList = append(ptList, nextTimestamp.UTC().Format(validTimeFormat))
		nextTimestamp = endOfMonth(nextTimestamp.AddDate(0, int(step), -1)) //remove one extra day for endOfMonth to work correctly
	}
	return ptList
}

func calculateTimestampsYears(step uint8, t1InLocation time.Time, t2InLocation time.Time) []string {
	ptList := make([]string, 0)
	nextTimestamp := endOfYear(t1InLocation).Truncate(1 * time.Hour)
	for nextTimestamp.Before(t2InLocation) {
		ptList = append(ptList, nextTimestamp.UTC().Format(validTimeFormat))
		nextTimestamp = nextTimestamp.AddDate(int(step), 0, 0)
	}
	return ptList
}
