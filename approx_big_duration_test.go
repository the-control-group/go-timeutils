package timeutils

import (
	"testing"
	"time"
)

func TestParseApproxBigDuration(t *testing.T) {
	var fixtures = []map[string]interface{}{
		{
			"bytes":    []byte(`15m`),
			"duration": 15 * time.Minute,
		},
		{
			"bytes":    []byte(`11h`),
			"duration": 11 * time.Hour,
		},
		{
			"bytes":    []byte(`1d1h`),
			"duration": 25 * time.Hour,
		},
		{
			"bytes":    []byte(`6d2h5s`),
			"duration": (6 * 24 * time.Hour) + (2 * time.Hour) + (5 * time.Second),
		},
		{
			"bytes":    []byte(`13s`),
			"duration": 13 * time.Second,
		},
		{
			"bytes":    []byte(`3mos`),
			"duration": 3 * 30 * 24 * time.Hour,
		},
	}

	for _, fixture := range fixtures {
		d, err := ParseApproxBigDuration(fixture["bytes"].([]byte))
		if err != nil {
			t.Error(err)
		}
		if time.Duration(d) != fixture["duration"].(time.Duration) {
			t.Errorf("parsed duration %s is %d (%s) instead of %d. Diff by %d seconds", fixture["bytes"], d, d.Pretty(), fixture["duration"], (time.Duration(d)-fixture["duration"].(time.Duration))/time.Second)
		}
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
