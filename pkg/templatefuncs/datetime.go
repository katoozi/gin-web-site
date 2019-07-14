package templatefuncs

import (
	"fmt"
	"time"
)

// FormatAsDate take given datetime and return string represent in special format
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", year, month, day)
}
