package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

/*
http服务端
默认的Server
ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。
Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

自定义Server
要管理服务端的行为，可以创建一个自定义的Server：

	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

log.Fatal(s.ListenAndServe())
*/
func main() {
	//将代码编译之后执行，打开你电脑上的浏览器地址栏输入127.0.0.1:8080回车，输出Hello Go Server.
	http.HandleFunc("/", index)
	http.HandleFunc("/getRequestWithoutParameters", getRequestWithoutParameters)   //无参GET请求
	http.HandleFunc("/getRequestWithParameters", getRequestWithParameters)         //有参GET请求
	http.HandleFunc("/postRequestWithForm", postRequestWithForm)                   //携带表单数据的POST请求
	http.HandleFunc("/postRequestWithBody", postRequestWithBody)                   //携带请求体的POST请求
	http.HandleFunc("/getRequestRandomSlowResponse", getRequestRandomSlowResponse) //GET请求随机出现慢响应

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

// 首页
func index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Fprintln(w, "首页")
}

// 无参GET请求
func getRequestWithoutParameters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Fprintln(w, "无参GET请求")
}

// 有参GET请求
func getRequestWithParameters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Printf("有参GET请求 name:%v age:%v\n", data.Get("name"), data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

// 携带表单数据的POST请求
func postRequestWithForm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	form := r.PostForm
	fmt.Println("携带表单数据的POST请求接收到的数据为：", r.PostForm) // 打印form数据
	name := form.Get("name")
	age := form.Get("age")
	fmt.Printf("获取数据 name:%v age:%v\n", name, age)
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

// 携带请求体的POST请求
func postRequestWithBody(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//请求类型是application/json时从r.Body读取数据
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println("携带请求体的POST请求的JSON数据", string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

// GET请求随机出现慢响应
func getRequestRandomSlowResponse(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10) // 耗时10秒的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}
