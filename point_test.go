package opentsdb_test

import (
	"github.com/adbjesus/go-opentsdb"
	"testing"
	"time"
)

func TestNewPoint(t *testing.T) {
	// Fake test data
	var err error
	metric := "metric"
	ts := time.Now().Unix()
	value := 123
	tags := map[string]string{"tag1": "foo", "tag2": "bar"}

	// Expect failure if metric is empty
	expected := "PointError: Metric can not be empty"
	_, err = opentsdb.NewPoint("", ts, value, tags)
	if err == nil || err.Error() != expected {
		t.Error(
			"Expected", expected,
			"Got", err,
		)
	}

	// Expect failure if value isn't of type integer, float
	expected = "PointError: value must of type int, int8, int16, int32, int64 or float32"
	_, err = opentsdb.NewPoint(metric, ts, nil, tags)
	if err == nil || err.Error() != expected {
		t.Error(
			"Expected", expected,
			"Got", err,
		)
	}

}
