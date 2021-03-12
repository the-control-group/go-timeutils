package timeutils

import (
	"testing"
	"time"
)

var testFixtures = []struct {
	String               string
	Pretty               string
	ApproxPretty         string
	Duration             time.Duration
	NegativeString       string
	NegativePretty       string
	NegativeApproxPretty string
}{
	{
		String:               `15m`,
		Pretty:               `15m`,
		ApproxPretty:         `15m`,
		Duration:             15 * time.Minute,
		NegativeString:       `-15m`,
		NegativePretty:       `-15m`,
		NegativeApproxPretty: `-15m`,
	},
	{
		String:               `11h`,
		Pretty:               `11h`,
		ApproxPretty:         `11h`,
		Duration:             11 * time.Hour,
		NegativeString:       `-11h`,
		NegativePretty:       `-11h`,
		NegativeApproxPretty: `-11h`,
	},
	{
		String:               `1d1h`,
		Pretty:               `1 days, 1 hours`,
		ApproxPretty:         `~ 1 days, 1 hours`,
		Duration:             25 * time.Hour,
		NegativeString:       `-1d1h`,
		NegativePretty:       `-1 days, 1 hours`,
		NegativeApproxPretty: `~ -1 days, 1 hours`,
	},
	{
		String:               `6d2h5s`,
		Pretty:               `6 days, 2 hours 5s`,
		ApproxPretty:         `~ 6 days, 2 hours 5s`,
		Duration:             (6 * 24 * time.Hour) + (2 * time.Hour) + (5 * time.Second),
		NegativeString:       `-6d2h5s`,
		NegativePretty:       `-6 days, 2 hours 5s`,
		NegativeApproxPretty: `~ -6 days, 2 hours 5s`,
	},
	{
		String:               `13s`,
		Pretty:               `13s`,
		ApproxPretty:         `13s`,
		Duration:             13 * time.Second,
		NegativeString:       `-13s`,
		NegativePretty:       `-13s`,
		NegativeApproxPretty: `-13s`,
	},
	{
		String:               `3mos`,
		Pretty:               `3 months`,
		ApproxPretty:         `~ 3 months`,
		Duration:             3 * 30 * 24 * time.Hour,
		NegativeString:       `-3mos`,
		NegativePretty:       `-3 months`,
		NegativeApproxPretty: `~ -3 months`,
	},
	{
		String:               `3mo`,
		Pretty:               `3 months`,
		ApproxPretty:         `~ 3 months`,
		Duration:             3 * 30 * 24 * time.Hour,
		NegativeString:       `-3mo`,
		NegativePretty:       `-3 months`,
		NegativeApproxPretty: `~ -3 months`,
	},
}

func TestApproxBigDuration(t *testing.T) {
	for _, fixture := range testFixtures {
		{
			d, err := ParseApproxBigDuration([]byte(fixture.String))
			if err != nil {
				t.Error(err)
			}
			if time.Duration(d) != fixture.Duration {
				t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture.String, d, d.Pretty(), fixture.Duration, (time.Duration(d)-fixture.Duration)/time.Second)
			}
		}
		{
			d, err := ParseApproxBigDuration([]byte(fixture.Pretty))
			if err != nil {
				t.Error(err)
			}
			if time.Duration(d) != fixture.Duration {
				t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture.Pretty, d, d.Pretty(), fixture.Duration, (time.Duration(d)-fixture.Duration)/time.Second)
			}
		}
		{
			d, err := ParseApproxBigDuration([]byte(fixture.ApproxPretty))
			if err != nil {
				t.Error(err)
			}
			if time.Duration(d) != fixture.Duration {
				t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture.ApproxPretty, d, d.Pretty(), fixture.Duration, (time.Duration(d)-fixture.Duration)/time.Second)
			}
		}
		{
			d, err := ParseApproxBigDuration([]byte(fixture.NegativeString))
			if err != nil {
				t.Error(err)
			}
			if time.Duration(d) != -fixture.Duration {
				t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture.NegativeString, d, d.Pretty(), -fixture.Duration, (time.Duration(d)+fixture.Duration)/time.Second)
			}
		}
		{
			d, err := ParseApproxBigDuration([]byte(fixture.NegativePretty))
			if err != nil {
				t.Error(err)
			}
			if time.Duration(d) != -fixture.Duration {
				t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture.NegativePretty, d, d.Pretty(), -fixture.Duration, (time.Duration(d)+fixture.Duration)/time.Second)
			}
		}
		{
			d, err := ParseApproxBigDuration([]byte(fixture.NegativeApproxPretty))
			if err != nil {
				t.Error(err)
			}
			if time.Duration(d) != -fixture.Duration {
				t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture.NegativeApproxPretty, d, d.Pretty(), -fixture.Duration, (time.Duration(d)+fixture.Duration)/time.Second)
			}
		}
		// {
		// 	d := ApproxBigDuration(fixture.Duration)
		// 	if d.String() != fixture.String {
		// 		t.Errorf("Duration %d string is %s instead of %s", d, d.String(), fixture.String)
		// 	}
		// }
	}
}

func TestInterfaceToApproxBigDurationInt(t *testing.T) {
	var value int = 5
	var iface interface{} = value
	d, err := InterfaceToApproxBigDuration(iface)
	if err != nil {
		t.Error(err)
	}
	if d.String() != "5ns" {
		t.Errorf("Expected 5ns but got %s", d)
	}
}

func TestInterfaceToApproxBigDurationFloat64(t *testing.T) {
	var value float64 = 5
	var iface interface{} = value
	d, err := InterfaceToApproxBigDuration(iface)
	if err != nil {
		t.Error(err)
	}
	if d.String() != "5ns" {
		t.Errorf("Expected 5ns but got %s", d)
	}
}

func TestInterfaceToApproxBigDurationInt64(t *testing.T) {
	var value int64 = 5
	var iface interface{} = value
	d, err := InterfaceToApproxBigDuration(iface)
	if err != nil {
		t.Error(err)
	}
	if d.String() != "5ns" {
		t.Errorf("Expected 5ns but got %s", d)
	}
}
