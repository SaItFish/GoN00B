// Package main
// @file: main.go
// @date: 2021/1/3
package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	appID    = 11111
	appKey   = "xxxxxx"
	fromLang = "auto"
	toLang   = "zh"
	baiduURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

//计算文本的md5值
func makeMD5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}

type TranslateModel struct {
	Q     string
	From  string
	To    string
	AppID int
	Salt  int
	Sign  string
}

type TranslateResponse struct {
	From string `json:"from"`
	To   string `json:"to"`

	TransResult []map[string]string `json:"trans_result"`
}

func NewTranslateModel(q, from, to string) TranslateModel {
	model := TranslateModel{
		Q:     q,
		From:  from,
		To:    to,
		AppID: appID,
	}
	model.Salt = time.Now().Second()
	content := strconv.Itoa(appID) + q + strconv.Itoa(model.Salt) + appKey
	model.Sign = makeMD5(content)
	return model
}

func (tran TranslateModel) Body() map[string]string {
	return map[string]string{
		"q":     tran.Q,
		"from":  tran.From,
		"to":    tran.To,
		"appid": strconv.Itoa(tran.AppID),
		"salt":  strconv.Itoa(tran.Salt),
		"sign":  tran.Sign,
	}
}

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}

	inputData := string(buf)
	inputData = strings.ReplaceAll(inputData, "\n\n", "#saltfish#")
	inputData = strings.ReplaceAll(inputData, "\n", " ")
	inputData = strings.ReplaceAll(inputData, "#saltfish#", "\n\n")

	model := NewTranslateModel(inputData, fromLang, toLang)
	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
		Data:    model.Body(),
	}
	resp, err := grequests.Post(baiduURL, ro)
	if err != nil {
		log.Println("Cannot post: ", err)
	}
	if resp.Ok != true {
		log.Println("Request did not return OK")
	}

	var data TranslateResponse
	if err := json.Unmarshal([]byte(resp.String()), &data); err != nil {
		log.Fatal(err)
	}
	res := data.TransResult

	var builder strings.Builder
	for i := 0; i < len(res); i++ {
		builder.WriteString(res[i]["dst"])
		builder.WriteString("\n")
	}

	err = ioutil.WriteFile(outputFile, []byte(builder.String()), 0644) // oct, not hex
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
}

/*
一次最多翻译30段
TODO: 每30段分成一个请求，每个1.5s请求1次（百度API普通用户1s只能请求1次）
*/
