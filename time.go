package gosugar

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// DateDuration I know all the reasons not to
// do this because I know how unreliable a Month is,
// but it's still useful if not for rotational math.
// Use with caution.
type DateDuration struct {
	Years   int `json:"years,omitempty"`
	Months  int `json:"months,omitempty"`
	Days    int `json:"days,omitempty"`
	Hours   int `json:"hours,omitempty"`
	Minutes int `json:"minutes,omitempty"`
	Seconds int `json:"seconds,omitempty"`
}

// ToString converts a DateDuration to the canonical query format.
func (d DateDuration) String() string {
	str := fmt.Sprintf("%dy%dM%dd%dh%dm%ds", d.Years, d.Months, d.Days, d.Hours, d.Minutes, d.Seconds)
	return str
}

// Truncate add the char for the field you want to round;
// All fields on the right will be set to zero.
// Valid characters are `yMdhms`.
func (d DateDuration) Truncate (s string) DateDuration {
	return DateDuration{}
}

// ToTime Danger Danger! This might not math as expected.
// But it's still sometimes useful you know. The math
// as expected is also an issue with the Go time lib in
// places, it's almost like time is relative!
func (d DateDuration) ToTime () time.Time {
	return time.Time{}
}

// ParseDateDuration returns a DateDuration type for a
// properly formatted string.
func ParseDateDuration(s string) (d DateDuration) {
	// dels is the map of delimiters for parsing the duration string, note
	// month has to be CAP'd!
	dels := map[string]int{"y": 0, "M": 0, "d": 0, "h": 0, "m": 0, "s": 0}
	for k := range dels {
		strSlc := strings.Split(s, k)
		if len(strSlc) > 1 {
			// This is essentially a trim left for whatever is
			// on the left hand side of the string... probably a better way.
			re := regexp.MustCompile(`[a-z]`)
			nums := re.Split(strSlc[0], -1)
			num, _ := strconv.Atoi(nums[len(nums)-1])
			dels[k] = num
		}
	}
	dd := DateDuration{
		Years:   dels["y"],
		Months:  dels["M"],
		Days:    dels["d"],
		Hours:   dels["h"],
		Minutes: dels["m"],
		Seconds: dels["s"],
	}
	return dd
}

// TimeSubDuration Go core has time.Add(Duration), time.Sub(time.Time),
//but no time.Sub(Duration)
func TimeSubDuration(t time.Time, d time.Duration) time.Time {
	return t.Add(-d)
}

// TimeAddDateDuration adds DateDuration to a time.Time.
func TimeAddDateDuration(t time.Time, d DateDuration) time.Time {
	return time.Time{}
}

// TimeSubDateDuration subtracts a DateDuration from a time and
// returns a time.
func TimeSubDateDuration(t time.Time, d DateDuration) time.Time {
	return time.Time{}
}

// SumDurations takes a slice of durations, adds them to zero time,
// then subtracts the resulting time from zero time.
func SumDurations(ds []time.Duration) time.Duration {
	startTime := time.Time{}
	for _, d := range ds {
		startTime = startTime.Add(d)
	}
	zeroTime := time.Time{}
	return startTime.Sub(zeroTime)
}