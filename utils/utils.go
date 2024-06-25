package utils

import (
	"fmt"
	"strings"
	"time"
)

type TimeSlot struct {
	Start time.Time
	End   time.Time
}

func ParseTimeSlot(timeSlotStr string) (TimeSlot, error) {
	layout := "15:04"
	times := strings.Split(timeSlotStr, "-")
	if len(times) != 2 {
		return TimeSlot{}, fmt.Errorf("invalid time slot format")
	}

	start, err := time.Parse(layout, times[0])
	if err != nil {
		return TimeSlot{}, err
	}

	end, err := time.Parse(layout, times[1])
	if err != nil {
		return TimeSlot{}, err
	}

	return TimeSlot{Start: start, End: end}, nil
}

func ConvertToTokyoTime() time.Time {
	// Load the Asia/Tokyo location
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return time.Now() // Fallback to UTC
	}

	// Convert the current time to Tokyo time
	tokyoTime := time.Now().In(loc)
	return tokyoTime
}

func IsCurrentTimeInSlot(slot TimeSlot) bool {

	now := ConvertToTokyoTime()
	start := time.Date(now.Year(), now.Month(), now.Day(), slot.Start.Hour(), slot.Start.Minute(), 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), slot.End.Hour(), slot.End.Minute(), 0, 0, now.Location())

	if now.After(start) && now.Before(end) {
		return true
	}
	return false
}

func CheckIfSlotIsInCurrentTimeWindow(slotStr string) bool {
	slot, err := ParseTimeSlot(slotStr)
	if err != nil {
		fmt.Println("Error parsing time slot:", err)
		return false
	}

	if IsCurrentTimeInSlot(slot) {
		return true
	}

	return false
}
