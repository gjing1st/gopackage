// Created by dolitTeam
//@Author : GJing
//@Time : 2020/10/23 11:46
//@File : log
package gplog

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"time"
)

// 用于应用初始化。
func initLog() {
	timeString := time.Now()
	//path := g.Cfg().GetString("logger.Path") + "/" + timeString.Format("2006-01") + "/" + timeString.Format("2006-01-02")
	path :=  "./logs/" + timeString.Format("2006-01") + "/" + timeString.Format("2006-01-02")
	//设置日志路径，自动创建目录
	glog.SetPath(path)
	//开启异步日志记录
	glog.SetAsync(true)
	//关闭控制台输出
	glog.SetStdoutPrint(false)
}

// LogFile
// @description: 日志记录到文件
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2020/10/23 11:35
// @return:
func LogFile(fileName string, v ...interface{}) {
	initLog()
	//使用对象设置方法，高并发时容易导致日志写入其他文件，改用链式操作
	//glog.SetFile(fileName)
	pathArr := strings.Split(fileName, "/")
	//字符串切割，如果是路由则取第一个路径为文件名
	if len(pathArr) > 1 {
		fileName = pathArr[1]
	}
	//如果文件名为空，则默认使用common
	if len(fileName) == 0 {
		fileName = "common.log"
	}

	//使用回溯值记录调用日志文件名和行号
	glog.Skip(1).Line(true).File(fileName).Println(v)
}

// Println
// @description: 兼容官方log包
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 11:33
// @return:
func Println(v ...interface{})  {
	fileName := "common.log"
	if len(v)>0{
		fileName = gconv.String(v[0])
	}
	LogFile(fileName,v)
}
