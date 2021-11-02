// send.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 14:19$

package message

import (
	"log"
	"net/url"
	"github.com/gjing1st/gopackage/workweixin"
)

type BaseParams struct {
	Touser                 string `json:"touser"`
	Toparty                string `json:"toparty"`
	Totag                  string `json:"totag"`
	Msgtype                string `json:"msgtype"`
	Agentid                int    `json:"agentid"`
	EnableUdTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}

// @description: 发送消息url
// @param: accessToken
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 22:48
// @success: url
func getRequestUrl(accessToken string) string {
	params := url.Values{}
	parseURL, err := url.Parse(workweixin.BaseUrl + "/message/send")
	if err != nil {
		log.Println("err")
		return ""
	}
	params.Set("access_token", accessToken)

	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	return urlPathWithParams
}
