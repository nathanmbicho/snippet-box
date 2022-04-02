package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	//slice of anonymous structs containing the test case name, input to humanDate and result
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, time.April, 12, 12, 0, 0, 0, time.UTC),
			want: "12 Apr 2022 at 12:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "EAT",
			tm:   time.Date(2022, time.April, 12, 12, 0, 0, 0, time.FixedZone("EAT", 3*60*60)),
			want: "12 Apr 2022 at 09:00",
		},
	}

	//loop over the test cases
	for _, tt := range tests {
		//use t.Run to run the subtest of each test case
		t.Run(tt.name, func(t *testing.T) {
			//check output from humanDate function is in format expected else use t.Errorf() function to log value and test failed
			hd := humanDate(tt.tm)
			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}
