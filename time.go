package time

import (
	"time"

	l "github.com/vsjadeja/logger"
)

type TimeProvider interface {
	Now() time.Time
	Elapsed() int64
	IsTimeout() bool
}

type TimeService struct {
	start          time.Time
	processTimeout float64
	logger         l.LoggerInterface
}

// NewTimeService creates a new instance of the TimeService object
func NewTimeService(logger l.LoggerInterface, processTimeout float64) *TimeService {
	s := &TimeService{logger: logger, processTimeout: processTimeout}
	s.start = s.Now()
	return s
}

// Now returns a current time
func (p *TimeService) Now() time.Time {
	return time.Now()
}

// Elapsed returns a elapsed time in milliseconds
func (p *TimeService) Elapsed() int64 {
	return time.Since(p.start).Milliseconds()
}

// IsTimeout returns timed out or not
func (p *TimeService) IsTimeout() bool {
	diff := time.Since(p.start).Seconds()

	return diff >= p.processTimeout
}
