package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

/*
log
实现简单日志服务
推荐使用第三方日志库：logrus、zap等
*/
func TestLog(t *testing.T) {
	useLog()              //使用自带的Logger
	setFlag()             //自定义日志格式
	setLogOutput()        //配置日志输出位置
	logger := diyLogger() //自定义Logger【非常有限，无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库】
	logger.Println("这是自定义的logger记录的日志。")
}

// 使用自带的Logger
func useLog() {
	/*
		使用Logger
		log包定义了Logger类型，该类型提供了一些格式化输出的方法。
		本包也提供了一个预定义的“标准”logger；
		可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用；
		比自行创建一个logger对象更容易使用。
	*/
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发fatal的日志。")
	//log.Panicln("这是一条会触发panic的日志。")
}

// 配置logger
func setFlag() {
	/*
		标准logger的配置
		默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法。

		log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。

			func Flags() int
			func SetFlags(flag int)

		flag选项
		log标准库提供了如下的flag选项，它们是一系列定义好的常量。
		const (
			// 控制输出日志信息的细节，不能控制输出的顺序和格式。
			// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
			Ldate         = 1 << iota     // 日期：2009/01/23
			Ltime                         // 时间：01:23:23
			Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
			Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
			Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
			LUTC                          // 使用UTC时间
			LstdFlags     = Ldate | Ltime // 标准logger的初始值
		)
	*/
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条带Flag的日志。")

	/*
		配置日志前缀
		func Prefix() string
		func SetPrefix(prefix string)
		其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。
	*/
	log.SetPrefix("[Meta39]")
	log.Println("这是一条带前缀的日志。")
}

// 配置日志输出位置
func setLogOutput() {
	wd, _ := os.Getwd()
	fmt.Println("获取当前工作目录（调试用）:", wd)
	logFile, err := os.OpenFile("../../log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//输出和配置不应该在一起，这里只是作示范。
	log.Println("这是一条带Flag的日志。")
	log.SetPrefix("[Meta39]")
	log.Println("这是一条带前缀的日志。")
}

// 使用标准的Logger，我们通常会把上面的配置操作写到init函数中。
/*func init() {
	wd, _ := os.Getwd()
	fmt.Println("获取当前工作目录（调试用）:", wd)
	logFile, err := os.OpenFile("../../log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}*/

/*
创建Logger
log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下：
func New(out io.Writer, prefix string, flag int) *Logger
New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。
*/
func diyLogger() *log.Logger {
	return log.New(os.Stdout, "【日志前缀】", log.Lshortfile|log.Ldate|log.Ltime)
}
