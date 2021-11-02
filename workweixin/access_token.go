// access_token.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 10:39$

package workweixin

import (
	"encoding/json"
	"errors"
	log "github.com/gjing1st/gopackage/gplog"
	"github.com/gjing1st/gopackage/net/gphttp"
	"net/url"
)

var baseUrl = "https://qyapi.weixin.qq.com/cgi-bin"

// AccessTokenRes
// 企业微信返回的数据
type AccessTokenRes struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

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
	return gphttp.GetRequest(urlPathWithParams)
}

// GetTokenRes
// @description: access_token转换
// @param: res返回的字符数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 11:19
// @return: access_token值
func GetTokenRes(res []byte) (string, error) {
	tokenRes := AccessTokenRes{}
	err := json.Unmarshal(res, &tokenRes)
	if err != nil {
		log.Println("token", "access_token转换失败", err)
	}
	if tokenRes.Errcode != 0 {
		return "", errors.New(tokenRes.Errmsg)
	}
	return tokenRes.AccessToken, nil
}

// GetAccessToken
// @description: 完整版直接获取token函数
// @param: corpid 企业ID
// @param: corpsecret 应用的凭证密钥
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 12:24
// @return: access_token值
func GetAccessToken(corpid, corpsecret string) (string, error) {
	res, _ := GetToken(corpid, corpsecret)
	return GetTokenRes(res)
}
