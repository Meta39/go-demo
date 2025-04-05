package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
HTTP客户端
基本的HTTP/HTTPS请求
Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。
resp, err := http.Get("http://example.com/")
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})

程序在使用完response后必须关闭回复的主体。
resp, err := http.Get("http://example.com/")

	if err != nil {
		// handle error
	}

defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)

自定义Client
要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

resp, err := client.Get("http://example.com")
// ...
req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
自定义Transport
要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}

client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
Client和Transport类型都可以安全的被多个goroutine同时使用。出于效率考虑，应该一次建立、尽量重用。
*/
func main() {
	httpRequestAddress := "http://localhost:8080"
	sendGetRequestWithoutParameters(httpRequestAddress + "/getRequestWithoutParameters") //发送无参GET请求
	sendGetRequestWithParameters(httpRequestAddress + "/getRequestWithParameters")       //发送带参数的GET请求
	sendPostRequestWithForm(httpRequestAddress + "/postRequestWithForm")                 //发送携带表单数据的POST请求
	sendPostRequestWithBody(httpRequestAddress + "/postRequestWithBody")                 //发送携带请求体的POST请求
}

// 发送无参GET请求
func sendGetRequestWithoutParameters(requestUrl string) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) //ioutil.ReadAll已经不推荐使用了，推荐使用：io.ReadAll
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Printf("不参数的GET请求%s的返回内容为：%s\n", requestUrl, string(body))
}

// 发送带参数的GET请求
func sendGetRequestWithParameters(requestUrl string) {
	// URL param
	data := url.Values{}
	data.Set("name", "Meta39")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	//完整请求地址
	fullRequestAddress := u.String()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Printf("带参数的GET请求%s返回的内容为：%s\n", fullRequestAddress, string(b))
}

// 发送POST请求（表单数据）
func sendPostRequestWithForm(requestUrl string) {
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=Meta39&age=18"
	resp, err := http.Post(requestUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Printf("发送POST请求（表单数据）%s的返回内容为：%s\n", requestUrl, string(b))
}

// 发送携带请求体的POST请求
func sendPostRequestWithBody(requestUrl string) {
	// body json
	contentType := "application/json"
	data := `{"name":"Meta39","age":18}`
	resp, err := http.Post(requestUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Printf("发送携带请求体的POST请求%s的返回内容为：%s\n", requestUrl, string(b))
}
