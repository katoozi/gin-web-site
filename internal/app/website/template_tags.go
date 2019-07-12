package website

import (
	"fmt"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", year, month, day)
}
