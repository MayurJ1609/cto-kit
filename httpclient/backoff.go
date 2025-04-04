package httpclient

import (
	"math"
	"time"

	"math/rand"
)

type BackoffPolicy func(attemptCount int) time.Duration

var (
	// DefaultBackoffPolicy is the default backoff policy used by the httpclient.
	DefaultBackoffPolicy = ExponentialBackoff(100*time.Millisecond, 2.0, 10*time.Second)

	ConstantBackoff = func(constantWait time.Duration, maxJitter time.Duration) BackoffPolicy {
		if constantWait < 0 {
			constantWait = 0
		}
		if maxJitter < 0 {
			maxJitter = 0
		}
		return func(attemptCount int) time.Duration {
			return constantWait + randJitter(maxJitter)
		}
	}

	LinearBackoff = func(minWait time.Duration, maxWait time.Duration, maxJitter time.Duration) BackoffPolicy {
		if minWait < 0 {
			minWait = 0
		}
		if maxJitter < 0 {
			maxJitter = 0
		}
		if maxWait < minWait {
			maxWait = 0
		}
		return func(attemptCount int) time.Duration {
			nextWait := time.Duration(attemptCount-1)*minWait + randJitter(maxJitter)
			if maxWait > 0 {
				nextWait = minDuration(nextWait, maxWait)
			}
			return nextWait
		}
	}

	ExponentialBackoff = func(minWait time.Duration, maxWait time.Duration, maxJitter time.Duration) BackoffPolicy {
		if minWait <= 0 {
			minWait = 0
		}
		if maxJitter < 0 {
			maxJitter = 0
		}
		if maxWait < minWait {
			maxWait = 0
		}
		return func(attemptCount int) time.Duration {
			nextWait := time.Duration(math.Pow(2, float64(attemptCount-1)))*minWait + randJitter(maxJitter)
			if maxWait > 0 {
				nextWait = minDuration(nextWait, maxWait)
			}
			return nextWait
		}
	}
)

func minDuration(duration1 time.Duration, duration2 time.Duration) time.Duration {
	if duration1 < duration2 {
		return duration1
	}
	return duration2
}

func randJitter(maxJitter time.Duration) time.Duration {
	if maxJitter <= 0 {
		return 0
	}
	return time.Duration(rand.Int63n(int64(maxJitter)))
}
