package message

import (
	"encoding/json"
	"gitee.com/gjing1st/gopackage/net/gphttp"
)

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}
type TextCardReq struct {
	BaseParams
	TextCard `json:"textcard"`
}

// GetRequestParams
// @description: 文字卡片参数组装
// @param: req TextCardReq
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 22:46
// @success:
func GetRequestParams(req TextCardReq) (reqBytesData []byte, err error) {
	text := TextCard{
		req.Title,
		req.Description,
		req.Url,
		req.Btntxt,
	}
	param := BaseParams{
		Touser:                 req.Touser,
		Msgtype:                "textcard",
		Agentid:                req.Agentid,
		EnableUdTrans:          req.EnableUdTrans,
		EnableDuplicateCheck:   req.EnableDuplicateCheck,
		DuplicateCheckInterval: req.DuplicateCheckInterval,
	}
	reqParams := TextCardReq{
		param,
		text,
	}
	reqBytesData, err = json.Marshal(reqParams)
	return
}

// SendCardMessage
// @description: 发送卡片消息
// @param: accessToken
// @param: req TextCardReq
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 22:47
// @success:
func SendCardMessage(accessToken string,req TextCardReq) (result []byte, err error) {
	reqUrl := getRequestUrl(accessToken)
	reqParams, err := GetRequestParams(req)
	if err != nil {
		return nil, err
	}
	return gphttp.PostJson(reqUrl, reqParams)

}
