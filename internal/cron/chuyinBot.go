package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func httpPostJson(url string, jsonStr string) {
	jsonByte := []byte(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		logrus.Fatalf(`req built error: `, err)
		// handle error
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf(`resp return err: `, err)
		// handle error
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	logrus.Infof(`%s %s %s`, string(body), statuscode, hea)
}

func getWeather() (map[string]interface{}, error) {
	resp, err := http.Get(shanghaiWeatherUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		logrus.Error("json marshal failed,err:", err)
		return nil, err
	}
	todayData := (((data[`data`].(map[string]interface{}))["forecast"]).([]interface{})[0]).(map[string]interface{})
	return todayData, nil
}

const shanghaiWeatherUrl = `http://t.weather.sojson.com/api/weather/city/101020100`
const chuyinUrl = `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=e3bd4082-db66-4479-9117-b9c68d7623ec`

func weather() {
	todayData, err := getWeather()
	if err != nil {
		logrus.Error(`fetch weather failed`)
	}
	data, err := json.Marshal(todayData)
	if err != nil {
		logrus.Errorf(`fetch weather failed`, err)
	}

	var postData = gin.H{
		"msgtype": "text",
		"text": gin.H{
			"content": fmt.Sprintf(`今天上海市天气预报：%s%s，最高温：%s，最低温：%s，风向：%s，风速:%s，日出：%s，日落：%s，温馨提示：%s`,
				todayData[`ymd`], todayData[`week`], todayData[`high`], todayData[`low`], todayData[`fx`], todayData[`fl`], todayData[`sunrise`], todayData[`sunset`], todayData[`notice`]),
		},
	}
	data, err = json.Marshal(postData)
	if err != nil {
		return
	}
	httpPostJson(chuyinUrl, string(data))
}

func hi() {
	var postData = gin.H{
		"msgtype": "text",
		"text": gin.H{
			"content": "新的一天来啦，工作也要元气满满呢～",
		},
	}
	data, err := json.Marshal(postData)
	if err != nil {
		return
	}
	httpPostJson(chuyinUrl, string(data))
}

func ping() {
	var postData = gin.H{
		"msgtype": "text",
		"text": gin.H{
			"content": "配置成功",
		},
	}
	data, err := json.Marshal(postData)
	if err != nil {
		return
	}
	httpPostJson(chuyinUrl, string(data))
}

func xinkulaNoon() {
	var postData = gin.H{
		"msgtype": "text",
		"text": gin.H{
			"content": "大家辛苦了，去吃午饭吧～",
		},
	}
	data, err := json.Marshal(postData)
	if err != nil {
		return
	}
	httpPostJson(chuyinUrl, string(data))
}

func xinkulaAfterWork() {
	var postData = gin.H{
		"msgtype": "text",
		"text": gin.H{
			"content": "大家辛苦了，准备下班吧～",
		},
	}
	data, err := json.Marshal(postData)
	if err != nil {
		return
	}
	httpPostJson(chuyinUrl, string(data))
}
