// auth_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/11$ 11:24$

package goauth

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	s,err := Auth("visit","17omqnu4s40cfj7zt38kcuk100cn2178")
	fmt.Println("s",s)
	fmt.Println("err",err)
}
