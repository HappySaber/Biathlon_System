package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseEventTime(timeStr string) (time.Time, error) {
	return time.Parse("15:04:05.000", strings.Trim(timeStr, "[]"))
}

func FormatTime(t time.Time) string {
	return t.Format("15:04:05.000")
}

func ParseStartDelta(deltaStr string) (time.Duration, error) {
	timeParts := strings.Split(deltaStr, ".")
	hasMilliseconds := len(timeParts) == 2

	hmsParts := strings.Split(timeParts[0], ":")
	if len(hmsParts) != 3 {
		return 0, fmt.Errorf("invalid time format, expected hh:mm:ss[.ms]")
	}

	h, err := strconv.Atoi(hmsParts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %v", err)
	}
	m, err := strconv.Atoi(hmsParts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %v", err)
	}
	s, err := strconv.Atoi(hmsParts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %v", err)
	}

	var ms int
	if hasMilliseconds {
		ms, err = strconv.Atoi(timeParts[1])
		if err != nil {
			return 0, fmt.Errorf("invalid milliseconds: %v", err)
		}
	}

	return time.Duration(h)*time.Hour +
		time.Duration(m)*time.Minute +
		time.Duration(s)*time.Second +
		time.Duration(ms)*time.Millisecond, nil
}

func FormatDuration(d time.Duration) string {
	total := int(d.Seconds())
	hours := total / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60
	millis := (d.Milliseconds() % 1000)
	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, millis)
}
func FormatTotalTime(start, end time.Time) string {
	return FormatDuration(end.Sub(start))
}
