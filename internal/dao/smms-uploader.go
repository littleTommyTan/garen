package dao

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

//用于 Upload() func
type SmmsResponce struct {
	Code string `json:"code"`
	Data struct {
		FileName  string `json:"filename"`
		StoreName string `json:"storename"`
		Size      int    `json:"size"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Hash      string `json:"hash"`
		Delete    string `json:"delete"`
		Url       string `json:"url"`
		Path      string `json:"path"`
		Msg       string `json:"msg"`
	} `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
}

func (d *Dao) SmmsUpload(file multipart.File, fileHeader *multipart.FileHeader) (url string, err error) {
	var (
		resp      *http.Response
		bodyBytes []byte
		ret       SmmsResponce
		bodyBuf   = new(bytes.Buffer)
	)
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("smfile", fileHeader.Filename)
	if err != nil {
		return
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return
	}
	defer file.Close()
	_ = bodyWriter.WriteField("ssl", "0")
	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()
	resp, err = http.Post("https://sm.ms/api/upload", contentType, bodyBuf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyBytes, &ret)
	if err != nil {
		log.Print(err)
		return
	}
	if ret.Code == "error" {
		err = errors.New(ret.Msg)
		return
	}
	url = ret.Data.Url
	return
}
