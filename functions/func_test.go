package functions

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	res,err := Auth("visit","2178")
	fmt.Println("res.Code=",res.Code)
	fmt.Println("err = ",err)
	fmt.Println(res.Message,err)
}

func TestRSA_Encrypt(t *testing.T) {
	//生成密钥对，保存到文件
	//GenerateRSAKey(2048)
	res := AuthRes{
		0,
		"ok",
		"",
	}
	//message:=[]byte("hello world")
	message,_:=json.Marshal(res)
	//加密
	cipherText:=RSA_Encrypt(message,"public.pem")
	fmt.Println("加密后为：",string(cipherText))
	//解密
	plainText := RSA_Decrypt(cipherText, "private.pem")
	fmt.Println("解密后为：",string(plainText))
}