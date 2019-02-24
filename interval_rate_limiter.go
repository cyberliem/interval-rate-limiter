package ratelimiter

import (
	"net/http"
	"time"
)

// IntervalRateLimiter is the object used to limit the rate
type IntervalRateLimiter struct {
	tk *time.Ticker
}

// NewIntervalRateLimiter return a new IntervalRateLimiter
func NewIntervalRateLimiter(rqInterval time.Duration) *IntervalRateLimiter {
	ticker := time.NewTicker(rqInterval)
	return &IntervalRateLimiter{
		tk: ticker,
	}
}

//ForwardRequest check if the request is eligible to forward.
func (irl *IntervalRateLimiter) ForwardRequest(c *http.Client, req *http.Request) (*http.Response, error) {
	<-irl.tk.C
	return c.Do(req)
}
