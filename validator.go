package main

import (
	"errors"
	"net/url"
	"time"
)

type requestValidator struct {
	validEndpoint    string
	validTimeFormat  string
	supportedPeriods []string
}

func newValidator(validEndpoint string, validTimeFormat string, supportedPeriods []string) requestValidator {
	return requestValidator{
		validEndpoint:    validEndpoint,
		supportedPeriods: supportedPeriods,
		validTimeFormat:  validTimeFormat,
	}
}

func (v requestValidator) validateAndParse(path string, query url.Values) (*request, error) {
	if path != v.validEndpoint {
		return nil, errors.New(error400EndpointDoesNotExist)
	}

	if query.Has("period") == false {
		return nil, errors.New(error400NoPeriod)
	}
	period := query.Get("period")
	periodSupported := false
	for _, supportedPeriod := range v.supportedPeriods {
		if supportedPeriod == period {
			periodSupported = true
			break
		}
	}
	if periodSupported == false {
		return nil, errors.New(error400PeriodUnsupported)
	}

	if query.Has("tz") == false {
		return nil, errors.New(error400NoTimezone)
	}
	location, err := time.LoadLocation(query.Get("tz"))
	if err != nil {
		return nil, errors.New(error400timezoneNotIANA)
	}

	if query.Has("t1") == false {
		return nil, errors.New(error400Not1)
	}
	t1, err := time.Parse(v.validTimeFormat, query.Get("t1"))
	if err != nil {
		return nil, errors.New(error400t1Invalid)
	}

	if query.Has("t2") == false {
		return nil, errors.New(error400Not2)
	}
	t2, err := time.Parse(v.validTimeFormat, query.Get("t2"))
	if err != nil {
		return nil, errors.New(error400t2Invalid)
	}

	if t2.After(t1) == false {
		return nil, errors.New(error400t2aftert1)
	}

	return &request{
		period:   period,
		location: location,
		t1:       t1,
		t2:       t2,
	}, nil
}
