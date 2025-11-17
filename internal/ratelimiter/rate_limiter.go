package ratelimiter

import "time"

type Limiter interface {
	Allow(ip string) (bool, time.Duration)
}

type RateLimiterConfig struct {
	MaxRequests int
	Duration    time.Duration
	Enabled     bool
}
