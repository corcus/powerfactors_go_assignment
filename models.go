package main

import "time"

type periodUnit string

const (
	hours  periodUnit = "h"
	days   periodUnit = "d"
	months periodUnit = "mo"
	years  periodUnit = "y"
)

type request struct {
	period   string
	t1       time.Time
	t2       time.Time
	location *time.Location
}
