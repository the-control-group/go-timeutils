package timeutils

import (
	"fmt"
	"time"
)

// ParseAnyLayouts is a collection of formats ParseAny() will attempt to use
var ParseAnyLayouts = []string{
	ISO8601_DATE,              // "2006-01-02"
	ISO8601_DATETIME,          // "2006-01-02T15:04:05Z"
	ISO8601_DATETIME_MILLI,    // "2006-01-02T15:04:05.000Z"
	ISO8601_DATETIME_TZ,       // "2006-01-02T15:04:05-07:00"
	ISO8601_DATETIME_MILLI_TZ, // "2006-01-02T15:04:05.000-07:00"
	time.ANSIC,                // "Mon Jan _2 15:04:05 2006"
	time.UnixDate,             // "Mon Jan _2 15:04:05 MST 2006"
	time.RubyDate,             // "Mon Jan 02 15:04:05 -0700 2006"
	time.RFC822,               // "02 Jan 06 15:04 MST"
	time.RFC822Z,              // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	time.RFC850,               // "Monday, 02-Jan-06 15:04:05 MST"
	time.RFC1123,              // "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC1123Z,             // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	time.RFC3339,              // "2006-01-02T15:04:05Z07:00"
	time.RFC3339Nano,          // "2006-01-02T15:04:05.999999999Z07:00"
	time.Stamp,                // "Jan _2 15:04:05"
	time.StampMilli,           // "Jan _2 15:04:05.000"
	time.StampMicro,           // "Jan _2 15:04:05.000000"
	time.StampNano,            // "Jan _2 15:04:05.000000000"
	"2006-01-02 15:04:05Z",
}

// ParseAny attempts to parse the given string in a collection of known formats
// You can add to the list of supported formats by appending ParseAnyLayouts
func ParseAny(timeString string) (t time.Time, err error) {
	for _, layout := range ParseAnyLayouts {
		t, err = time.Parse(layout, timeString)
		if err == nil {
			return t, nil
		}
	}
	return t, fmt.Errorf("No matching layout for time string")
}
