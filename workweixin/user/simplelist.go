// simplelist.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 13:45$

package user

import (
	"encoding/json"
	"errors"
	"github.com/gjing1st/gopackage/net/gphttp"
	"github.com/gjing1st/gopackage/workweixin"
	"log"
	"net/url"
	"strconv"
)

// SimpleUser
// 简要用户信息
type SimpleUser struct {
	UserId     int    `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	OpenUserid string `json:"open_userid"`
}

// SimpleListResponse
// 简要部门成员
type SimpleListResponse struct {
	workweixin.BaseResponse
	UserList []SimpleUser `json:"userlist"`
}

// ReqSimpleList
// @description: 请求企业微信获取部门成员
// @param: accessToken 通讯录的access_token
// @param: depId 部门id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 13:55
// @return:
func ReqSimpleList(accessToken string, depId int) ([]byte, error) {
	params := url.Values{}
	parseURL, err := url.Parse(workweixin.BaseUrl + "/user/simplelist")
	if err != nil {
		log.Println("err")
		return nil, err
	}
	params.Set("access_token", accessToken)
	if depId > 0 {
		params.Set("id", strconv.Itoa(depId))
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	return gphttp.GetRequest(urlPathWithParams)
}

// GetSimpleListRes
// @description: 部门成员
// @param:  res 企业微信返回的字节数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 13:54
// @return:
func GetSimpleListRes(res []byte) ([]SimpleUser, error) {
	listRes := SimpleListResponse{}
	err := json.Unmarshal(res, &listRes)
	if err != nil {
		log.Println("user", "SimpleListResponse转换失败", err)
	}
	if listRes.Errcode != 0 {
		return nil, errors.New(listRes.Errmsg)
	}
	return listRes.UserList, nil
}

// GetSimpleList
// @description: 获取部门成员
// @param: accessToken 通讯录的access_token
// @param: depId 部门id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 13:54
// @return: 部门成员列表
func GetSimpleList(accessToken string, depId int) ([]SimpleUser, error) {
	res, err := ReqSimpleList(accessToken, depId)
	if err != nil {
		return nil, err
	}
	return GetSimpleListRes(res)
}
