// mqtt_test.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/16$ 10:13$

package gpmqtt

import (
	"fmt"
	"strings"
	"testing"
)

func TestConnect(t *testing.T) {
	//add := "tcp://10.9.0.1:1883"
	sn := "A1620S80004"
	b := []byte(sn)
	a := [32]byte{}
	copy(a[0:len(b)],b)
	fmt.Println("a",a)
	fmt.Println("len(a)",len(a))

	c := string(a[:])
	fmt.Println("c",strings.TrimRight(c,string([]byte{})))
	fmt.Println("c",c[12:])

	d:=0
	for i:= 0;i<len(a);i++{
		fmt.Println("a[i]",a[i])
		if a[i] == 0 {
			fmt.Println("a[i]=====",a[i])
			d=i
			break
		}
	}
	fmt.Println("d===",d)
	fmt.Println(string(a[:d]))
	//StartMqttClient("",add)
}
