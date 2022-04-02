package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	//initialize new time.Time object and pass it in humanDate function
	tm := time.Date(2022, time.April, 12, 12, 0, 0, 0, time.UTC)
	hd := humanDate(tm)

	//check output from humanDate function is in format expected else use t.Errorf() function to log value and test failed
	if hd != "12 Apr 2022 at 12:00" {
		t.Errorf("want %q; got %q", "12 Apr 2022 at 12:00", hd)
	}
}
