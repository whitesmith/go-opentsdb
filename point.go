package opentsdb

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Point struct {
	// Required
	// Opentsdb metric e.g.: "sys.info.cpu"
	Metric string `json:"metric"`

	// Required
	// Timestamp unix time e.g.: time.Now().Unix()
	Timestamp int64 `json:"timestamp"`

	// Required
	// Value to save, this can be either integer(8,16,32,64), float32
	// Different types will have unexpected behaviour
	Value interface{} `json:"value"`

	// Required
	// Map of tags, example: {"host": "deskop"}
	Tags map[string]string `json:"tags"`
}

func NewPoint(metric string, timestamp int64, value interface{}, tags map[string]string) (*Point, error) {
	if metric == "" {
		return nil, errors.New("PointError: Metric can not be empty")
	}

	t := reflect.TypeOf(value)
	if t == nil || (t.Kind() != reflect.Int &&
		t.Kind() != reflect.Int8 &&
		t.Kind() != reflect.Int16 &&
		t.Kind() != reflect.Int32 &&
		t.Kind() != reflect.Int64 &&
		t.Kind() != reflect.Float32) {
		return nil, errors.New("PointError: value must of type int, int8, int16, int32, int64 or float32")
	}

	return &Point{
		Metric:    metric,
		Timestamp: timestamp,
		Value:     value,
		Tags:      tags,
	}, nil
}

type BatchPoints struct {
	Points []*Point `json:""`
}

func NewBatchPoints() *BatchPoints {
	return &BatchPoints{}
}

func (bp *BatchPoints) AddPoint(p *Point) {
	bp.Points = append(bp.Points, p)
}

func (bp *BatchPoints) ToJson() ([]byte, error) {
	return json.Marshal(bp.Points)
}
