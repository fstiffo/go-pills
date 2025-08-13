package utils

import (
	"testing"
	"time"
)

func TestToDateOnly(t *testing.T) {
	// Test with a time that has hour, minute, second, nanosecond
	originalTime := time.Date(2023, 12, 15, 14, 30, 45, 123456789, time.UTC)
	expected := time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC)
	
	result := ToDateOnly(originalTime)
	
	if !result.Equal(expected) {
		t.Errorf("ToDateOnly() = %v, want %v", result, expected)
	}
}

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