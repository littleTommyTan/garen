package cron

import (
	"github.com/jasonlvhit/gocron"
	"log"
)

func ring() {
	log.Print("goCron ring ~")
}
func Cron() {
	gocron.Every(60).Seconds().Do(ring)
	gocron.Start()

	//gocron.RunAll()
}
