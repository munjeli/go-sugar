package gosugar

import (
	"fmt"
	"regexp"
	"strconv"
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
func (d DateDuration) Truncate(s string) DateDuration {
	precedenceMap := map[string]int{
		"y": 5,
		"M": 4,
		"d": 3,
		"h": 2,
		"m": 1,
		"s": 0,
	}
	precedence, found := precedenceMap[s]
	if !found {
		return d
	}
	// This is absurdly redundant, but I'm still thinking about the data structure
	// of a DateDuration, and how to map the string values for the query vs. the
	// properties of the type.
	dd := DateDuration{
		Years:   setDurationPropByPrecedence(d.Years, 5, precedence),
		Months:  setDurationPropByPrecedence(d.Months, 4, precedence),
		Days:    setDurationPropByPrecedence(d.Days, 3, precedence),
		Hours:   setDurationPropByPrecedence(d.Hours, 2, precedence),
		Minutes: setDurationPropByPrecedence(d.Minutes, 1, precedence),
		Seconds: setDurationPropByPrecedence(d.Seconds, 0, precedence),
	}
	return dd
}

func setDurationPropByPrecedence(d int, stringForProp, precedence int) int {
	if stringForProp > precedence {
		return d
	}
	return 0
}

// ToTime Danger Danger! This might not math as expected.
// But it's still sometimes useful you know. Skip the
// months for accuracy, and use days instead.
func (d DateDuration) ToTime() time.Time {
	month := time.Month(d.Months)
	date := time.Date(d.Years, month, d.Days, d.Hours, d.Minutes, d.Seconds, 0, time.UTC)
	return date
}

// ParseDateDuration returns a DateDuration type for a
// properly formatted string.
func ParseDateDuration(s string) (d DateDuration) {
	// dels is the map of delimiters for parsing the duration string, note
	// month has to be CAP'd!
	dels := map[string]int{"y": 0, "M": 0, "d": 0, "h": 0, "m": 0, "s": 0}
	for letter := range dels {
		pat := `\d+` + letter
		var re = regexp.MustCompile(pat)
		cycle := re.FindAllString(s, 1)
		if len(cycle) != 0 {
			num := regexp.MustCompile(`\d+`).FindAllString(cycle[0], 1)
			if len(num) != 0 {
				n, _ := strconv.Atoi(num[0])
				dels[letter] = n
			}
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
