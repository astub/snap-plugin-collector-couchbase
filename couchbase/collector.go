package couchbase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// NewCollector constructs new metricCollector that will request given stats and
// fail if they are unavailable.
func NewCollector() *metricCollector {
	self := new(metricCollector)
	return self
}

// metricCollector implements logic for discovering available metrics
type metricCollector struct {
}

// Collect performs given set of calls (indicated by true value in metrics map).
// returns map of metric values (accessible by metric name). If any of requested
// calls fail error is returned.
func (mc *metricCollector) Collect(metrics map[int]bool) (map[string]interface{}, error) {
	s, err := GetSamples()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Discover performs metric discovery. Returns valid metric names and associated
// Call id's. If mandatory request fails error is returned. No error is returned
// when master or slave stats can't be read because server may not be configured
// to work in master-slave mode.
func (mc *metricCollector) Discover() ([]metric, error) {
	samples, err := GetSamples()
	if err != nil {
		return nil, err
	}

	res := []metric{}

	for key := range samples {
		res = append(res, metric{Name: key, Call: 0})
	}

	return res, nil
}

// metric contains name of metric and id of call that collects particular metric.
type metric struct {
	Name string
	Call int
}

func GetSamples() (samples map[string]interface{}, err error) {
	var username string = "admin"
	var passwd string = "password"

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:32817/pools/default/buckets/travel-sample/stats", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var couchdata map[string]interface{}
	err = json.Unmarshal(body, &couchdata)
	if err != nil {
		return
	}

	samples = couchdata["op"].(map[string]interface{})["samples"].(map[string]interface{})

	return
}