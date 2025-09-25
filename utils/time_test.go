package utils

import (
	"testing"
	"time"
)

// TestToDateOnly verifies that the ToDateOnly function correctly strips the time component
// (hour, minute, second, nanosecond) from a time.Time value, returning a new time.Time
// representing only the date portion (year, month, day) in the same location.
func TestToDateOnly(t *testing.T) {
	// Test with a time that has hour, minute, second, nanosecond
	originalTime := time.Date(2023, 12, 15, 14, 30, 45, 123456789, time.UTC)
	expected := time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC)

	result := ToDateOnly(originalTime)

	if !result.Equal(expected) {
		t.Errorf("ToDateOnly() = %v, want %v", result, expected)
	}
}

// TestIsSameDateAs tests the IsSameDateAs function to ensure it correctly determines
// whether two time.Time values represent the same calendar date, regardless of the time.
// It verifies that IsSameDateAs returns true for times on the same date and false for times on different dates.
func TestIsSameDateAs(t *testing.T) {
	time1 := time.Date(2023, 12, 15, 9, 30, 0, 0, time.UTC)
	time2 := time.Date(2023, 12, 15, 18, 45, 30, 0, time.UTC)
	time3 := time.Date(2023, 12, 16, 9, 30, 0, 0, time.UTC)

	if !IsSameDateAs(time1, time2) {
		t.Error("IsSameDateAs() should return true for same dates with different times")
	}

	if IsSameDateAs(time1, time3) {
		t.Error("IsSameDateAs() should return false for different dates")
	}
}

// TestIsDateBefore tests the IsDateBefore function to ensure it correctly determines
// whether the first date is before the second date. It verifies that the function
// returns true when the first date precedes the second, and false otherwise.
func TestIsDateBefore(t *testing.T) {
	earlier := time.Date(2023, 12, 15, 18, 30, 0, 0, time.UTC)
	later := time.Date(2023, 12, 16, 9, 30, 0, 0, time.UTC)

	if !IsDateBefore(earlier, later) {
		t.Error("IsDateBefore() should return true when first date is before second")
	}

	if IsDateBefore(later, earlier) {
		t.Error("IsDateBefore() should return false when first date is after second")
	}
}

// TestIsDateAfterOrEqual tests the IsDateAfterOrEqual function to ensure it correctly determines
// whether the first date is the same as or after the second date, ignoring the time component.
// It verifies the following scenarios:
//   - Returns true for dates with the same day, regardless of time.
//   - Returns true when the first date is after the second date.
//   - Returns false when the first date is before the second date.
func TestIsDateAfterOrEqual(t *testing.T) {
	date1 := time.Date(2023, 12, 15, 9, 30, 0, 0, time.UTC)
	date2 := time.Date(2023, 12, 15, 18, 45, 0, 0, time.UTC) // Same date, different time
	date3 := time.Date(2023, 12, 16, 9, 30, 0, 0, time.UTC)
	date4 := time.Date(2023, 12, 14, 9, 30, 0, 0, time.UTC)

	// Same date should return true
	if !IsDateAfterOrEqual(date1, date2) {
		t.Error("IsDateAfterOrEqual() should return true for same dates")
	}

	// Later date should return true
	if !IsDateAfterOrEqual(date3, date1) {
		t.Error("IsDateAfterOrEqual() should return true when first date is after second")
	}

	// Earlier date should return false
	if IsDateAfterOrEqual(date4, date1) {
		t.Error("IsDateAfterOrEqual() should return false when first date is before second")
	}
}
