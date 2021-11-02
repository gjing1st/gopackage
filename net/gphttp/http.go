// http.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 10:22$

package gphttp

import (
	log "github.com/gjing1st/gopackage/gplog"
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) ([]byte,error){
	resp, err := http.Get(url)
	if err != nil {
		log.LogFile("request","请求",url,"失败,err=",err)
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.LogFile("request","fetch: reading %s: %v\n", url, err)
		return nil, err
	}
	return b,nil
}

