package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
)

func ring() {
	logrus.Infof("goCron ring ~")
}

// Cron sheduler
func Cron() {
	gocron.Every(300).Seconds().Do(ring)
	gocron.Start()

	gocron.RunAll()
}
