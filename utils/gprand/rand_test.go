// rand_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/11$ 14:10$

package gprand

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	s := Str("1234567890",1)
	fmt.Println(s)
	n := Intn(2)
	fmt.Println(n)
}
