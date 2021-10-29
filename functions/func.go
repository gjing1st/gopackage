//Created by dolitTeam
//@Author : GJing
//@Time : 2020/10/23 13:56
//@File : functions
//@Description: 公共函数库
package functions

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gogf/gf/util/gconv"
	"io"
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
