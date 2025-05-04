package main

import "biathlon_system/src/controllers"

func main() {
	controllers.ReadFromJSON()
	//controllers.ReadEvents()
	controllers.TrackCompetitors(controllers.ReadEvents())
}
