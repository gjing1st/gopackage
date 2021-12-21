package goauth

import (
	"encoding/json"
	log "github.com/gjing1st/gopackage/gplog"
	"github.com/gjing1st/gopackage/net/gphttp"
	"net/url"
)

type AuthRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Auth
// @description: 请求授权中心
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/7 12:47
// @success:
func Auth(sys, token string) (AuthRes, error) {
	//BaseUrl := "http://127.0.0.1:8199/auth/"
	BaseUrl := "http://auth.zdhr.top/auth/"
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
