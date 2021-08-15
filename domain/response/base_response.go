package response

import "time"

//Base is
type Base struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp" example:"2020-12-21T08:01:55.570498561Z"`
	Data       interface{} `json:"data"`
}
