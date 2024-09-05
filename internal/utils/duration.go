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
	if strings.Contains(date, "Z") {
		format = "2006-01-02T15:04:05.000000Z"
	} else if strings.Contains(date, "T") {
		format = "2006-01-02T15:04:05"
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
	duration := t.Sub(now)
	sign := " "
	if duration < 0 {
		sign = "-"
		duration = -duration
	}

	if duration.Hours() > hoursInYear {
		return strings.Replace(fmt.Sprintf("%s%.1fy", sign, duration.Hours()/hoursInYear), ".0", "", 1)
	} else if duration.Hours() > hoursInMonth {
		return fmt.Sprintf("%s%dmo", sign, int(duration.Hours()/hoursInMonth))
	} else if duration.Hours() > hoursInWeek {
		return fmt.Sprintf("%s%dw", sign, int(duration.Hours()/hoursInWeek))
	} else if duration.Hours() > hoursInDay {
		return fmt.Sprintf("%s%dd", sign, int(duration.Hours()/hoursInDay))
	} else if duration.Hours() > 1 {
		return fmt.Sprintf("%s%dh", sign, int(duration.Hours()))
	} else if duration.Minutes() > 1 {
		return fmt.Sprintf("%s%dm", sign, int(duration.Minutes()))
	} else if duration.Seconds() > 1 {
		return fmt.Sprintf("%s%ds", sign, int(duration.Seconds()))
	}

	return fallback
}
