package cron

import (
"github.com/jasonlvhit/gocron"
)

// Cron sheduler
func Cron() {
	gocron.Every(1).Day().At("8:00").Do(weather)
	gocron.Every(1).Day().At("10:00").Do(hi)
	gocron.Every(1).Day().At("12:00").Do(xinkulaNoon)
	gocron.Every(1).Day().At("13:30").Do(buyaoshuila)
	gocron.Every(1).Day().At("16:00").Do(xiaoxie)
	gocron.Every(1).Day().At("18:30").Do(xinkulaAfterWork)
	weather()
	gocron.Start()
}
