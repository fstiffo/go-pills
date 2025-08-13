package utils

import "time"

// ToDateOnly converts a time.Time to date-only (midnight) for date comparisons.
// This removes the time component, keeping only year, month, and day.
func ToDateOnly(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// IsSameDateAs checks if two times represent the same date (ignoring time).
func IsSameDateAs(t1, t2 time.Time) bool {
	return ToDateOnly(t1).Equal(ToDateOnly(t2))
}

// IsDateBefore checks if t1's date is before t2's date (ignoring time).
func IsDateBefore(t1, t2 time.Time) bool {
	return ToDateOnly(t1).Before(ToDateOnly(t2))
}

// IsDateAfter checks if t1's date is after t2's date (ignoring time).
func IsDateAfter(t1, t2 time.Time) bool {
	return ToDateOnly(t1).After(ToDateOnly(t2))
}

// IsDateBeforeOrEqual checks if t1's date is before or equal to t2's date (ignoring time).
func IsDateBeforeOrEqual(t1, t2 time.Time) bool {
	date1, date2 := ToDateOnly(t1), ToDateOnly(t2)
	return date1.Before(date2) || date1.Equal(date2)
}

// IsDateAfterOrEqual checks if t1's date is after or equal to t2's date (ignoring time).
func IsDateAfterOrEqual(t1, t2 time.Time) bool {
	date1, date2 := ToDateOnly(t1), ToDateOnly(t2)
	return date1.After(date2) || date1.Equal(date2)
}