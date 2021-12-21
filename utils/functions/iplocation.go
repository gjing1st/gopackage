// iplocation.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/16$ 10:54$

package functions

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall"
	"time"
)

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type QQIpLocation struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Ip       string      `json:"ip"`
		Location GeoLocation `json:"location"`
		AdInfo   struct {
			Nation   string `json:"nation"`
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			AdCode   int    `json:"adcode"`
		} `json:"ad_info"`
	} `json:"result"`
}


// GetLocationFromQQ
// @description: 通过ip获取所在省市区
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/16 10:55
// @return:
func GetLocationFromQQ() (l *QQIpLocation, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 2,
	}
	req, err := http.NewRequest("GET", "https://apis.map.qq.com/ws/location/v1/ip?key=LSNBZ-WT2W3-XC23P-YTTOH-F2D73-W2F4R", nil)
	if err != nil {
		return
	}
	req.Header.Add("Referer", "https://prv.cartesianshield.com/ipquery.service")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("getGeoLocation failed: %s\n", err)
		fmt.Printf("getGeoLocation failed: %s\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("getGeoLocation failed: status: %s\n", resp.Status)
		fmt.Printf("getGeoLocation failed: status: %s\n", resp.Status)
		err = syscall.ENODATA
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	loc := &QQIpLocation{}
	err = json.Unmarshal(body, loc)
	if err == nil {
		l = loc
	}
	return
}
