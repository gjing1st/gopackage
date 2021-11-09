// list_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/4$ 8:49$

package user

import (
	"fmt"
	"testing"
)

func TestGetDetailList(t *testing.T) {
	token:="5j6WOn2hG1rjV2zIiIw83W0MeQEe8mBbbEbZSpMxOM0KgdxkfptdVqgM-xiPHZK6obFUEDZ1wsDAKuJzhIaLmk16SMqiFuVcqD4QvyoR5JoybGRHpL0ZwFAHz9gSGuzzdGDuV5Omknaa-6ENcN1o_6cr4ryMmhqRzeyTvF8Ngn41vMzMkCYZXhnkrFPdUB63iLWpUijcO5jYU9ACDH4WFA"

	//res,err := GetDetailList(token,2)
	//fmt.Println(res,err)
	//g.Dump(res)
	r,err := GetUserInfoByCode(token,"123")
	fmt.Println(r,err)
}
