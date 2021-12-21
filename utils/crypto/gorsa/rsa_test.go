// rsa_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/11$ 11:03$

package gorsa

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRSA_Encrypt(t *testing.T) {
	//生成密钥对，保存到文件
	//GenerateRSAKey(2048)
	res := "hello"
	//message:=[]byte("hello world")
	message,_:=json.Marshal(res)
	//加密
	cipherText:=RSA_Encrypt(message,"public.pem")
	fmt.Println("加密后为：",string(cipherText))
	//解密
	plainText := RSA_Decrypt(cipherText, "private.pem")
	fmt.Println("解密后为：",string(plainText))
}
