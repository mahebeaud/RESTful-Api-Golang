package controllers

import (
	"time"
)

// GetCurrentDateISO8601 returns the current date in ISO 8601 format
func GetCurrentDateISO8601() string {
	currentTime := time.Now()
	return currentTime.Format(time.RFC3339)
}