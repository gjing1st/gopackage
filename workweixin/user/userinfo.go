// userinfo.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/8$ 16:31$

package user

import (
	"encoding/json"
	"errors"
	log "gitee.com/gjing1st/gopackage/gplog"
	"gitee.com/gjing1st/gopackage/net/gphttp"
	"gitee.com/gjing1st/gopackage/workweixin"
	"net/url"
)

type UserInfoResp struct {
	workweixin.BaseResponse
	UserId   string `json:"UserId"`
	DeviceId string `json:"DeviceId"`
}

// ReqUserInfoByCode
// @description: 通过企业微信授权返回的code获取用户信息字节
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/8 16:40
// @return:
func ReqUserInfoByCode(accessToken string, code string) ([]byte, error) {
	params := url.Values{}
	parseURL, err := url.Parse(workweixin.BaseUrl + "/user/getuserinfo")
	if err != nil {
		log.Println("err")
		return nil, err
	}
	params.Set("access_token", accessToken)
	params.Set("code", code)
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	return gphttp.GetRequest(urlPathWithParams)
}

// ReqUserInfoByCodeRes
// @description: 通过企业微信授权返回的字节数组获取用户信息
// @param:  res 企业微信返回的字节数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/9 9:02
// @return:
func ReqUserInfoByCodeRes(res []byte) (UserInfoResp, error) {
	listRes := UserInfoResp{}
	err := json.Unmarshal(res, &listRes)
	if err != nil {
		log.Println("user", "SimpleListResponse转换失败", err)
	}
	if listRes.Errcode != 0 {
		return UserInfoResp{}, errors.New(listRes.Errmsg)
	}
	return listRes, nil
}

// GetUserInfoByCode
// @description: 通过企业微信授权返回的code获取用户信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/9 9:03
// @return:
func GetUserInfoByCode(accessToken string, code string) (UserInfoResp, error)  {
	resBytes,err := ReqUserInfoByCode(accessToken,code)
	if err != nil {
		return UserInfoResp{}, err
	}
	return ReqUserInfoByCodeRes(resBytes)

}
