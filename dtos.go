package main

// response dtos
type badRequestResponse struct {
	Status string
	Error  string
}

type SuccessResponse []string

const (
	error400EndpointDoesNotExist = "The requested endpoint does not exist"
	error400PeriodUnsupported    = "The requested period is not supported"
	error400NoPeriod             = "No period value provided"
	error400t1Invalid            = "The t1 time value is invalid"
	error400Not1                 = "No t1 value provided"
	error400t2Invalid            = "The t2 time value is invalid"
	error400Not2                 = "No t2 value provided"
	error400timezoneNotIANA      = "The timezone provided is not an IANA timezone"
	error400NoTimezone           = "No timezone value provided"
	error400t2aftert1            = "t2 must be after t1"
)
