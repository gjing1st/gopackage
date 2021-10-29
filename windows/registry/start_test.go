// $
// Created by dkedTeam.
// User: GJing
// Date: 2021/8/26$ 11:30$

package registry

import (
	"strings"
	"testing"
)

// @description: 测试应用启动
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/8/26 11:34
// 1.应用名称转数组
// 2.查找注册表名称DisplayName是否包含应用名称
// 3.查找安装目录，安装目标不存在时找图标位置
// 4.去除无效字符，返回路径是否带有.exe，不带去找返回路径下对应名称的.exe或者其bin目录下对应名称.exe
// 5.检测文件是否有效
// 6.执行启动
func TestStart(t *testing.T) {
	//names := "腾讯QQ,QQ"
	names := "微信,WeChat"
	//names := "Notepad++"
	//appName := []string{"微信"}
	appName := strings.Split(names, ",")
	StartApp(appName)

}
