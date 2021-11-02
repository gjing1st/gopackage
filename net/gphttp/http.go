// http.go
// Created by BestTeam.
// User: GJing
// WeChat: ks_kdb
// Date: 2021/11/2$ 10:22$

package gphttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/gjing1st/gopackage/gplog"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// GetRequest
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 14:11
// @return:
func GetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.LogFile("request", "请求", url, "失败,err=", err)
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.LogFile("request", "fetch: reading %s: %v\n", url, err)
		return nil, err
	}
	return b, nil
}

// GetRequestWithHeader
// @description: 带请求头的get请求
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 14:13
// @return:
func GetRequestWithHeader(url string)  ([]byte, error){
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("err")
	}
	// 添加请求头
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Add("header", "header")
	// 添加cookie
	cookie1 := &http.Cookie{
		Name:  "aaa",
		Value: "aaa-value",
	}
	req.AddCookie(cookie1)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	return b,err
}

// UrlPost
// @description: post请求
// @param: apiUrl 要请求的url地址
// @param: postParam 请求参数
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 16:36
// @return:
func UrlPost(apiUrl string,postParam map[string]string)(result map[string]interface{}, err error){
	postValue := url.Values{}
	for key, value := range postParam{
		postValue.Set(key, value)
	}
	response, err := http.Post(apiUrl, "application/x-www-form-urlencoded", strings.NewReader(postValue.Encode()))
	obj := make(map[string]interface{})
	if err != nil{
		return nil, err
	}
	text, err2 := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err2 != nil{
		return nil, err2
	}
	err3 :=  json.Unmarshal(text, &obj)
	return obj, err3
}

// PostJson
// @description: post请求，参数为josn
// @param: reqUrl 要请求的url地址
// @param: bytesData json字节
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/11/2 22:18
// @success: reqUrl接口返回的字节数组
func PostJson(reqUrl string,bytesData []byte) (result []byte, err error) {

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", reqUrl, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	result, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	////byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&result))
	//fmt.Println(*str)
	return result,nil
}

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

//body提交二进制数据
func DoBytesPost(url string, data []byte) ([]byte, error) {

	body := bytes.NewReader(data)
	request, err := http.NewRequest("post", url, body)
	if err != nil {
		log.Println("http.NewRequest,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	request.Header.Set("Connection", "Keep-Alive")
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
		return []byte(""), err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http.Do failed,[err=%s][url=%s]", err, url)
	}
	return b, err
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, err
}


const HTTP_CT_JSON = "application/json"

func JsonRestRequest(method, url string, req, res interface{}) (sc int, err error) {
	var reader io.Reader
	var resp *http.Response
	var data []byte
	sc = -1
	if req != nil {
		data, err = json.Marshal(req)
		if err != nil {
			return
		}
		reader = bytes.NewReader(data)
	}
	if method == "GET" {
		resp, err = http.Get(url)
	} else if method == "POST" {
		resp, err = http.Post(url, HTTP_CT_JSON, reader)
	} else {
		var req *http.Request
		req, err = http.NewRequest(method, url, reader)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", HTTP_CT_JSON)
		client := http.Client{}
		resp, err = client.Do(req)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	sc = resp.StatusCode
	if sc < 200 || sc >= 300 {
		err = syscall.EINVAL
		return
	}
	if res != nil {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, res)
	}
	return
}
