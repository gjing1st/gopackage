// list.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/4$ 8:37$

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

type DetailList struct {
	SimpleUser
	Position string `json:"position"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type DetailListResp struct {
	workweixin.BaseResponse
	UserList []DetailList `json:"userlist"`
}

// ReqDetailList
// @description: 请求企业微信获取部门成员详细
// @param: accessToken 通讯录的access_token
// @param: depId 部门id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 13:55
// @return:
func ReqDetailList(accessToken string, depId int) ([]byte, error) {
	params := url.Values{}
	parseURL, err := url.Parse(workweixin.BaseUrl + "/user/list")
	if err != nil {
		log.Println("err")
		return nil, err
	}
	params.Set("access_token", accessToken)
	if depId > 0 {
		params.Set("department_id", strconv.Itoa(depId))
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	return gphttp.GetRequest(urlPathWithParams)
}

// GetDetailListRes
// @description: 部门成员详细信息
// @param:  res 企业微信返回的字节数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 13:54
// @return:
func GetDetailListRes(res []byte) ([]DetailList, error) {
	listRes := DetailListResp{}
	err := json.Unmarshal(res, &listRes)
	if err != nil {
		log.Println("user", "SimpleListResponse转换失败", err)
	}
	if listRes.Errcode != 0 {
		return nil, errors.New(listRes.Errmsg)
	}
	return listRes.UserList, nil
}

// GetDetailList
// @description: 获取部门成员详细信息
// @param: accessToken 通讯录的access_token
// @param: depId 部门id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 13:54
// @return: 部门成员列表
func GetDetailList(accessToken string, depId int) ([]DetailList, error) {
	res, err := ReqDetailList(accessToken, depId)
	if err != nil {
		return nil, err
	}
	return GetDetailListRes(res)
}