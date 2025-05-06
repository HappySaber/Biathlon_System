package controllers

import (
	"biathlon_system/src/models"
	"biathlon_system/src/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

var EventsList = map[string]string{
	"1":  "The competitor(%s) registered",
	"2":  "The start time for the competitor(%s) was set by a draw to %s)",
	"3":  "The competitor(%s) is on the start line",
	"4":  "The competitor(%s) has started",
	"5":  "The competitor(%s) is on the firing range(%s)",
	"6":  "The target(%s) has been hit by competitor(%s)",
	"7":  "The competitor(%s) left the firing range",
	"8":  "The competitor(%s) entered the penalty laps",
	"9":  "The competitor(%s) left the penalty laps",
	"10": "The competitor(%s) ended the main lap",
	"11": "The competitor(%s) can't continue%s",
	"32": "The competitor(%s) is disqualified",
	"33": "The competitor(%s) has finished",
}

func TrackCompetitors(records []models.Event) {
	models.CompetitorInfo = AddCompetitorsInSlice(records)

	for i := range records {

		competitorID, err := strconv.Atoi(records[i].CompetitorID)
		if err != nil {
			log.Fatal("Error converting CompetitorID:", err)
		}

		if IsDisqualified(records[i]) {
			continue
		}
		ResultOfRace := ""

		switch records[i].ID {
		case "2":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID, records[i].ExtraParam)
			fmt.Println(message)

			ScheduledStart, err := utils.ParseEventTime(records[i].ExtraParam)
			if err != nil {
				log.Fatal("Error converting ScheduledStart:", err)
			}
			models.CompetitorInfo[competitorID-1].ScheduledStart = ScheduledStart

		case "4":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID)
			fmt.Println(message)

			if !IsStratedInTime(records[i]) {
				models.CompetitorInfo[competitorID-1].Disqualified = true
				ResultOfRace = "32"
			}
			startTime, err := utils.ParseEventTime(records[i].Time)
			if err != nil {
				log.Fatal("error parsing start time:", err)
			}
			models.CompetitorInfo[competitorID-1].LapStartTime = startTime
			models.CompetitorInfo[competitorID-1].CurrentLap = 0

		case "5":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID, records[i].ExtraParam)
			fmt.Println(message)
			models.CompetitorInfo[competitorID-1].Shots += 5
		case "6":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].ExtraParam, records[i].CompetitorID)
			fmt.Println(message)
			models.CompetitorInfo[competitorID-1].Hits++
		case "8":
			penaltyStartTime, err := utils.ParseEventTime(records[i].Time)
			if err != nil {
				log.Fatal("error parsing time:", err)
			}
			models.CompetitorInfo[competitorID-1].PenaltyStart = penaltyStartTime

			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID)
			fmt.Println(message)
		case "9":
			penaltyEndTime, err := utils.ParseEventTime(records[i].Time)
			if err != nil {
				log.Fatal("error parsing time:", err)
			}

			if !models.CompetitorInfo[competitorID-1].PenaltyStart.IsZero() {
				duration := penaltyEndTime.Sub(models.CompetitorInfo[competitorID-1].PenaltyStart)
				models.CompetitorInfo[competitorID-1].PenaltyDuration += duration
				models.CompetitorInfo[competitorID-1].PenaltyStart = time.Time{}
			}

			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID)
			fmt.Println(message)
		case "10":
			endTime, err := utils.ParseEventTime(records[i].Time)
			if err != nil {
				log.Fatal("error parsing lap end time:", err)
			}
			lapDuration := endTime.Sub(models.CompetitorInfo[competitorID-1].LapStartTime)

			models.CompetitorInfo[competitorID-1].LapTimes = append(models.CompetitorInfo[competitorID-1].LapTimes,
				utils.FormatDuration(lapDuration))

			models.CompetitorInfo[competitorID-1].LapStartTime = endTime
			models.CompetitorInfo[competitorID-1].CurrentLap++

			models.CompetitorInfo[competitorID-1].DidNotFinish = true

			if models.CompetitorInfo[competitorID-1].CurrentLap == models.Cfg.Laps {
				models.CompetitorInfo[competitorID-1].TotalTime = utils.FormatTotalTime(models.CompetitorInfo[competitorID-1].ActualStart, endTime)
				models.CompetitorInfo[competitorID-1].DidNotFinish = false
				ResultOfRace = "33"
			}
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID)
			fmt.Println(message)
		case "11":
			models.CompetitorInfo[competitorID-1].DidNotFinish = true
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID, ": "+records[i].ExtraParam)
			fmt.Println(message)

			for i := len(models.CompetitorInfo[competitorID-1].LapTimes); i < models.Cfg.Laps; i++ {
				models.CompetitorInfo[competitorID-1].LapTimes = append(models.CompetitorInfo[competitorID-1].LapTimes, "00:00:00.000")
			}

		default:
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID)
			fmt.Println(message)
		}

		switch ResultOfRace {
		case "32":
			message := fmt.Sprintf(EventsList["32"], records[i].CompetitorID)
			fmt.Println(message)
			ResultOfRace = ""
		case "33":
			message := fmt.Sprintf(EventsList["33"], records[i].CompetitorID)
			fmt.Println(message)
			ResultOfRace = ""
		}

	}

}

func CheckCompetitorInSlice(Competitors []models.Competitor, ID string) bool {
	for i := range Competitors {
		if Competitors[i].ID == ID {
			return false
		}
	}
	return true
}

func AddCompetitorsInSlice(records []models.Event) []models.Competitor {
	var CompetitorsInfo []models.Competitor

	for i := range records {
		if CheckCompetitorInSlice(CompetitorsInfo, records[i].CompetitorID) {
			Competitor := models.Competitor{
				ID: records[i].CompetitorID,
			}
			CompetitorsInfo = append(CompetitorsInfo, Competitor)
		}
	}
	return CompetitorsInfo
}

func IsDisqualified(record models.Event) bool {
	competitorID, err := strconv.Atoi(record.CompetitorID)
	if err != nil {
		log.Fatal("Error converting CompetitorID:", err)
	}
	return models.CompetitorInfo[competitorID-1].Disqualified
}
