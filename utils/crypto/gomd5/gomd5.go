// gomd5.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/11$ 10:38$

package gomd5

import (
	"crypto/md5"
	"fmt"
)

// EncryptBytes
// @description:
// @param: data 要加密的字节数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/11 10:42
// @return:
func EncryptBytes(data []byte) (encrypt string, err error) {
	h := md5.New()
	if _, err = h.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// EncryptByString
// @description:
// @param: data 要加密的字符串
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/11 10:42
// @return:
func EncryptByString(data string) (string,error) {
	return EncryptBytes([]byte(data))
}