package controllers

import (
	"biathlon_system/src/models"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFromJSON() models.Config {
	jsonFile, err := os.Open("sunny_5_skiers/config.json")

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var config models.Config

	json.Unmarshal(byteValue, &config)

	fmt.Println(config)
	return config
}

func ReadEvents() []models.Event {
	file, err := os.Open("sunny_5_skiers/events")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully Opened events.txt")
	defer file.Close()

	var records []models.Event
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var record models.Event
		if len(parts) == 3 {
			record = models.Event{
				Time:         parts[0],
				ID:           parts[1],
				CompetitorID: parts[2],
			}

		} else if len(parts) > 3 {
			record = models.Event{
				Time:         parts[0],
				ID:           parts[1],
				CompetitorID: parts[2],
				ExtraParam:   strings.Join(parts[3:], " "),
			}
		}
		records = append(records, record)
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

	// for i := range records {
	// 	fmt.Println(records[i].CompetitorID)
	// }
	return records
}
