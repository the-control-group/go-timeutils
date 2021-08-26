package timeutils

import (
	"testing"
)

var TestFormats = map[string]string{
	"ISO8601_DATE":              "2006-01-02",
	"ISO8601_DATETIME":          "2006-01-02T15:04:05Z",
	"ISO8601_DATETIME_MILLI":    "2006-01-02T15:04:05.000Z",
	"ISO8601_DATETIME_TZ":       "2006-01-02T15:04:05-07:00",
	"ISO8601_DATETIME_MILLI_TZ": "2006-01-02T15:04:05.000-07:00",
	"time.ANSIC":                "Mon Jan 2 15:04:05 2006",
	"time.UnixDate":             "Mon Jan 2 15:04:05 MST 2006",
	"time.RubyDate":             "Mon Jan 02 15:04:05 -0700 2006",
	"time.RFC822":               "02 Jan 06 15:04 MST",
	"time.RFC822Z":              "02 Jan 06 15:04 -0700", // RFC822 with numeric zone
	"time.RFC850":               "Monday, 02-Jan-06 15:04:05 MST",
	"time.RFC1123":              "Mon, 02 Jan 2006 15:04:05 MST",
	"time.RFC1123Z":             "Mon, 02 Jan 2006 15:04:05 -0700", // RFC1123 with numeric zone
	"time.RFC3339":              "2006-01-02T15:04:05-07:00",
	"time.RFC3339Z":             "2006-01-02T15:04:05Z",
	"time.RFC3339Nano":          "2006-01-02T15:04:05.999999999-07:00",
	"time.RFC3339NanoZ":         "2006-01-02T15:04:05.999999999Z",
	"time.Stamp":                "Jan 2 15:04:05",
	"time.StampMilli":           "Jan 2 15:04:05.000",
	"time.StampMicro":           "Jan 2 15:04:05.000000",
	"time.StampNano":            "Jan 2 15:04:05.000000000",
	"$.event.created 01":        "2020-12-01T00:00:21+00:00",
	"$.event.created 02":        "2020-12-01T00:19:51.481Z",
	"$.event.created 03":        "2020-12-01 23:05:36Z",
	"$.event.created 04":        "2020-12-01 23:05:36.000",
}

func TestParseAny(t *testing.T) {
	for name, format := range TestFormats {
		p, err := ParseAny(format)
		if err != nil {
			t.Error(err, name, format)
		} else {
			t.Log(p.Unix(), name, format)
		}
	}
}

func TestParseAnyMaybe(t *testing.T) {
	for name, format := range TestFormats {
		p := ParseAnyMaybe(format)
		if p == nil {
			t.Error(name, format)
		} else {
			t.Log(p.Unix(), name, format)
		}
	}
}
