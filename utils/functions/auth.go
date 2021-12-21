package functions

import (
	"encoding/json"
	log "gitee.com/gjing1st/gopackage/gplog"
	"gitee.com/gjing1st/gopackage/net/gphttp"
	"net/url"
)

var BaseUrl = "http://auth.zdhr.top/auth/"

type AuthRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Auth
// @description: 请求授权中心
// @param: sys 系统名称
// @param: token 分配的授权token
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/7 12:47
// @success:
func Auth(sys, token string) (AuthRes, error) {
	params := url.Values{}
	parseURL, err := url.Parse(BaseUrl + sys)
	if err != nil {
		log.Println("err")
		return AuthRes{}, err
	}
	params.Set("token", token)
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resBytes, err := gphttp.GetRequest(urlPathWithParams)
	if err != nil {
		return AuthRes{}, err
	}
	authRes := AuthRes{}
	err = json.Unmarshal(resBytes, &authRes)
	//if  authRes.Data != ""{
	//	authRes.Data = string(RSA_Decrypt([]byte(authRes.Data),"private.pem"))
	//}
	return authRes, err
}
