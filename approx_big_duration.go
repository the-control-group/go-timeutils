package timeutils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"bytes"
	"encoding/json"
)

type ApproxBigDuration time.Duration

const Day = time.Hour * 24
const Month = Day * 30
const Year = Month * 12

var negativePattern = regexp.MustCompile("^(~ )?-")
var nanosPattern = regexp.MustCompile("(\\d)+ ?(?:ns|nanos|nanosecond|nanoseconds)")
var microsPattern = regexp.MustCompile("(\\d)+ ?(?:µ|µs|ns|nanos|nanosecond|nanoseconds)")
var millisPattern = regexp.MustCompile("(\\d)+ ?(?:ms)")
var secsPattern = regexp.MustCompile("(\\d+) ?(?:s|sec|secs)")
var minsPattern = regexp.MustCompile("(\\d+) ?(?:m[^o]|m$|min|mins)")
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
	if d < 0 {
		return "-" + (-d).String()
	}
	switch {
	case time.Duration(d) < 4*Day:
		return time.Duration(d).String()
	case time.Duration(d) < 1*Month:
		var days = int64(float64(d) / float64(Day))
		var hours = (int64(d) - (days * int64(Day))) / int64(time.Hour)
		var remainder = time.Duration(int64(d) - (days * int64(Day)) - (hours * int64(time.Hour)))
		return fmt.Sprintf("%dd%dh%s", days, hours, remainder)
	case time.Duration(d) < 12*Month:
		var months = int64(float64(d) / float64(Month))
		var days = int64(int64(d)-(months*int64(Month))) / int64(Day)
		var remainder = time.Duration(int64(d) - (months * int64(Month)) - (days * int64(Day)))
		return fmt.Sprintf("%dmo%dd%s", months, days, remainder)
	default:
		var years = int64(float64(d) / float64(Year))
		var months = int64(int64(d)-(years*int64(Year))) / int64(Month)
		var remainder = time.Duration(int64(d) - (years * int64(Year)) - (months * int64(Month)))
		return fmt.Sprintf("%dy%dmo%s", years, months, remainder)
	}
}

func (d ApproxBigDuration) Pretty() string {
	if d < 0 {
		return "- " + (-d).Pretty()
	}
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
	if d < 0 {
		return strings.Replace((-d).ApproxPretty(), "~ ", "~ -", 1)
	}
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
	_d, err := ParseApproxBigDuration(data)
	if err != nil {
		return err
	}
	*d += _d
	return nil
}

// ParseApproxBigDuration parses a string for an approximate duration string
func ParseApproxBigDuration(data []byte) (ApproxBigDuration, error) {
	var d ApproxBigDuration
	if nanosPattern.Match(data) {
		v, err := strconv.Atoi(string(nanosPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * time.Nanosecond)
	}
	if microsPattern.Match(data) {
		v, err := strconv.Atoi(string(microsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * time.Microsecond)
	}
	if millisPattern.Match(data) {
		v, err := strconv.Atoi(string(millisPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * time.Millisecond)
	}
	if secsPattern.Match(data) {
		v, err := strconv.Atoi(string(secsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * time.Second)
	}
	if minsPattern.Match(data) {
		v, err := strconv.Atoi(string(minsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * time.Minute)
	}
	if hoursPattern.Match(data) {
		v, err := strconv.Atoi(string(hoursPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * time.Hour)
	}
	if daysPattern.Match(data) {
		v, err := strconv.Atoi(string(daysPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * Day)
	}
	if monthsPattern.Match(data) {
		v, err := strconv.Atoi(string(monthsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * Month)
	}
	if yearsPattern.Match(data) {
		v, err := strconv.Atoi(string(yearsPattern.FindSubmatch(data)[1]))
		if err != nil {
			return d, err
		}
		d += ApproxBigDuration(time.Duration(v) * Year)
	}
	if negativePattern.Match(data) {
		d = -1 * d
	}
	return d, nil
}

// InterfaceToApproxBigDuration handles converting various types
// Can be used as a template function
func InterfaceToApproxBigDuration(i interface{}) (ApproxBigDuration, error) {
	switch exp := i.(type) {
	case []byte:
		return ParseApproxBigDuration(exp)
	case string:
		return ParseApproxBigDuration([]byte(exp))
	// case int, int32, int64, float64:
	// 	return ApproxBigDuration(exp), nil
	case int:
		return ApproxBigDuration(exp), nil
	case int32:
		return ApproxBigDuration(exp), nil
	case int64:
		return ApproxBigDuration(exp), nil
	case float32:
		return ApproxBigDuration(exp), nil
	case float64:
		return ApproxBigDuration(exp), nil
	case json.Number:
		_i, err := exp.Int64()
		if err != nil {
			return 0, err
		}
		return ApproxBigDuration(_i), nil
	default:
		return 0, fmt.Errorf("Could not convert type %T to ApproxBigDuration", exp)
	}
}
