package main

import (
	"testing"
	"time"

	"github.com/print-agent/snippet-box/internal/assert"
)

func TestHumanDate(t *testing.T) {

	testCases := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "Start of the year 2000 in UTC",
			tm:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: "01 Jan 2000 at 00:00",
		},
		{
			name: "End of the year 2099 in UTC",
			tm:   time.Date(2099, time.December, 31, 23, 59, 0, 0, time.UTC),
			want: "31 Dec 2099 at 23:59",
		},
		{
			name: "April 15, 2020 at 12:30 in UTC",
			tm:   time.Date(2020, time.April, 15, 12, 30, 0, 0, time.UTC),
			want: "15 Apr 2020 at 12:30",
		},
		{
			name: "Midnight on July 4, 1776 in UTC",
			tm:   time.Date(1776, time.July, 4, 0, 0, 0, 0, time.UTC),
			want: "04 Jul 1776 at 00:00",
		},
		{
			name: "Year 1 in UTC",
			tm:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: "01 Jan 0001 at 00:00",
		},
		{
			name: "Year 9999 in UTC",
			tm:   time.Date(9999, time.December, 31, 23, 59, 0, 0, time.UTC),
			want: "31 Dec 9999 at 23:59",
		},
		{
			name: "Leap year in UTC",
			tm:   time.Date(2000, time.February, 29, 12, 0, 0, 0, time.UTC),
			want: "29 Feb 2000 at 12:00",
		},
		{
			name: "January 1, 2000 at 00:00 in UTC-5",
			tm:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.FixedZone("UTC-5", -5*60*60)),
			want: "01 Jan 2000 at 05:00",
		},
		{
			name: "January 1, 2000 at 00:00 in UTC+5",
			tm:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.FixedZone("UTC+5", 5*60*60)),
			want: "31 Dec 1999 at 19:00",
		},
		{
			name: "Empty time (zero value)",
			tm:   time.Time{},
			want: "01 Jan 0001 at 00:00",
		},
		{
			name: "DST in America/New_York (winter)",
			tm:   time.Date(2023, time.January, 1, 12, 0, 0, 0, time.FixedZone("EST", -5*60*60)),
			want: "01 Jan 2023 at 17:00",
		},
		{
			name: "DST in America/New_York (summer)",
			tm:   time.Date(2023, time.July, 1, 12, 0, 0, 0, time.FixedZone("EDT", -4*60*60)),
			want: "01 Jul 2023 at 16:00",
		},
		{
			name: "Time zone with non-integer offset (UTC+5:30)",
			tm:   time.Date(2023, time.January, 1, 12, 0, 0, 0, time.FixedZone("UTC+5:30", (5*60+30)*60)),
			want: "01 Jan 2023 at 06:30",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			hd := humanDate(tC.tm)

			assert.Equal(t, hd, tC.want)
		})
	}
}
