package controllers

import (
	"biathlon_system/src/models"
	"math"
	"testing"
	"time"
)

func TestIsStartedInTime(t *testing.T) {
	models.CompetitorInfo = []models.Competitor{}

	models.Cfg.StartDelta = "00:01:30"

	scheduledStart := time.Date(2000, 1, 1, 9, 30, 0, 0, time.UTC)
	actualStart := time.Date(2000, 1, 1, 9, 30, 2, 0, time.UTC)

	competitor1 := models.Competitor{
		ID:             "1",
		ScheduledStart: scheduledStart,
		ActualStart:    actualStart,
	}

	models.CompetitorInfo = append(models.CompetitorInfo, competitor1)

	event := models.Event{
		ID:           "1",
		Time:         "09:30:02.000",
		CompetitorID: "1",
	}

	result := IsStratedInTime(event)

	if !result {
		t.Errorf("Result was incorrect, got: %t, want: %t", result, true)
	}
	t.Log("Tested result1")
}

func TestSetStatusOrTime(t *testing.T) {
	competitor1 := models.Competitor{
		ID:           "1",
		DidNotFinish: true,
	}
	competitor2 := models.Competitor{
		ID:           "2",
		Disqualified: true,
	}
	competitor3 := models.Competitor{
		ID:        "3",
		TotalTime: "10:30:00.000",
	}
	result1 := SetStatusOrTime(competitor1)

	if result1 != "[NotFinished]" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result1, "[NotFinished]")
	}
	t.Log("Tested SetStatusOrTime 1")
	result2 := SetStatusOrTime(competitor2)

	if result2 != "[NotStarted]" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result2, "[NotStarted]")
	}
	t.Log("Tested SetStatusOrTime 2")
	result3 := SetStatusOrTime(competitor3)

	if result3 != "[10:30:00.000]" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result3, "[10:30:00.000]")
	}
	t.Log("Tested SetStatusOrTime 3")
}

func TestCalculateAvgPenanltySpeeds(t *testing.T) {
	models.Cfg.PenaltyLen = 2
	competitor1 := models.Competitor{
		ID:              "1",
		Shots:           10,
		Hits:            9,
		PenaltyDuration: 1 * time.Second,
	}
	result := CalculateAvgPenanltySpeeds(competitor1)
	if result != 2 {
		t.Errorf("Result was incorrect, got: %f, want: %d", result, 2)
	}
	t.Log("Tested CalculateAvgPenanltySpeeds")
}

func TestCalculateAvgSpeeds(t *testing.T) {
	models.Cfg.LapLen = 3500
	competitor1 := models.Competitor{
		ID:           "1",
		LapTimes:     []string{"00:13:20.939", "00:13:01.202"},
		LapAvgSpeeds: []float64{0, 0},
	}

	avgSpeeds := CalculateAvgSpeeds(competitor1)

	round := func(x float64, prec int) float64 {
		pow := math.Pow(10, float64(prec))
		return math.Round(x*pow) / pow
	}
	if round(avgSpeeds[0], 6) != 4.369871 || round(avgSpeeds[1], 6) != 4.480275 {
		t.Errorf("Result was incorrect, got: %.6f, want: %.6f and got: %.6f, want: %.6f",
			avgSpeeds[0], 4.369871, avgSpeeds[1], 4.480275)
	}
}

func TestSortCompetitorsByTime(t *testing.T) {
	comps := []models.Competitor{
		{
			ID:        "1",
			TotalTime: "10:00:00.000",
		},
		{
			ID:        "2",
			TotalTime: "10:02:00.000",
		},
		{
			ID:        "3",
			TotalTime: "10:01:00.000",
		},
	}

	sorted := SortCompetitorsByTime(comps)

	if sorted[0].ID != "1" && sorted[1].ID != "3" && sorted[2].ID != "2" {
		t.Errorf("Result was incorrect, got: %s, want: %s,%s,%s", "1,3,2", sorted[0].ID, sorted[1].ID, sorted[2].ID)
	}
}
