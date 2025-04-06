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

## tcp
如果没有tcp/client和tcp/server项目，则按以下方式创建。
### 创建tcp/client客户端
1. 创建tcp目录
2. 在tcp目录创建client
3. client目录下执行go mod init client命令生成go.mod文件

### 创建tcp/server服务端
1. 进入tcp目录
2. 在tcp目录创建server
3. server目录下执行go mod init server命令生成go.mod文件

## http
如果没有http/client和http/server项目，则按以下方式创建。

注意：net/http不适用于企业级开发。企业级开发建议选择：Gin、Echo等第三方开源库。

### 何时选择 net/http？
1. 小型工具或对依赖极度敏感的项目。
2. 需要完全控制底层实现（如自定义协议）。
3. 学习HTTP原理的练手场景。

### 创建http/client客户端
1. 创建http目录
2. 在http目录创建client
3. client目录下执行go mod init client命令生成go.mod文件

### 创建http/server服务端
1. 进入http目录
2. 在http目录创建server
3. server目录下执行go mod init server命令生成go.mod文件

### context（http Server 上下文）注意事项
1. 推荐以参数的方式显示传递Context
2. 以Context作为参数的函数方法，应该把Context作为第一个参数。
3. 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
4. Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
5. Context是线程安全的，可以放心的在多个goroutine中传递

## database/mysql
```text
Go操作MySQL数据库
1、简单的增查改删(CRUD)
2、预编译SQL，防止SQL注入
3、MySQL事务控制(ACID)
事务必须满足4个条件：原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）
```
注意：
1. 企业级开发推荐使用第三方开源ORM库：GORM、sqlx等
2. 企业级开发推荐使用第三方开源Web库：Gin、Echo等

