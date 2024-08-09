package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	hoursInDay   = 24
	hoursInWeek  = 24 * 7
	hoursInMonth = 24 * 30
	hoursInYear  = 24 * 365
)

func parseDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, errors.New("nil date")
	}

	// Automatic date parsing
	format := "2006-01-02"
	if strings.Contains(date, "T") {
		format = "2006-01-02T15:04:05Z"
	}

	t, err := time.Parse(format, date)
	if err != nil {
		return t, err
	}

	return t, nil
}

func ShortDuration(date string, fallback string) string {
	t, err := parseDate(date)
	if err != nil {
		return fallback
	}

	now := time.Now()
	duration := now.Sub(t)

	if duration.Hours() > hoursInYear {
		return strings.Replace(fmt.Sprintf("%.1fy", duration.Hours()/hoursInYear), ".0", "", 1)
	} else if duration.Hours() > hoursInMonth {
		return fmt.Sprintf("%dmo", int(duration.Hours()/hoursInMonth))
	} else if duration.Hours() > hoursInWeek {
		return fmt.Sprintf("%dw", int(duration.Hours()/hoursInWeek))
	} else if duration.Hours() > hoursInDay {
		return fmt.Sprintf("%dd", int(duration.Hours()/hoursInDay))
	} else if duration.Hours() > 1 {
		return fmt.Sprintf("%dh", int(duration.Hours()))
	} else if duration.Minutes() > 1 {
		return fmt.Sprintf("%dm", int(duration.Minutes()))
	} else if duration.Seconds() > 1 {
		return fmt.Sprintf("%ds", int(duration.Seconds()))
	}

	return fallback
}
