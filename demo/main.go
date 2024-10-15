package main

import (
	"Demo/source/config"
	//"Demo/source/database"
	"Demo/source/logger"
	"Demo/source/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")
	// database.Init()
	// if config.Appconfig.GetBool("seeddata") {
	// 	//Logic to seed data to database
	// 	logger.InfoLn("Data seeded successfully")
	// }
	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
