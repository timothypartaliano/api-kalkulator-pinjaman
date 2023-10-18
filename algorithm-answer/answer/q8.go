package answer

import (
	"fmt"
)

func calculateWorkDuration(checkIn, checkOut string) (int, int) {
	checkInHour, checkInMinute := convertTo24Hour(checkIn)
	checkOutHour, checkOutMinute := convertTo24Hour(checkOut)

	totalMinutes := (checkOutHour*60 + checkOutMinute) - (checkInHour*60 + checkInMinute)

	workHours := totalMinutes / 60
	workMinutes := totalMinutes % 60

	return workHours, workMinutes
}

func convertTo24Hour(time12Hour string) (int, int) {
	var hour, minute int
	_, err := fmt.Sscanf(time12Hour, "%d:%d", &hour, &minute)
	if err != nil {
		panic("Invalid time format")
	}

	if hour == 12 {
		hour = (hour - 12) % 12
	}

	return hour, minute
}

func Q8() {
	checkInTime := "8:00"
	checkOutTime := "12:30"

	workHours, workMinutes := calculateWorkDuration(checkInTime, checkOutTime)
	fmt.Printf("Work Duration: %d hours and %d minutes\n", workHours, workMinutes)
}