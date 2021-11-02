// access_token.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 10:39$

package workweixin

import (
	"log"
	"net/url"
	"github.com/gjing1st/gopackage/net/gphttp"
)

var baseUrl = "https://qyapi.weixin.qq.com/cgi-bin"

// GetToken
// @description: 获取应用的access_token
// @param: corpid 企业ID
// @param: corpsecret 应用的凭证密钥
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 10:54
// @return: 
func GetToken(corpid, corpsecret string) ([]byte, error) {
	params := url.Values{}
	parseURL, err := url.Parse(baseUrl + "/gettoken")
	if err != nil {
		log.Println("err")
		return nil, err
	}
	params.Set("corpid", corpid)
	params.Set("corpsecret", corpsecret)
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	return gphttp(urlPathWithParams)
	//resp, err := http.Get(urlPathWithParams)
	//if err != nil {
	//	log.Println("err")
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println("err")
	//	return nil, err
	//}
	//return b, nil
}
