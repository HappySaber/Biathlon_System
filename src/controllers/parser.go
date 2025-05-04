package controllers

import (
	"biathlon_system/src/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReaderFromJSON() {
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
}
