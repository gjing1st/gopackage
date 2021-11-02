// access_token_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 10:58$

package workweixin

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	b,err := GetToken("ww887a89a56b64260f","65nTxxqhCmVmiq151HFs4M_lRSBpBOp61Mw2IA1XCUQ")
	fmt.Println(1111,b)
	fmt.Println(22222,err)
}
