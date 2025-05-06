package models

import "time"

type Competitor struct {
	ID              string
	TotalTime       string
	LapTimes        []string
	LapStartTime    time.Time
	CurrentLap      int
	LapAvgSpeeds    []float64
	PenaltyStart    time.Time
	PenaltyDuration time.Duration
	PenaltyTime     time.Time
	PenaltyAvgSpeed float64
	Hits            int
	Shots           int
	Disqualified    bool
	DidNotFinish    bool
	ActualStart     time.Time
	ScheduledStart  time.Time
}

var CompetitorInfo []Competitor
