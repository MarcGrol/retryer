package web

import "time"

type TaskStatus struct {
	UID            string
	Timestamp      time.Time
	Method         string
	RequestURL     string
	ResponseStatus int
	RequestBody    string `datastore:",noindex"`
	ResponseBody   string `datastore:",noindex"`
	Success        bool
	NumAttempts    int32
	MaxAttempts    int32
	Done           bool
}
