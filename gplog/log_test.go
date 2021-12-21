// log_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/22$ 15:18$

package gplog

import (
	"testing"
	"time"
)

func TestPrintlnDay(t *testing.T) {
	//LogPath = "D:/tmp/logs"
	//PrintlnPath("path1.log","path1")
	//PrintlnDay("test.log")
	//PrintlnDay("test.log","1111","2222222")
	//PrintlnMonth("test1.log")
	PrintlnPath("path2","path2")
	time.Sleep(time.Second)
}
