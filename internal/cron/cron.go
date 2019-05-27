package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
)

func ring() {
	logrus.Infof("goCron ring ~")
}
func Cron() {
	gocron.Every(60).Seconds().Do(ring)
	gocron.Start()

	gocron.RunAll()
}
