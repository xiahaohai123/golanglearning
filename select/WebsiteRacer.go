package _select

import (
	"net/http"
	"time"
)

func Racer(urlA string, urlB string) string {
	durationA := measureResponseTime(urlA)
	durationB := measureResponseTime(urlB)

	if durationA < durationB {
		return urlA
	} else {
		return urlB
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
