package main

import (
	"net/url"
	"testing"
)

var supportedPeriods = []string{"1h", "1d", "1mo"}

func TestValidator_AllOK(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"tz":     []string{"Europe/Athens"},
		"t1":     []string{"20060102T150405Z"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err != nil {
		t.Error("There should be no error")
	}

}

func TestValidator_WrongEndpoint(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"tz":     []string{"Europe/Athens"},
		"t1":     []string{"20060102T150405Z"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("wrong", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400EndpointDoesNotExist {
		t.Error("Wrong error message")
	}
}

func TestValidator_PeriodDoesNotExist(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, []string{"invalid"})
	query := url.Values{
		"tz": []string{"Europe/Athens"},
		"t1": []string{"20060102T150405Z"},
		"t2": []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400NoPeriod {
		t.Error("Wrong error message")
	}
}

func TestValidator_UnsupportedPeriod(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"2h"},
		"tz":     []string{"Europe/Athens"},
		"t1":     []string{"20060102T150405Z"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400PeriodUnsupported {
		t.Error("Wrong error message")
	}
}

func TestValidator_TimezoneDoesNotExist(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"t1":     []string{"20060102T150405Z"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400NoTimezone {
		t.Error("Wrong error message")
	}
}

func TestValidator_TimezoneInvalid(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"tz":     []string{"invalid"},
		"t1":     []string{"20060102T150405Z"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400timezoneNotIANA {
		t.Error("Wrong error message")
	}
}

func TestValidate_T1DoesNotExist(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"tz":     []string{"Europe/Athens"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400Not1 {
		t.Error("Wrong error message")
	}
}

func TestValidate_T1Invalid(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"tz":     []string{"Europe/Athens"},
		"t1":     []string{"20060102T150405Z+00"},
		"t2":     []string{"20070102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400t1Invalid {
		t.Error("Wrong error message")
	}
}

//Similarly for t2

func TestValidate_T1AfterT2(t *testing.T) {
	//arrange
	v := newValidator("right", validTimeFormat, supportedPeriods)
	query := url.Values{
		"period": []string{"1h"},
		"tz":     []string{"Europe/Athens"},
		"t1":     []string{"20060102T150405Z"},
		"t2":     []string{"20050102T150405Z"},
	}
	//act
	_, err := v.validateAndParse("right", query)

	//assert
	if err == nil {
		t.Error("There should be an error")
	}

	if err.Error() != error400t2aftert1 {
		t.Error("Wrong error message")
	}
}
