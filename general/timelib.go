package general

import (
	"strconv"
	"time"
)

type Timelib struct {
	Current time.Time
}


func (self Timelib) Timestamp() string {
	timestamp := strconv.FormatInt(self.Current.UTC().UnixNano(), 10)
	timestamp = timestamp[:10]
	return timestamp
}