package controllers

import (
	"biathlon_system/src/models"
	"testing"
)

func TestIsDisqualified(t *testing.T) {
	competitor1 := models.Competitor{
		ID:           "2",
		Disqualified: false,
	}
	event := models.Event{
		CompetitorID: competitor1.ID,
	}
	models.CompetitorInfo = append(models.CompetitorInfo, competitor1)
	result1 := IsDisqualified(event)
	if result1 {
		t.Errorf("Result was incorrect, got: %t, want: %s.", result1, "true")

	}
	t.Log("Tested result1")

	competitor2 := models.Competitor{
		ID:           "3",
		Disqualified: true,
	}

	event = models.Event{
		CompetitorID: competitor2.ID,
	}
	models.CompetitorInfo = append(models.CompetitorInfo, competitor2)
	result2 := IsDisqualified(event)
	if !result2 {
		t.Errorf("Result was incorrect, got: %t, want: %s.", result2, "false")

	}
	t.Log("Tested result2")
}

func TestAddCompetitorsInSlice(t *testing.T) {

	event := []models.Event{
		{
			Time:         " ",
			ID:           "1",
			CompetitorID: "1",
			ExtraParam:   " ",
		},
		{
			Time:         " ",
			ID:           "2",
			CompetitorID: "2",
			ExtraParam:   " ",
		},
		{
			Time:         " ",
			ID:           "4",
			CompetitorID: "3",
			ExtraParam:   " ",
		},
	}

	competitors := AddCompetitorsInSlice(event)
	if len(competitors) != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(competitors), "3")

	}
	t.Log("Tested result1")
	event2 := []models.Event{
		{
			Time:         " ",
			ID:           "1",
			CompetitorID: "1",
			ExtraParam:   " ",
		},
		{
			Time:         " ",
			ID:           "2",
			CompetitorID: "3",
			ExtraParam:   " ",
		},
		{
			Time:         " ",
			ID:           "4",
			CompetitorID: "3",
			ExtraParam:   " ",
		},
	}

	competitors2 := AddCompetitorsInSlice(event2)
	if len(competitors2) != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(competitors), "2")

	}
	t.Log("Tested result2")
}

func TestCheckCompetitorInSlice(t *testing.T) {
	competitors := []models.Competitor{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
		{
			ID: "3",
		},
	}
	res1 := CheckCompetitorInSlice(competitors, "1")

	if res1 {
		t.Errorf("Result was incorrect, got: %t, want: %s.", res1, "true")
	}
	t.Log("Tested result1")
	res2 := CheckCompetitorInSlice(competitors, "6")

	if !res2 {
		t.Errorf("Result was incorrect, got: %t, want: %s.", res2, "false")
	}
	t.Log("Tested result2")
}
