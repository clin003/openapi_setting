package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gitee.com/lyhuilin/util"
)

func postUserInfo(req interface{}, apiURL string) (retText string, err error) {
	dataBody := util.JsonEncode(req)

	rep, err := http.NewRequest("POST", apiURL, strings.NewReader(dataBody))
	if err != nil {
		return
	}
	rep.Header.Add("Content-Type", "application/json")
	// rep.Header.Add("HLTYClient", "haowu_video")

	client := &http.Client{Timeout: time.Second * 30}

	resp, err := client.Do(rep)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(body) > 0 {
		retText = string(body)
	}
	// log.Println("updateHaowuVideoStat100ToMyAdmin()", "retText", retText)
	return
}

func RegUserInfo(req interface{}) (retText string, err error) {
	apiUrl := "https://vip.lyhuilin.com/v1/user"
	return postUserInfo(req, apiUrl)
}
func LoginUserInfo(req interface{}) (retText string, err error) {
	apiUrl := "https://vip.lyhuilin.com/login"
	return postUserInfo(req, apiUrl)
}

type ResponseToken struct {
	Code int64 `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
}

// 返回Token
func RegUserInfoEx(req interface{}) (retText string, err error) {
	respText, err := RegUserInfo(req)
	if err != nil {
		return
	}
	var t ResponseToken
	if err1 := util.JsonDecode(respText, &t); err1 != nil {
		err = fmt.Errorf("%s:%v", respText, err1)
		return
	}
	if len(t.Data.Token) > 0 {
		retText = t.Data.Token
	} else {
		err = fmt.Errorf("%s", retText)
	}
	return
}

// 返回Token
func LoginUserInfoEx(req interface{}) (retText string, err error) {
	respText, err := LoginUserInfo(req)
	if err != nil {
		return
	}
	var t ResponseToken
	if err1 := util.JsonDecode(respText, &t); err1 != nil {
		err = fmt.Errorf("%s:%v", respText, err1)
		return
	}
	if len(t.Data.Token) > 0 {
		retText = t.Data.Token
	} else {
		err = fmt.Errorf("%s", retText)
	}
	return
}
