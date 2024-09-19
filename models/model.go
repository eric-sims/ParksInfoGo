package models

type DisneyQueues struct {
	QueueTimes []Ride `json:"queue_times"`
}

// QueueTimesResponse - follows the return structure of the QueueTimes api
// Found at: https://queue-times.com/parks/<park_id>/queue_times.json
type QueueTimesResponse struct {
	Lands []Land `json:"lands"`
}

type Land struct {
	Name  string `json:"name"`
	Rides []Ride `json:"rides"`
}

type Ride struct {
	Name     string `json:"name"`
	IsOpen   bool   `json:"is_open"`
	WaitTime int    `json:"wait_time"`
}
