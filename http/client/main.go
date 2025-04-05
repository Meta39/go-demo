package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
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
	// 定义一个100毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()                                                                            // 调用cancel释放子goroutine资源
	sendGetRequestRandomSlowResponse(ctx, httpRequestAddress+"/getRequestRandomSlowResponse") //发送GET请求随机出现慢响应
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

type respData struct {
	resp *http.Response
	err  error
}

/*
一、Context概述
Context用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。
对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。
它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。
当一个上下文被取消时，它派生的所有上下文也被取消。

二、Context接口
context.Context是一个接口，该接口定义了四个需要实现的方法。具体签名如下：

	type Context interface {
	    Deadline() (deadline time.Time, ok bool)
	    Done() <-chan struct{}
	    Err() error
	    Value(key interface{}) interface{}
	}

其中：
（1）Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
（2）Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
（3）Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
（4）如果当前Context被取消就会返回Canceled错误；
（5）如果当前Context超时就会返回DeadlineExceeded错误；
（6）Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

三、Background()和TODO()
Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。
我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。
Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。
background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

四、With系列函数
（1）func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。
取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
（2）func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
返回父上下文的副本，并将deadline调整为不迟于d。
如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
（3）func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。
（4）func WithValue(parent Context, key, val interface{}) Context
返回父节点的副本，其中与key关联的值为val。
仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。
所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
WithValue的用户应该为键定义自己的类型。为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。
或者，导出的上下文关键变量的静态类型应该是指针或接口。

五、使用Context的注意事项
（1）推荐以参数的方式显示传递Context
（2）以Context作为参数的函数方法，应该把Context作为第一个参数。
（3）给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
（4）Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
（5）Context是线程安全的，可以放心的在多个goroutine中传递
*/
//发送GET请求随机出现慢响应
func sendGetRequestRandomSlowResponse(ctx context.Context, requestUrl string) {
	transport := http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives: true}
	client := http.Client{
		Transport: &transport,
	}

	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		//fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		//transport.CancelRequest(req)
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := io.ReadAll(result.resp.Body)
		fmt.Printf("发送GET请求随机出现慢响应%s的返回内容为：%v\n", requestUrl, string(data))
	}
}
