// sms.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/12$ 10:03$

package gojuhe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendSms(mobile,tplId,key,tplValue string) {
	//请求地址
	juheURL := "http://v.juhe.cn/sms/send"

	//初始化参数
	param := url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("mobile", mobile) //接受短信的用户手机号码
	param.Set("tpl_id", tplId) //您申请的短信模板ID，根据实际情况修改
	if tplValue !=""{
		param.Set("tpl_value", "#code#=1234&#company#=聚合数据") //您设置的模板变量，根据实际情况
	}
	param.Set("key", key)  //应用APPKEY(应用详细页查询)

	//发送请求
	data, err := Post(juheURL, param)
	if err != nil {
		fmt.Errorf("请求异常,错误信息:\r\n%v", err)
	} else {
		var netReturn map[string]interface{}
		json.Unmarshal(data, &netReturn)

		fmt.Println(netReturn)
	}

}

// Get 网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values) (rs []byte, err error) {
	fmt.Println("apiURL:",apiURL)
	fmt.Println("params:",params)
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

