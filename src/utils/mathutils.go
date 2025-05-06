package utils

import "time"

func CalculateSpeed(distanceMeters int, duration time.Duration) float64 {
	return float64(distanceMeters) / duration.Seconds()
}
