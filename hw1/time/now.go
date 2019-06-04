package time

import (
	"github.com/beevik/ntp"
	"time"
)

func Now() string {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return ""
	}

	return currentTime.Format(time.RFC3339)
}
