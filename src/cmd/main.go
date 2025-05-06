package main

import (
	"biathlon_system/src/controllers"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Do you want to check singleEvent or multiEvent (1/2)")
	choose := ""
	fmt.Fscan(os.Stdin, &choose)
	switch choose {
	case "1":
		controllers.ReadFromJSON("singleEvent/config.json")
		controllers.TrackCompetitors(controllers.ReadEvents("singleEvent/events"))
		controllers.GenerateReport()
	case "2":
		controllers.ReadFromJSON("multiEvent/config.json")
		controllers.TrackCompetitors(controllers.ReadEvents("multiEvent/events"))
		controllers.GenerateReport()
	default:
		fmt.Println("Wrong value")
	}
}
