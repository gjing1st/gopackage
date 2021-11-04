// list_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/4$ 8:49$

package user

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestGetDetailList(t *testing.T) {
	token:="4s2mPvDtcEcI20rNtjPBgwbduPNjL31fq3P0-l3_IJpOfHJQodinjZVGSSPMTTSoV35xaKfpc4ifsG3c3vSUvLdnDgxaG5MzzoJxM5ppwbJ7joHTMk51Tt_A1iAv-rOpIljHFdi4Jy2gbaXTzQKIoW6JrIIzND8_Wr7webWuohnOjrUY6TReaN2wzOW075F21yLhKAeWKuD99SwP1sZ5cA"

	res,err := GetDetailList(token,2)
	fmt.Println(res,err)
	g.Dump(res)
}
