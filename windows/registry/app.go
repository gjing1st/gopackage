// $ 通过注册表查询应用启动
// Created by dkedTeam.
// User: GJing
// Date: 2021/8/26$ 11:28$

package registry

import (
	registry "golang.org/x/sys/windows/registry"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// 1.应用名称转数组
// 2.查找注册表名称DisplayName是否包含应用名称
// 3.查找安装目录，安装目标不存在时找图标位置
// 4.去除无效字符，返回路径是否带有.exe，不带去找返回路径下对应名称的.exe或者其bin目录下对应名称.exe
// 5.检测文件是否有效
// 6.执行启动
var path = `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`

// StartApp
// @description: 启动App
// @param: appName 应用名称，同一个应用，多种叫法数组
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/8/30 13:53
// @success:
func StartApp(appName []string) {
	path := findAppPath(appName)
	//返回路径中带有"需要去除"
	path = strings.ReplaceAll(path, `"`, "")

	fileExt := filepath.Ext(path)
	if fileExt == "" {
		path = findFinallyPath(appName, path)
	}
	//文件是否存在
	pathBool := pathExists(path)
	if !pathBool {
		log.Println("路径不存在")
		return
	}

	cmd := exec.Command("cmd", "/C", path)
	log.Println("cmd===", cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal("---", err)
	}
	defer stdout.Close()                // 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}
}

// @description: 查找应用程序的启动路径
// @param: appName 应用名词
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/8/24 17:20
// @success: 返回路径字符串，失败返回空字符串
func findAppPath(appName []string) string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.READ)
	var appPath string
	if err != nil {
		log.Println("err==", err)
		return appPath
	}
	defer k.Close()
	// 读取：一个项下的所有子项
	keys, _ := k.ReadSubKeyNames(0)
	appPathCh := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(len(keys))
	for _, subKey := range keys {
		//detailPath := path + `\` + subKey
		//k, _ := registry.OpenKey(registry.LOCAL_MACHINE, detailPath, registry.READ)
		//name, _, _ := k.GetStringValue("DisplayName")
		//if name == appName {
		//	displayPath, _, err := k.GetStringValue("DisplayIcon")
		//	if err != nil {
		//		log.Println("app DisplayIcon not found", err)
		//		appPath, _, _ = k.GetStringValue("InstallLocation")
		//	}else {
		//		appPath = displayPath
		//	}
		//}

		go findMatchAppName(appName, subKey, appPathCh, &wg)
	}
	appPath = <-appPathCh
	go func(appPathCh chan string) {
		wg.Wait()
		close(appPathCh)
	}(appPathCh)
	//有些应用返回的路径带有,0
	finallyPath := strings.Split(appPath, ",")
	return finallyPath[0]
}

// @description: 查找匹配的应用名称
// @param: appName 应用名数组
// @param: subKey 注册表中每项的值
// @param: appPathCh 应用路径管道
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/8/24 17:47
// @success: 放入管道
func findMatchAppName(appName []string, subKey string, appPathCh chan string, wg *sync.WaitGroup) {
	var appPath string
	detailPath := path + `\` + subKey
	k, _ := registry.OpenKey(registry.LOCAL_MACHINE, detailPath, registry.READ)
	defer k.Close()
	defer wg.Done()
	displayName, _, _ := k.GetStringValue("DisplayName")
	var matchBool bool
	for _, val := range appName {
		//if val == displayName {
		//改为使用字符串包含，有些displayName实在太奇葩
		if strings.Contains(displayName, val) {
			matchBool = true
		}
	}
	if matchBool {
		// TODO 图标路径有些为卸载程序
		//displayPath, _, err := k.GetStringValue("DisplayIcon")
		// 改为先找安装路径
		installLocation, _, err := k.GetStringValue("InstallLocation")
		if err != nil {
			log.Println("app DisplayIcon not found", err)
			//installLocation, _, _ := k.GetStringValue("InstallLocation")
			displayPath, _, _ := k.GetStringValue("DisplayIcon")
			//appPath = installLocation + "bin" + `\` + displayName + ".exe"
			//appPath = installLocation
			appPath = displayPath
		} else {
			//appPath = displayPath
			appPath = installLocation
		}
	}
	if appPath != "" {
		appPathCh <- appPath
		//close(appPathCh)
	}

}

// @description: 查找应用最终路径
// @param: name 应用名称数组
// @param: path 应用上级目录
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/8/25 10:36
// @success: 文件的完整路径否则返回空字符串
func findFinallyPath(name []string, path string) string {
	pathLen := len(path)
	// 有些注册表的安装位置最后没有加\，此处补上
	if path[pathLen-1] != '\\' {
		path = path + `\`
	}
	// TODO 可尝试找目录下的所有exe，取名称对比看有没有是合适的程序
	for _, v := range name {
		//路径中不含文件名
		newPath := path + v + ".exe"
		pathBool := pathExists(newPath)
		if pathBool {
			return newPath
		}
	}
	//同级目录未找到对应应用
	for _, v := range name {
		//路径中不含文件名
		newPath := path + "bin" + `\` + v + ".exe"
		pathBool := pathExists(newPath)
		if pathBool {
			return newPath
		}
	}
	return ""

}

// @description: 文件或者目录是否存在
// @param: path 绝对路径
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/8/25 9:59
// @success: 存在返回true，不存在false
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
