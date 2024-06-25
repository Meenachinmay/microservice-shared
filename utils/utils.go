package utils

import (
	"fmt"
	"log"
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

	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return TimeSlot{}, err
	}

	start, err := time.ParseInLocation(layout, times[0], location)
	if err != nil {
		return TimeSlot{}, err
	}

	end, err := time.ParseInLocation(layout, times[1], location)
	if err != nil {
		return TimeSlot{}, err
	}

	return TimeSlot{Start: start, End: end}, nil
}

func IsCurrentTimeInSlot(slot TimeSlot) bool {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	now := time.Now().In(location)
	start := time.Date(now.Year(), now.Month(), now.Day(), slot.Start.Hour(), slot.Start.Minute(), 0, 0, location)
	end := time.Date(now.Year(), now.Month(), now.Day(), slot.End.Hour(), slot.End.Minute(), 0, 0, location)

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
