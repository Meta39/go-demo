package test

import (
	"fmt"
	"testing"
	"time"
)

/*
Go 语言内置的 time 包的基本用法。
time 包提供了一些关于时间显示和测量用的函数。
time 包中日历的计算采用的是公历，不考虑润秒。
*/

func TestTime(t *testing.T) {
	fmt.Println("============ time ============")
	timeDemo()      //日期时间
	timezoneDemo()  //时区
	timestampDemo() //Unix Time是自1970年1月1日 00:00:00 UTC 至当前时间经过的总秒数
	timeOption()    //时间间隔操作
	formatDemo()    //时间格式化【时间格式化成字符串，推荐：now.Format("2006-01-02 15:04:05")】
	parseDemo()     //解析字符串格式的时间【字符串根据对应的格式转为时间，推荐：time.ParseInLocation("2006-01-02 15:04:05", "2025-04-04 13:34:14", time.Local)】
	timeZone()      //解决序列化和反序列化，时区丢失的问题
	tickDemo()      //定时器

	fmt.Println("============ time ============")
}

// 时间结构体的年月日时分秒
func timeDemo() {
	/*
		时间类型
		Go 语言中使用time.Time类型表示时间。我们可以通过time.Now函数获取当前的时间结构体，然后从时间结构体中可以获取到年、月、日、时、分、秒等信息。
	*/
	now := time.Now() // 获取当前时间
	fmt.Printf("获取当前时间:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Printf("年:%v 月:%v 日:%v 时:%v 分:%v 秒:%v\n", year, month, day, hour, minute, second)
}

// 时区示例
func timezoneDemo() {
	/*
		Go 语言中使用 location 来映射具体的时区。时区（Time Zone）是根据世界各国家与地区不同的经度而划分的时间定义，全球共分为24个时区。
		中国差不多跨5个时区，但为了使用方便只用东八时区的标准时即北京时间为准。
		在日常编码过程中使用时间结构体的时候一定要注意其时区信息；
		由于time.LoadLocation依赖系统的时区数据库，在不太确定程序运行环境的情况下建议先定义时区偏移量然后使用time.FixedZone的方式指定时区。
		下面的示例代码中使用beijing来表示东八区8小时的偏移量，其中time.FixedZone和time.LoadLocation这两个函数则是用来获取location信息。
	*/
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds()) //获取8小时的秒数
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	// 加载上海所在的时区
	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

	// 创建时间结构体需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing) //true
	fmt.Println("北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间:", timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork) //true
	fmt.Println("纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间", timesAreEqual)

	/*
		如果你的程序需要在不同的国家运行，但是使用的是同一个数据库，则使用以下方式。
		1.获取当地时间，然后转为UTC时间存储到数据库，这样不同国家的时间都是UTC时间的方式存储
		2.当需要不同国家展示当地时间的时候，取出数据库的UTC时间，然后通过time.Local获取本地时区信息，通过UTC.In(time.Local)转为当地时间进行展示。
	*/
	localNow := time.Now() //获取当地时间
	fmt.Println("当前时间为：", localNow)
	//数据库存储UTC时间
	utcTime := localNow.UTC() //当地时间转为UTC时间
	fmt.Println("当前UTC时间为：", utcTime)
	//前端展示时间（UTC + 当地时区）
	localNow2 := utcTime.In(time.Local) //数据库UTC时间根据当地时区转为当地时间
	fmt.Println("UTC转为当地时间：", localNow2)

}

// 时间戳
func timestampDemo() {
	/*
		Unix Time是自1970年1月1日 00:00:00 UTC 至当前时间经过的总秒数。
	*/
	now := time.Now()        // 获取当前时间
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Printf("秒级时间戳:%v，毫秒时间戳:%v，微秒时间戳:%v，纳秒时间戳:%v\n", timestamp, milli, micro, nano)

	// 获取北京时间所在的东八区时区对象
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 北京时间 2025-04-22 14:00:30.000000022 +0800 CST
	t := time.Date(2025, 04, 22, 14, 00, 30, 5, beijing)

	var (
		sec  = t.Unix()
		msec = t.UnixMilli()
		usec = t.UnixMicro()
	)

	// 将秒级时间戳转为时间结构体（第二个参数为不足1秒的纳秒数）
	timeObj := time.Unix(sec, 5)
	fmt.Println("秒级时间戳转为时间结构体:", timeObj)  // 2025-04-22 14:00:30.000000005 +0800 CST
	timeObj = time.UnixMilli(msec)         // 毫秒级时间戳转为时间结构体
	fmt.Println("毫秒级时间戳转为时间结构体:", timeObj) // 2025-04-22 14:00:30 +0800 CST
	timeObj = time.UnixMicro(usec)         // 微秒级时间戳转为时间结构体
	fmt.Println("微秒级时间戳转为时间结构体:", timeObj) // 2025-04-22 14:00:30 +0800 CST
}

// 时间间隔操作
func timeOption() {
	/*
		time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。time.Duration表示一段时间间隔，可表示的最长时间段大约290年。
		time 包中定义的时间间隔类型的常量如下：

		const (
			Nanosecond  Duration = 1//纳秒
			Microsecond          = 1000 * Nanosecond//微秒
			Millisecond          = 1000 * Microsecond//毫秒
			Second               = 1000 * Millisecond//秒
			Minute               = 60 * Second//分
			Hour                 = 60 * Minute//时
		)

		func (t Time) Add(d Duration) Time//传入的时间+时间间隔常量【t.Add(time.Hour/Minute/Second/Millisecond/Microsecond/Nanosecond)】
		func (t Time) Sub(u Time) Duration//t - u 的差值(得到的是2个时间相差的时间)
		func (t Time) Equal(u Time) bool//判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
		func (t Time) Before(u Time) bool//如果t代表的时间点在u之前，返回真；否则返回假。
		func (t Time) After(u Time) bool//如果t代表的时间点在u之后，返回真；否则返回假。
	*/
	now := time.Now()
	fmt.Println("当前时间：", now)
	nowAdd := now.Add(time.Hour)
	fmt.Println("加一个小时以后的时间：", nowAdd)
	timeSub := nowAdd.Sub(now)
	fmt.Printf("时间差：【%v】 和 【%v】 相差 【%v】\n", nowAdd, now, timeSub)
	fmt.Printf("时间差：【%v】 和 【%v】 相差 【%v】\n", now, nowAdd, timeSub)
	newYorkLoadLocation, _ := time.LoadLocation("America/New_York")
	newYorkNow := now.In(newYorkLoadLocation)
	fmt.Println("北京当前时间：", now)
	fmt.Println("美国/纽约当前时间：", newYorkNow)
	fmt.Println("时间是否相等【不同时区标准的时间也可以正确比较】，2个不同时区的当前时间：北京时间和纽约时间是否相等：", now.Equal(newYorkNow))
	fmt.Printf("%v 是否在 %v 之前？结果：%v\n", now, nowAdd, now.Before(nowAdd))
	fmt.Printf("%v 是否在 %v 之后？结果：%v\n", now, nowAdd, now.After(nowAdd))
}

// 时间格式化【时间格式化成字符串】
func formatDemo() {
	now := time.Now()
	/*
		Go 和 Java 时间格式化性能对比
		Go 的更为高效，因为其 time 包是底层实现，直接操作时间数据结构。
		Java 略低于 Go，因为 DateTimeFormatter 是不可变的，每次格式化都会创建新的对象（尽管 JVM 会优化）。
	*/
	// 24小时制格式化的模板为 2006-01-02 15:04:05【等价于java的：yyyy-MM-dd HH:mm:ss】
	// 12小时制格式化的模板为 2006-01-02 03:04:05【等价于java的：yyyy-MM-dd hh:mm:ss】

	// 24小时制
	fmt.Println("24小时制【展示标准输出】", now.Format("2006-01-02 15:04:05"))
	fmt.Println("24小时制", now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制【用的少】
	fmt.Println("12小时制", now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println("小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数", now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println("小数点后写9，会省略末尾可能出现的0", now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	fmt.Println("只格式化时分秒部分", now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println("只格式化日期部分", now.Format("2006.01.02"))
}

// 解析字符串格式的时间
func parseDemo() {
	// time.Parse默认将无时区字符串解析为UTC时间，导致时间点偏差（如12:00:00 UTC+8 和 12:00:00 UTC 时间点不同）
	timeObj, _ := time.Parse("2006-01-02 15:04:05", "2025-04-04 13:34:14")
	fmt.Println("time.Parse 返回UTC时间【不推荐】：", timeObj) // 2025-04-04 13:34:14 +0000 UTC
	//time.ParseInLocation允许指定解析时使用的时区（如time.Local），确保解析后的时间与原始时间的时区一致，从而时间点相同。
	timeObj2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-04-04 13:34:14", time.Local)
	fmt.Println("time.ParseInLocation指定解析时使用的时区（如time.Local）【推荐】：", timeObj2) //推荐

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	//time.RFC3339 "2006-01-02T15:04:05Z07:00"
	/*
		Z 是时区偏移量的占位符 + 表示东时区（UTC 时间之后的时间）
		+ 表示东时区（UTC 时间之后的时间）。如：+08:00，表示 UTC + 8 小时
		- 表示西时区（UTC 时间之前的时间）。如：-08:00，表示 UTC - 8 小时
	*/
	timeObj, _ = time.Parse("2006-01-02 15:04:05Z07:00", "2025-04-04 13:34:14+08:00")
	fmt.Println(timeObj) // 2025-04-04 13:34:14 +0800 CST
}

// 解决序列化和反序列化，时区丢失的问题
func timeZone() {
	t := time.Now()
	tString := t.Format("2006-01-02 15:04:05")
	t2, _ := time.ParseInLocation(time.DateTime, tString, time.Local)
	fmt.Println(t.Equal(t2)) // 输出 true
}

// 定时器
func tickDemo() {
	ticker := time.Tick(3 * time.Second) //定义一个3秒间隔的定时器
	times := 0
	for i := range ticker {
		//执行三次以后退出，防止一直执行，卡住程序。一般情况下就是要一直去执行的，但是这里为了让后面的程序也能正常运行，所以执行三次后停止执行
		if times == 3 {
			break
			//return//用return也行
		}
		fmt.Println("每3秒都会执行的任务，i表示当前时间：", i) //每3秒都会执行的任务
		times++
	}
}
