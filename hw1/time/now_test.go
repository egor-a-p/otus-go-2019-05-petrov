package time

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	currentTime := Now()

	_, err := time.Parse(time.RFC3339, currentTime)

	if err != nil {
		t.Errorf("Fail: %s", err)
	} else {
		t.Logf("Success: %s", currentTime)
	}
}
