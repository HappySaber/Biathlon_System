package controllers

import (
	"biathlon_system/src/models"
	"fmt"
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
	CompetitorsInfo := AddCompetitorsInSlice(records)
	for i := range CompetitorsInfo {
		fmt.Println(CompetitorsInfo[i])
	}
	for i := range records {

		switch records[i].ID {
		case "2", "5":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID, records[i].ExtraParam)
			fmt.Println(message)
		case "6":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].ExtraParam, records[i].CompetitorID)
			fmt.Println(message)
		case "11":
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID, ": "+records[i].ExtraParam)
			fmt.Println(message)
		default:
			message := fmt.Sprintf(EventsList[records[i].ID], records[i].CompetitorID)
			fmt.Println(message)
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
