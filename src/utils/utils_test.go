package utils

import (
	"math"
	"testing"
	"time"
)

func TestCalculateSpeed(t *testing.T) {
	distance := 1000
	duration := 5*time.Minute + 30*time.Second
	wantSpeed := 3.030303
	res := CalculateSpeed(distance, duration)

	round := func(x float64, prec int) float64 {
		pow := math.Pow(10, float64(prec))
		return math.Round(x*pow) / pow
	}
	if round(res, 6) != round(wantSpeed, 6) {
		t.Errorf("Result was incorrect, got: %f, want: %f.", res, wantSpeed)
	}
}

func TestFormatTime(t *testing.T) {
	tm := time.Date(0, 1, 1, 9, 15, 30, 500000000, time.UTC)
	want := "09:15:30.500"
	res := FormatTime(tm)

	if res != want {
		t.Errorf("Result was incorrect, got: %s, want: %s", res, want)
	}
}

func TestParseStartDelta(t *testing.T) {
	tests := []struct {
		input string
		want  time.Duration
	}{
		{"01:30:45", 1*time.Hour + 30*time.Minute + 45*time.Second},
		{"00:01:30.500", 1*time.Minute + 30*time.Second + 500*time.Millisecond},
	}

	for _, tt := range tests {
		res, err := ParseStartDelta(tt.input)
		if err != nil {
			t.Errorf("Unexpected error for %s: %v", tt.input, err)
		}
		if res != tt.want {
			t.Errorf("For %s got: %v, want: %v", tt.input, res, tt.want)
		}
	}
}

func TestFormatDuration(t *testing.T) {
	d := 2*time.Hour + 5*time.Minute + 30*time.Second + 123*time.Millisecond
	want := "02:05:30.123"
	res := FormatDuration(d)

	if res != want {
		t.Errorf("Result was incorrect, got: %s, want: %s", res, want)
	}
}

func TestFormatTotalTime(t *testing.T) {
	start := time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 1, 12, 30, 15, 500000000, time.UTC)
	want := "02:30:15.500"
	res := FormatTotalTime(start, end)

	if res != want {
		t.Errorf("Result was incorrect, got: %s, want: %s", res, want)
	}
}
