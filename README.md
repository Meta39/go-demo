# GOPROXY 要配置相同

如果代理不一致可能会导致拉依赖go.sum校验和出问题。

## Windows cmd
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## macOS or Linux
```shell
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

# go-demo
go-demo

## golang是基础包
基础包下面./base/test包下面包含
1. test测试用例
2. Go语言常用标准库
```text
一、fmt(fmt与格式化占位符)
二、time(时间和日期)
三、flag(命令行)
四、log(简单日志服务)
五、file(文件操作)
六、strconv(基本数据类型和字符串处理)
七、net/http(HTTP客户端和服务端)
八、context(上下文)
......
```
### time格式化和解析问题
```text
相同国家推荐：
时间：2025-10-01 10:00:00
格式化为：now.Format("2006-01-02 15:04:05")//string(当地时间)
解析：time.ParseInLocation("2006-01-02 15:04:05", "2025-10-01 10:00:00", time.Local)//Time(当地时区)

不同国家推荐：
时间：2025-10-01 10:00:00
格式化为：now.Format("2006-01-02 15:04:05")//格式化时间前转为UTC时间，防止解析丢失时区string(UTC)
解析：time.Parse("2006-01-02 15:04:05", "2025-10-01 10:00:00")//Time(UTC)
```

## overtime是给golang本地调用的测试包

## tcp是Go语言实现TCP通信
如果没有tcp/client和tcp/server项目，则按以下方式创建。
### 创建tcp/client客户端
1. 创建tcp目录
2. 在tcp目录创建client
3. client目录下执行go mod init client命令生成go.mod文件

### 创建tcp/server服务端
1. 进入tcp目录
2. 在tcp目录创建server
3. server目录下执行go mod init server命令生成go.mod文件