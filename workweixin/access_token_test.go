// access_token_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 10:58$

package workweixin

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetToken(t *testing.T) {
	//b, err := workweixin.ReqToken("ww887a89a56b64260f", "65nTxxqhCmVmiq151HFs4M_lRSBpBOp61Mw2IA1XCUQ")
	//fmt.Println(1111, b)
	//fmt.Println(22222, err)
	mqttAddr := "tc:///"  + ":1883"
	u,e :=url.Parse(mqttAddr)
	fmt.Println(u)
	fmt.Println(e)
}

