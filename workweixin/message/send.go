// send.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 14:19$

package message

import (
	"log"
	"net/url"
)

type FormParams struct {
	Touser string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
	Btntxt string `json:"btntxt"`
	EnableUdTrans string `json:"enable_id_trans"`
	EnableDuplicateCheck string `json:"enable_duplicate_check"`
	DuplicateCheckInterval string `json:"duplicate_check_interval"`
}

type TextCard struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
}


func reqUrl(accessToken string)  string {
	params := url.Values{}
	parseURL, err := url.Parse(workweixin.BaseUrl + "/message/send")
	if err != nil {
		log.Println("err")
		return ""
	}
	params.Set("access_token", accessToken)

	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	return urlPathWithParams
}

func send(url string)  {

}