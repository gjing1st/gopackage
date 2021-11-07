//Created by dolitTeam
//@Author : GJing
//@Time : 2020/10/23 13:56
//@File : functions
//@Description: 公共函数库
package functions

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gjing1st/gopackage/net/gphttp"
	"github.com/gogf/gf/util/gconv"
	"io"
	"log"
	"net/url"
)

// Md5
// Author: GJing
// Email: gjing1st@gmail.com
// Date: 2020/10/23 13:57
// Description: md5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

// InArray
// @description: 判断一个字符串是否在数组中
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2020/11/13 14:42
// @success:
func InArray(value string, arr []string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// ReserveNumber
// @description: 截取保留小数点后m位，舍去后面位数
// @param: i 接口，int,float,string类型
// @param: m 保留的位数
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2020/11/25 上午 10:31
// @success: 返回截取后的字符串
func ReserveNumber(i interface{}, m int) string {
	s := gconv.String(i)
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i+m+1]
		}
	}
	return s
}

type AuthRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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
	return authRes, err
}
