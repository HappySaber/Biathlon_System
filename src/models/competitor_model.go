package models

import "time"

type Competitor struct {
	ID              string
	TotalTime       string // Общее время (или "NotStarted"/"NotFinished")
	LapTimes        []string
	LapAvgSpeeds    []float64
	PenaltyTime     string
	PenaltyAvgSpeed float64
	Hits            int
	Shots           int
	Disqualified    bool
	DidNotFinish    bool
	ActualStart     time.Time
	ScheduledStart  time.Time
}

var CompetitorInfo []Competitor
