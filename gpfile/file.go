// file.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/10/29$ 16:39$

package gpfile

import (
	"fmt"
	"gitee.com/gjing1st/gopackage/utils/functions"
	"github.com/gogf/gf/os/gfile"
	"path/filepath"
	"strings"
)

// IsAllVideoFile
// @description: 该路径下如果有文件是否都是视频文件(包含递归子目录)
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2020/11/12 16:37
// @success: 全是视频文件返回true，否则返回false
func IsAllVideoFile(sourcePath string) bool {
	var VideoType = []string{"MP4", "AVI", "WMV", "MPEG", "QuickTime", "RealVideo", "Flash", "Mpeg-4", "MPG", "DAT", "MOV", "FLV"}
	var ok = true
	names, _ := gfile.DirNames(sourcePath)
	if len(names) > 0 {
		for _, v := range names {
			if gfile.IsDir(sourcePath + "/" + v) {
				b := IsAllVideoFile(sourcePath + "/" + v)
				if b == false {
					return false
				}
			} else {
				ext := strings.TrimLeft(filepath.Ext(v), ".")
				if ext != "" && !functions.InArray(ext, VideoType) {
					return false
				}
			}
		}
	}
	return ok
}

// IsEmptyCatalog
// @description: 检查给定目录下是否为空文件(递归子目录)
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2020/11/12 17:12
// @success: 如果没有文件，它将返回true
func IsEmptyCatalog(path string) bool {
	if gfile.IsFile(path) {
		return false
	}
	var ok = true
	names, _ := gfile.DirNames(path) //目录下所有文件
	fmt.Println("所有文件：", names)
	if len(names) > 0 {
		for _, v := range names {
			isCatalog := gfile.IsFile(path + "/" + v) //是否是文件
			fmt.Println("是否文件", v, isCatalog)
			if isCatalog {
				return false
			} else {
				//有子目录
				b := IsEmptyCatalog(path + "/" + v)
				if !b {
					return false
				}
			}

		}
	}
	return ok
}


