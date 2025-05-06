package controllers

import (
	"biathlon_system/src/models"
	"biathlon_system/src/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func GenerateReport() {

	fmt.Println("\n\nResulting table \n ```")
	sortedCompetitors := SortCompetitorsByTime(models.CompetitorInfo)

	for i := range sortedCompetitors {

		sortedCompetitors[i].LapAvgSpeeds = make([]float64, models.Cfg.Laps)
		sortedCompetitors[i].LapAvgSpeeds = CalculateAvgSpeeds(sortedCompetitors[i])
		sortedCompetitors[i].TotalTime = SetStatusOrTime(sortedCompetitors[i])
		competitor := sortedCompetitors[i]
		if competitor.Disqualified {
			fmt.Println(competitor.TotalTime, competitor.ID, "{-}", "{-}", "{-}")
			continue
		}
		arrayOfLapAndAvgSpeed := make([]string, models.Cfg.Laps)

		for j := range arrayOfLapAndAvgSpeed {

			if competitor.LapTimes[j] != "00:00:00.000" {
				arrayOfLapAndAvgSpeed[j] = "{" + competitor.LapTimes[j] + ", " + strconv.FormatFloat(competitor.LapAvgSpeeds[j], 'f', 2, 64) + "}"
			} else {
				arrayOfLapAndAvgSpeed[j] = "{" + ", " + "}"
			}
		}
		competitor.PenaltyAvgSpeed = CalculateAvgPenanltySpeeds(competitor)
		outputOfLapAndAvgSpeed := "[" + strings.Join(arrayOfLapAndAvgSpeed, `, `) + "]"
		outputOfPenalty := ""
		if competitor.PenaltyAvgSpeed == 0 {
			outputOfPenalty = "{, }"

		} else {
			outputOfPenalty = "{" + utils.FormatDuration(competitor.PenaltyDuration) + ", " + strconv.FormatFloat(competitor.PenaltyAvgSpeed, 'f', 2, 64) + "}"

		}

		HitsShots := strconv.Itoa(competitor.Hits) + "/" + strconv.Itoa(competitor.Shots)
		fmt.Println(competitor.TotalTime, competitor.ID, outputOfLapAndAvgSpeed, outputOfPenalty, HitsShots)
	}
}

func IsStratedInTime(record models.Event) bool {
	competitorID, err := strconv.Atoi(record.CompetitorID)
	if err != nil {
		log.Fatal("Error converting CompetitorID:", err)
	}
	StartTime, err := utils.ParseEventTime(record.Time)

	if err != nil {
		log.Fatal("Error converting time of start:", err)
	}

	StartDelta, err := utils.ParseStartDelta(models.Cfg.StartDelta)
	if err != nil {
		log.Fatal("Error converting time of start delta:", err)
	}

	models.CompetitorInfo[competitorID-1].ActualStart = StartTime
	EdgeStartTime := models.CompetitorInfo[competitorID-1].ScheduledStart.Add(StartDelta)
	return !models.CompetitorInfo[competitorID-1].ActualStart.After(EdgeStartTime)
}

func SetStatusOrTime(competitor models.Competitor) string {
	if competitor.DidNotFinish {
		return "[NotFinished]"
	} else if competitor.Disqualified {
		return "[NotStarted]"
	}
	return "[" + competitor.TotalTime + "]"
}

func SortCompetitorsByTime(competitors []models.Competitor) []models.Competitor {
	sorted := make([]models.Competitor, len(competitors))
	copy(sorted, competitors)

	sort.Slice(sorted, func(i, j int) bool {
		timeI, errI := utils.ParseEventTime(sorted[i].TotalTime)
		if errI != nil {
			return false
		}

		timeJ, errJ := utils.ParseEventTime(sorted[j].TotalTime)
		if errJ != nil {
			return true
		}
		return timeI.Before(timeJ)
	})

	return sorted
}

func CalculateAvgSpeeds(competitor models.Competitor) []float64 {
	for i := range competitor.LapTimes {
		delta, err := utils.ParseStartDelta(competitor.LapTimes[i])
		if err != nil {
			log.Fatal("Error while parsing string to time.Duration", err)
		}
		competitor.LapAvgSpeeds[i] = utils.CalculateSpeed(models.Cfg.LapLen, delta)

	}
	return competitor.LapAvgSpeeds
}

func CalculateAvgPenanltySpeeds(competitor models.Competitor) float64 {
	PenaltyLen := (competitor.Shots - competitor.Hits) * models.Cfg.PenaltyLen
	if PenaltyLen == 0 {
		return 0
	}
	penaltyAvgSpeed := utils.CalculateSpeed(PenaltyLen, competitor.PenaltyDuration)
	return penaltyAvgSpeed
}
