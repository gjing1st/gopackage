// mqtt.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/16$ 9:36$

package gpmqtt

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"syscall"
	"time"
)

const (
	MQTT_WILL = "{\"action\":\"offline\"}"
)

var (
	mqttClient  MQTT.Client
	mqttMutex   sync.Mutex
	statusTopic string
	noticeTopic string
)

type MqttNotice struct {
	Action string `json:"action"`
}

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type AddressInfo struct {
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	ZipCode  string `json:"zipcode"`
}

type MqttStatus struct {
	Action   string       `json:"action"`
	Ip       string       `json:"ip,omitempty"`
	Location *GeoLocation `json:"location,omitempty"`
	AdInfo   *AddressInfo `json:"ad_info,omitempty"`
}

func mqttOnDisconnect(client MQTT.Client, err error) {
	log.Printf("MqttClient Disconnected: err = %s\n", err)
	fmt.Printf("MqttClient Disconnected: err = %s\n", err)
}

func mqttOnConnect(client MQTT.Client) {
	log.Printf("MqttClient Connected\n")
	fmt.Printf("MqttClient Connected\n")
	mqttClient.Subscribe(noticeTopic, 1, nil)
	go reportOnline()
}

func mqttOnMessage(client MQTT.Client, msg MQTT.Message) {
	notice := MqttNotice{}
	payload := msg.Payload()
	err := json.Unmarshal(payload, &notice)
	if err != nil {
		log.Printf("mqttMessage invalid: %s\n", string(payload))
		fmt.Printf("mqttMessage invalid: %s\n", string(payload))
		//tnaerr.LastErr.Set(int(syscall.EBADMSG), string(payload), affairErrSrc)
		return
	}
	log.Printf("MqttNotice: %s\n", notice.Action)
	fmt.Printf("MqttNotice: %s\n", notice.Action)
	switch notice.Action {
	case "resetpin":
		log.Printf("resetpin\n")
		fmt.Printf("resetpin\n")
	case "updateacl":
	default:
		log.Printf("Unknown notice\n")
		fmt.Printf("Unknown notice\n")
	}
}

func StartMqttClient(id, svr string) {
	mqttMutex.Lock()
	defer mqttMutex.Unlock()
	if mqttClient != nil {
		mqttClient.Disconnect(1)
		mqttClient = nil
	}
	if mqttClient == nil {
		statusTopic = fmt.Sprintf("tnabox/%s/status", id)
		noticeTopic = fmt.Sprintf("tnabox/%s/notice", id)
		opts := MQTT.NewClientOptions()
		opts.AddBroker(svr).
			SetAutoReconnect(true).
			//SetCleanSession(true).
			SetClientID(id).
			SetConnectionLostHandler(mqttOnDisconnect). //连接断开
			SetOnConnectHandler(mqttOnConnect). //连接成功
			SetOrderMatters(false).
			SetKeepAlive(time.Second*20).
			SetDefaultPublishHandler(mqttOnMessage).
			SetWill(statusTopic, MQTT_WILL, 0, true)
		mqttClient = MQTT.NewClient(opts)
	}
	mqttClientConnect()
	time.Sleep(time.Second*2)
}

func mqttClientConnect() (err error) {
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		err = token.Error()
		log.Printf("MqttClient failed to connect: %s\n", err)
		fmt.Printf("MqttClient failed to connect: %s\n", err)
		//tnaerr.LastErr.Set(-1, "MqttClient failed", affairErrSrc)
	}
	return
}

func mqttSubscribeNotice() {
	if token := mqttClient.Subscribe(noticeTopic, 1, nil); token.Wait() && token.Error() != nil {
		err := token.Error()
		log.Printf("MqttClient failed to subscribe notice: %s\n", err)
		fmt.Printf("MqttClient failed to subscribe notice: %s\n", err)
		//tnaerr.LastErr.Set(-1, "MqttClient failed to subscribe notice", affairErrSrc)
	}
}

func reportOnline() {
	status := MqttStatus{Action: "online"}
	l, err := getLocationFromQQ()
	if err != nil {
		log.Printf("getGeoLocation failed: %s\n", err)
		fmt.Printf("getGeoLocation failed: %s\n", err)
	} else {
		status.Ip = l.Result.Ip
		status.Location = &l.Result.Location
		status.AdInfo = &AddressInfo{l.Result.AdInfo.Nation, l.Result.AdInfo.Province, l.Result.AdInfo.City, l.Result.AdInfo.District, ""}
		if l.Result.AdInfo.AdCode != 0 {
			status.AdInfo.ZipCode = fmt.Sprintf("%06d", l.Result.AdInfo.AdCode)
		}
	}
	data, _ := json.Marshal(&status)
	log.Printf("MqttStatus: %s\n", string(data))
	fmt.Printf("MqttStatus: %s\n", string(data))
	mqttClient.Publish(statusTopic, 0, true, data)
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

func getLocationFromQQ() (l *QQIpLocation, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 2,
	}
	req, err := http.NewRequest("GET", "https://apis.map.qq.com/ws/location/v1/ip?key=XNOBZ-FMFKK-2WHJ7-AT52S-JDDZ6-FOFFG", nil)
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
