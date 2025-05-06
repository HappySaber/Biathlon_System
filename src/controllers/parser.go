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

func ReadFromJSON(fileName string) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened config.json")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var config models.Config

	json.Unmarshal(byteValue, &config)

	models.Cfg = config
}

func ReadEvents(fileName string) []models.Event {
	file, err := os.Open(fileName)
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
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

	return records
}
