// department.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 12:34$

package department

import (
	"encoding/json"
	"errors"
	"github.com/gjing1st/gopackage/workweixin"
	"github.com/gjing1st/gopackage/net/gphttp"
	"log"
	"net/url"
	"strconv"
)

type Department struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	Parentid string `json:"parentid"`
	Order    int    `json:"order"`
}
type DepartmentListResponse struct {
	workweixin.BaseResponse
	Department []Department `json:"department"`
}

// ReqList
// @description: 请求企业微信部门列表
// @param: accessToken 通讯录的access_token
// @param: depId 部门id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 12:43
// @return: 
func ReqList(accessToken string, depId int) ([]byte, error) {
	params := url.Values{}
	parseURL, err := url.Parse(workweixin.BaseUrl + "/department/list")
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

// GetListRes
// @description: 部门列表
// @param: res 企业微信返回的字节数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 12:53
// @return: 
func GetListRes(res []byte) ([]Department, error){
	listRes := DepartmentListResponse{}
	err := json.Unmarshal(res,&listRes)
	if err != nil {
		log.Println("depart", "departmentList转换失败", err)
	}
	if listRes.Errcode != 0 {
		return nil, errors.New(listRes.Errmsg)
	}
	return listRes.Department, nil
}

// GetList
// @description: 获取部门列表
// @param: accessToken 通讯录的access_token
// @param: depId 部门id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 12:57
// @return: 部门列表
func GetList(accessToken string, depId int) ([]Department, error) {
	res,err := ReqList(accessToken,depId)
	if err != nil{
		return nil, err
	}
	return GetListRes(res)
}