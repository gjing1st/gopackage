package message

import (
	"fmt"
	"github.com/gjing1st/gopackage/workweixin"
	"github.com/gjing1st/gopackage/workweixin/message"
	"testing"
)

func TestSendCardMessage(t *testing.T) {
	text := message.TextCard{
		"待审批通知",
		" <div class=\\\"normal\\\">您收到一条待审批访客单</div><div class=\\\"highlight\\\">请登录系统查看审批</div>",
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=ww887a89a56b64260f&redirect_uri=http://visit.sd1st.top/login&response_type=code&scope=snsapi_base",
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
