package opentsdb_test

import (
	"testing"
	"time"

	"github.com/whitesmith/go-opentsdb"
)

var testOptions = opentsdb.Options{
	Host: "127.0.0.1",
	Port: 4242,
}

var testClient, _ = opentsdb.NewClient(testOptions)

func TestPut(t *testing.T) {
	tim := time.Now().Unix()
	p, _ := opentsdb.NewPoint("app-rankings.rank",
		tim,
		tim%10,
		map[string]string{"country": "pt", "id": "1"})

	p2, _ := opentsdb.NewPoint("app-rankings.rank",
		tim,
		(tim%10)+10,
		map[string]string{"country": "pt", "id": "9"})

	bp := opentsdb.NewBatchPoints()
	bp.AddPoint(p)
	bp.AddPoint(p2)

	_, err := testClient.Put(bp, "details")
	if err != nil {
		t.Error(
			"Expected", nil,
			"Got", err,
		)
	}
}

func TestGet(t *testing.T) {
	q, _ := opentsdb.NewQueryParams()
	q.Start = "6h-ago"
	q.Queries = append(q.Queries, opentsdb.Query{Aggregator: "sum", Metric: "app-rankings.rank"})

	_, err := testClient.Query(q)
	if err != nil {
		t.Error(
			"Expected", nil,
			"Got", err,
		)
	}
}
