package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/tommytan/garen/internal/app/wechat"
)

func Cron() {
	gocron.Every(7000).Seconds().Do(wechat.FetchAccessToken)
	gocron.Start()

	gocron.RunAll()
}
