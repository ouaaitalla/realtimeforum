package helpers

import (
	"sync"
	"time"
)

type attempt struct {
	count    int
	firstAt  time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	attempts map[string]*attempt
	max      int
	window   time.Duration
	stopCh   chan struct{}
}

func NewRateLimiter(max int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		attempts: make(map[string]*attempt),
		max:      max,
		window:   window,
		stopCh:   make(chan struct{}),
	}

	// Start periodic cleanup every minute to prevent memory leaks
	go rl.cleanup()

	return rl
}

func (rl *RateLimiter) Stop() {
	close(rl.stopCh)
}

// Allow checks if a request from the given key (e.g. IP) is allowed.
// Returns (allowed bool, remaining int, resetAfter time.Duration).
func (rl *RateLimiter) Allow(key string) (bool, int, time.Duration) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	a, exists := rl.attempts[key]

	if !exists {
		rl.attempts[key] = &attempt{
			count:   0,
			firstAt: now,
		}

		a = rl.attempts[key]
	}

	// Reset if window has expired
	if now.Sub(a.firstAt) > rl.window {
		a.count = 0
		a.firstAt = now
	}

	a.count++

	remaining := rl.max - a.count

	if remaining < 0 {
		remaining = 0
	}

	resetAfter := rl.window - now.Sub(a.firstAt)

	if resetAfter < 0 {
		resetAfter = 0
	}

	if a.count > rl.max {
		return false, remaining, resetAfter
	}

	return true, remaining, resetAfter
}

// Reset clears all attempts for the given key (used on successful login).
func (rl *RateLimiter) Reset(key string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	delete(rl.attempts, key)
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)

	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
			rl.mu.Lock()

			now := time.Now()

			for key, a := range rl.attempts {
				if now.Sub(a.firstAt) > rl.window {
					delete(rl.attempts, key)
				}
			}

			rl.mu.Unlock()

		case <-rl.stopCh:
			return
		}
	}
}

// LoginRateLimiter is the shared instance for login endpoint: 5 attempts per minute.
var LoginRateLimiter = NewRateLimiter(5, 1*time.Minute)
