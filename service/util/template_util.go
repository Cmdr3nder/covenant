package util

import (
	"fmt"
	"time"
)

// MonthAsInt formats a month as its integer representation with exactly 2 digits. Ex: June -> 06
func MonthAsInt(month time.Month) string {
	return fmt.Sprintf("%02d", month)
}
