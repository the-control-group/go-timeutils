package timeutils

import (
	"fmt"
	"time"
	"regexp"
	"strconv"
	// "encoding/json"
	"bytes"
)

type ApproxBigDuration time.Duration

const Day = time.Hour * 24
const Month = Day * 30
const Year = Month * 12

var nanosPattern = regexp.MustCompile("(\\d)+ ?(?:ns|nanos|nanosecond|nanoseconds)")
var microsPattern = regexp.MustCompile("(\\d)+ ?(?:µ|µs|ns|nanos|nanosecond|nanoseconds)")
var millisPattern = regexp.MustCompile("(\\d)+ ?(?:ms)")
var secsPattern = regexp.MustCompile("(\\d+) ?(?:s|sec|secs)")
var minsPattern = regexp.MustCompile("(\\d+) ?(?:m|min|mins)")
var hoursPattern = regexp.MustCompile("(\\d+) ?(?:h|hr|hrs)")
var daysPattern = regexp.MustCompile("(\\d+) ?(?:d|day|days)")
var monthsPattern = regexp.MustCompile("(\\d+) ?(?:mo|mos)")
var yearsPattern = regexp.MustCompile("(\\d+) ?(?:y|yr|yrs)")

func (d ApproxBigDuration) ApproxString() string {
	switch {
	case time.Duration(d) < 4*Day:
		return time.Duration(d).String()
	case time.Duration(d) < 1*Month:
		var days = int64(float64(d) / float64(Day))
		var hours = (int64(d) - (days * int64(Day))) / int64(time.Hour)
		return fmt.Sprintf("~%dd%dh", days, hours)
	case time.Duration(d) < 12*Month:
		var months = int64(float64(d) / float64(Month))
		var days = int64(int64(d)-(months*int64(Month))) / int64(Day)
		return fmt.Sprintf("~%dmo%dd", months, days)
	default:
		var years = int64(float64(d) / float64(Year))
		var months = int64(int64(d)-(years*int64(Year))) / int64(Month)
		return fmt.Sprintf("~%dy%dmo", years, months)
	}
}

func (d ApproxBigDuration) String() string {
	switch {
	case time.Duration(d) < 4*Day:
		return time.Duration(d).String()
	case time.Duration(d) < 1*Month:
		var days = int64(float64(d) / float64(Day))
		var hours = (int64(d) - (days * int64(Day))) / int64(time.Hour)
		var remainder = time.Duration(int64(d) - (days * int64(Day)) - (hours * int64(time.Hour)))
		return fmt.Sprintf("%dd%dh %s", days, hours, remainder)
	case time.Duration(d) < 12*Month:
		var months = int64(float64(d) / float64(Month))
		var days = int64(int64(d)-(months*int64(Month))) / int64(Day)
		var remainder = time.Duration(int64(d) - (months * int64(Month)) - (days * int64(Day)))
		return fmt.Sprintf("%dmo%dd %s", months, days, remainder)
	default:
		var years = int64(float64(d) / float64(Year))
		var months = int64(int64(d)-(years*int64(Year))) / int64(Month)
		var remainder = time.Duration(int64(d) - (years * int64(Year)) - (months * int64(Month)))
		return fmt.Sprintf("%dy%dmo %s", years, months, remainder)
	}
}

func (d ApproxBigDuration) Pretty() string {
	switch {
	case time.Duration(d) < 4*Day:
		return time.Duration(d).String()
	case time.Duration(d) < 1*Month:
		var days = int64(float64(d) / float64(Day))
		var hours = (int64(d) - (days * int64(Day))) / int64(time.Hour)
		var remainder = time.Duration(int64(d) - (days * int64(Day)) - (hours * int64(time.Hour)))
		return fmt.Sprintf("%d days, %d hours %s", days, hours, remainder)
	case time.Duration(d) < 12*Month:
		var months = int64(float64(d) / float64(Month))
		var days = int64(int64(d)-(months*int64(Month))) / int64(Day)
		var remainder = time.Duration(int64(d) - (months * int64(Month)) - (days * int64(Day)))
		return fmt.Sprintf("%d months, %d days %s", months, days, remainder)
	default:
		var years = int64(float64(d) / float64(Year))
		var months = int64(int64(d)-(years*int64(Year))) / int64(Month)
		var remainder = time.Duration(int64(d) - (years * int64(Year)) - (months * int64(Month)))
		return fmt.Sprintf("%d years, %d months %s", years, months, remainder)
	}
}

func (d ApproxBigDuration) ApproxPretty() string {
	switch {
	case time.Duration(d) < 4*Day:
		return time.Duration(d).String()
	case time.Duration(d) < 1*Month:
		var days = int64(float64(d) / float64(Day))
		var hours = (int64(d) - (days * int64(Day))) / int64(time.Hour)
		return fmt.Sprintf("~ %d days, %d hours", days, hours)
	case time.Duration(d) < 12*Month:
		var months = int64(float64(d) / float64(Month))
		var days = int64(int64(d)-(months*int64(Month))) / int64(Day)
		return fmt.Sprintf("~ %d months, %d days", months, days)
	default:
		var years = int64(float64(d) / float64(Year))
		var months = int64(int64(d)-(years*int64(Year))) / int64(Month)
		return fmt.Sprintf("~ %d years, %d months", years, months)
	}
}

func (d *ApproxBigDuration) UnmarshalJSON(data []byte) error {
	if bytes.Compare([]byte(`null`), bytes.ToLower(data)) == 0 {
		return nil
	}
	if nanosPattern.Match(data) {
		v, err := strconv.Atoi(string(nanosPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * time.Nanosecond)
	}
	if microsPattern.Match(data) {
		v, err := strconv.Atoi(string(microsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * time.Microsecond)
	}
	if millisPattern.Match(data) {
		v, err := strconv.Atoi(string(millisPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * time.Millisecond)
	}
	if secsPattern.Match(data) {
		v, err := strconv.Atoi(string(secsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * time.Second)
	}
	if minsPattern.Match(data) {
		v, err := strconv.Atoi(string(minsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * time.Minute)
	}
	if hoursPattern.Match(data) {
		v, err := strconv.Atoi(string(hoursPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * time.Hour)
	}
	if daysPattern.Match(data) {
		v, err := strconv.Atoi(string(daysPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * Day)
	}
	if monthsPattern.Match(data) {
		v, err := strconv.Atoi(string(monthsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * Month)
	}
	if yearsPattern.Match(data) {
		v, err := strconv.Atoi(string(yearsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return err
		}
		*d += ApproxBigDuration(time.Duration(v) * Year)
	}
	return nil
}