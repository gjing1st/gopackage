package message

import (
	"fmt"
	"github.com/gjing1st/gopackage/workweixin/message"
	"gopackage/workweixin"
	"testing"
)

func TestSendCardMessage(t *testing.T) {
	text := message.TextCard{
		"程序员节领奖通知",
		"<div class=\\\"gray\\\">2021年10月21日</div> <div class=\\\"normal\\\">恭喜你抽中iPhone 13一台，领奖码：XSDEF</div><div class=\\\"highlight\\\">请于2021年10月25日前联系行政同事领取</div>",
		"https://www.baidu.com",
		"查看详情",
	}
	param := message.BaseParams {
		Touser: "GuoJing",
		Msgtype: "textcard",
		Agentid: 1000002,
		EnableUdTrans:0,
		EnableDuplicateCheck: 0,
		DuplicateCheckInterval:1800,
	}
	req := message.TextCardReq{
		param,
		text,
	}
	//token := "fMbJLQyslN7kA6zQud4_gzE3hcCJYG0eM9bEgQQc9MJR5NH99GeGvRZU_ZVeKVO10WqFV5POok3Kgofv_4_iK4Q3ZR7BPge3YzujAm3rSXizp8BFyL-JOByKlLWyaC-frb1E5nGpIb19Oml9zibnwNbEkjnBgl1w01l32dmuQACNwrGw0qWyyFuKpXTO1JZ8PsyIdDjAm-2Ir2cMI19Akw"
	corpid:= "ww887a89a56b64260f"
	corpsecret := "Xh0Cn3oC1n997ImN3ZQlsWFkGIk8rlsgDPIQUp4PPpA"
	accessToken,_:= workweixin.GetAccessToken(corpid,corpsecret)
	fmt.Println("accessToken",accessToken)
	res,err := message.SendCardMessage(accessToken,req)
	fmt.Println(string(res),err)

}
