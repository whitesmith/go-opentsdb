package opentsdb_test

import (
	"testing"
	"time"

	"github.com/adbjesus/go-opentsdb"
)

var testOptions = opentsdb.Options{
	Host: "127.0.0.1",
	Port: 4242,
}

var testClient, _ = opentsdb.NewClient(testOptions)

func TestPut(t *testing.T) {
	p, _ := opentsdb.NewPoint("app-rankings.rank",
		time.Now().Unix(),
		time.Now().Unix()%10,
		map[string]string{"country": "us"})

	p2, _ := opentsdb.NewPoint("app-rankings.rank",
		time.Now().Unix()+1,
		time.Now().Unix()%10+1,
		map[string]string{"country": "us"})

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
