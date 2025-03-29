package main

/*
执行go get github.com/Meta39/gohello/v2@v2.0.0命令时要注意 go.sum 是否一致，如果始终无法导入，则删除里面的内容。
"github.com/Meta39/gohello" 和 "github.com/Meta39/gohello/v2"是同一个项目不通版本，go不知道用哪个版本，因此要用gohello2别名区分
*/
import (
	"fmt"
	"github.com/Meta39/gohello"
	gohello2 "github.com/Meta39/gohello/v2"
	"github.com/Meta39/overtime"
	"golang/base/demo"
)

/*
Package 在工程化的Go语言开发项目中，Go语言的源码复用是建立在包（package）基础之上的。

包与依赖管理
在工程化的Go语言开发项目中，Go语言的源码复用是建立在包（package）基础之上的。
Go语言中如何定义包、如何导出包的内容及如何引入其他包。
如何在项目中使用go module管理依赖。

包介绍
Go语言中支持模块化的开发理念，在Go语言中使用包（package）来支持代码模块化和代码复用。
一个包是由一个或多个Go源码文件（.go结尾的文件）组成，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如fmt、os、io等。

定义包
我们可以根据自己的需要创建自定义包。
一个包可以简单理解为一个存放.go文件的文件夹。
该文件夹下面的所有.go文件都要在非注释的第一行添加如下声明，声明该文件归属的包。
package packagename
其中：
package：声明包的关键字
packagename：包名，可以不与文件夹的名称一致，不能包含 - 符号，最好与其实现的功能相对应。
另外需要注意一个文件夹下面直接包含的文件只能归属一个包，同一个包的文件不能在多个文件夹下。
包名为main的包是应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。
PS：如果发布的包错了，此时又已经拉了错误的包。go.mod文件下的go.sum就会记录错误的包的加密哈希值，因此重新打包后，用 go mod tidy重新生成依赖 或 手动解决go.sum冲突（优先保留最新的变更）。
*/
func Package() {
	fmt.Println("============ 包 ============")
	/*
		标识符可见性 PS：类似java的作用域关键字public、private、default、protected，只不过go只有public（大写字母开头）和private（小写字母开头）
		在同一个包内部声明的标识符都位于同一个命名空间下，在不同的包内部声明的标识符就属于不同的命名空间。
		想要在包的外部使用包内部的标识符就需要添加包名前缀，例如fmt.Println("Hello world!")，就是指调用fmt包中的Println函数。
		如果想让一个包中的标识符（如变量、常量、类型、函数等）能被外部的包使用，那么标识符必须是对外可见的（public）。
		在Go语言中是通过标识符的首字母大/小写来控制标识符的对外可见（public）/不可见（private）的。在一个包内部只有首字母大写的标识符才是对外可见的。
		例如我们定义一个名为demo的包，在其中定义了若干标识符。在另外一个包中并不是所有的标识符都能通过demo.前缀访问到，因为只有那些首字母是大写的标识符才是对外可见的。
	*/
	demo.Add(1, 2)

	/*
		包的引入
		要在当前包中使用另外一个包的内容就需要使用import关键字引入这个包，并且import语句通常放在文件的开头，package声明语句的下方。
		完整的引入声明语句格式如下:
		import importname "path/to/package"
		其中：
		importname：引入的包名，通常都省略。默认值为引入包的包名。
		path/to/package：引入包的路径名称，必须使用双引号包裹起来。
		Go语言中禁止循环导入包。

		一个Go源码文件中可以同时引入多个包，例如：
		import "fmt"
		import "net/http"
		import "os"

		//批量引入
		import (
			"fmt"
			"net/http"
			"os"
		)

		当引入的多个包中存在相同的包名或者想自行为某个引入的包设置一个新包名时，都需要通过importname指定一个在当前文件中使用的新包名。
		例如，在引入fmt包时为其指定一个新包名f。
		import f "fmt"
		这样在当前这个文件中就可以通过使用f来调用fmt包中的函数了。
		f.Println("Hello world!")

		如果引入一个包的时候为其设置了一个特殊_作为包名，那么这个包的引入方式就称为匿名引入。
		一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。
		被匿名引入的包中的init函数将被执行并且仅执行一遍。
		import _ "github.com/go-sql-driver/mysql"
		匿名引入的包与其他方式导入的包一样都会被编译到可执行文件中。
		需要注意的是，Go语言中不允许引入包却不在代码中使用这个包的内容，如果引入了未使用的包则会触发编译错误。
		PS：GoLand IDE 会帮我们自动去除不存在的包和导入需要的包，因此很少会出现上面注意的情况。其它编辑器就不知道了。
	*/

	/*
		init初始化函数 PS：类似 java 中 class 的 static 块，只执行一次。
		在每一个Go源文件中，都可以定义任意个如下格式的特殊函数：
		func init(){
		  // ...
		}
		这种特殊的函数不接收任何参数也没有任何返回值，我们也不能在代码中主动调用它。当程序启动的时候，init函数会按照它们声明的顺序自动执行
		一个包的初始化过程是按照代码中引入的顺序来进行的，所有在该包中声明的init函数都将被串行调用并且仅调用执行一次。
		每一个包初始化的时候都是先执行依赖的包中声明的init函数再执行当前包中声明的init函数。
		确保在程序的main函数开始执行时所有的依赖包都已初始化完成。
		PS：类似java的 static 块，只会执行一次。
	*/

	/*
		go module （PS：可以简单理解为Maven、Gradle、Ant等依赖管理工具）
		Go module 是 Go1.11 版本发布的依赖管理方案，从 Go1.14 版本开始推荐在生产环境使用，于Go1.16版本默认开启。
		Go module 提供了以下命令供我们使用：
		命令				介绍
		go mod init		初始化项目依赖，生成go.mod文件
		go mod download	根据go.mod文件下载依赖
		go mod tidy		比对项目文件中引入的依赖与go.mod进行比对
		go mod graph	输出依赖关系图
		go mod edit		编辑go.mod文件
		go mod vendor	将项目的所有依赖导出至vendor目录
		go mod verify	检验一个依赖包是否被篡改过
		go mod why		解释为什么需要某个依赖

		Go语言在 go module 的过渡阶段提供了 GO111MODULE 这个环境变量来作为是否启用 go module 功能的开关。
		考虑到 Go1.16 之后 go module 已经默认开启，所以不再介绍该配置，对于刚接触Go语言的而言完全没有必要了解这个历史包袱。

		GOPROXY （PS：类似Maven、Gradle、Ant公共仓库代理）
		这个环境变量主要是用于设置 Go 模块代理（Go module proxy），其作用是用于使 Go 在后续拉取模块版本时能够脱离传统的 VCS 方式，直接通过镜像站点来快速拉取。
		GOPROXY 的默认值是：https://proxy.golang.org,direct，由于某些原因国内无法正常访问该地址，所以我们通常需要配置一个可访问的地址。
		目前社区使用比较多的有两个https://goproxy.cn和https://goproxy.io，当然如果你的公司有提供GOPROXY地址那么就直接使用。
		设置GOPAROXY的命令如下：
		go env -w GOPROXY=https://goproxy.cn,direct
		GOPROXY 允许设置多个代理地址，多个地址之间需使用英文逗号 “,” 分隔。最后的 “direct” 是一个特殊指示符，用于指示 Go 回源到源地址去抓取（比如 GitHub 等）。
		当配置有多个代理地址时，如果第一个代理地址返回 404 或 410 错误时，Go 会自动尝试下一个代理地址，当遇见 “direct” 时触发回源，也就是回到源地址去抓取。

		GOPRIVATE（PS：类似Maven、Gradle、Ant私有仓库代理）
		设置了GOPROXY 之后，go 命令就会从配置的代理地址拉取和校验依赖包。
		当我们在项目中引入了非公开的包（公司内部git仓库或 github 私有仓库等），此时便无法正常从代理拉取到这些非公开的依赖包，这个时候就需要配置 GOPRIVATE 环境变量。
		GOPRIVATE用来告诉 go 命令哪些仓库属于私有仓库，不必通过代理服务器拉取和校验。
		GOPRIVATE 的值也可以设置多个，多个地址之间使用英文逗号 “,” 分隔。我们通常会把自己公司内部的代码仓库设置到 GOPRIVATE 中，例如：
		go env -w GOPRIVATE="git.mycompany.com"
		这样在拉取以git.mycompany.com为路径前缀的依赖包时就能正常拉取了。
		此外，如果公司内部自建了 GOPROXY 服务，那么我们可以通过设置 GONOPROXY=none，允许通内部代理拉取私有仓库的包。
	*/

	/*
		使用go module引入本地的一个包
		如果你想要导入本地的一个包，并且这个包也没有发布到到其他任何代码仓库，这时候你可以在go.mod文件中使用replace语句将依赖临时替换为本地的代码包。
		例如在我的电脑上有另外一个名为 github.com/Meta39/overtime 的项目，它位于 golang 项目同级目录下。
		由于 github.com/Meta39/overtime 包只存在于我本地，并不能通过网络获取到这个代码包，这个时候应该如何在 golang 项目中引入它呢？
		我们可以在 golang/go.mod 文件中正常引入 github.com/Meta39/overtime 包，然后像下面的示例那样使用 replace 语句将这个依赖替换为使用相对路径表示的本地包。
		在go-demo文件夹下面创建overtime文件夹并在文件夹下进行初始化。
		即：
		1、cd ./overtime
		2、go mod init overtime
		3、创建overtime.go文件
		4、创建一个公共函数Hello，可以在全局访问。以供当前项目调用
		func Hello() {
			fmt.Println("Hello Meta39")
		}
		5、golang/go.mod 引入 github.com/Meta39/overtime
			module golang
			go [版本号]
			require github.com/Meta39/gohello v0.1.0
			require github.com/Meta39/overtime v0.0.0
			replace github.com/Meta39/overtime  => ../overtime
	*/
	overtime.Hello() //引用golang同级的overtime项目，并调用Hello()函数

	/*
		使用go module发布包
		1、编写一个代码包并将它发布到github.com仓库，让它能够被全球的Go语言开发者使用。
		2、在自己的 github 账号下新建一个 gohello 只包含 README.md 的项目，并把它下载到本地。git clone https://github.com/Meta39/gohello
		3、初始化项目。cd gohello => go mod init github.com/Meta39/gohello
		4、创建一个 hello.go 文件并创建 GoHello 函数，输出 GoHello。代码如下：
		package gohello

		import "fmt"

		func GoHello() {
			fmt.Println("GoHello")
		}
		5、然后将该项目的代码 push 到仓库的远端分支，这样就对外发布了一个Go包。其他的开发者可以通过 github.com/Meta39/gohello 这个引入路径下载并使用这个包了。
			git config user.name Meta39
			git config user.email 5399553@qq.com
			git add .
			git commit -m "描述"
			git remote add [远程仓库分支别名] [远程SSH或者Https]
			git push [远程仓库分支别名] [远程仓库分支名]
			//推送github.com/Meta39/gohello中的hello.go和GoHello函数以后再发布版本
			git tag -a v0.1.0 -m "release version v0.1.0"
			git push [远程仓库分支别名] v0.1.0
			//然后就会在自己的github仓库Releases中有发布一个名为v0.1.0的Tag，供当前项目使用
		6、一个设计完善的包应该包含开源许可证及文档等内容，并且我们还应该尽心维护并适时发布适当的版本。github 上发布版本号使用git tag为代码包打上标签即可。
		7、Go modules中建议使用语义化版本控制，其建议的版本号格式如下：v主版本号.次版本号.修订号
			主版本号：发布了不兼容的版本迭代时递增（breaking changes）。
			次版本号：发布了功能性更新时递增。
			修订号：发布了bug修复类更新时递增。
		8、go.mod 配置 require github.com/Meta39/gohello v0.1.0
		9、在当前项目的根路径下(golang)输入：go mod download下载gohello项目
		10、导入 import github.com/Meta39/gohello 包
		11、调用GoHello函数
		PS：如果gohello打错包，比如没有修改v2直接打包。但是命名时v2.0.0
		此时go get github.com/Meta39/gohello/v2@v2.0.0会找不到包，修改gohello go.mod文件后要删除原来的tag，重新打tag。
		golang 再次执行go get github.com/Meta39/gohello/v2@v2.0.0前需要删除go.sum里面有关v2的包，否则会拉取失败。
	*/
	fmt.Println("调用自己开源的包github.com/Meta39/gohello.GoHello函数输出GoHello")
	gohello.GoHello()

	/*
		发布新的主版本
		1、现在我们的 gohello 项目要进行与之前版本不兼容的更新，我们计划让 SayHi 函数支持向指定人发出问候。更新后的 SayHi 函数内容如下：
		package gohello

		import "fmt"

		// SayHi 向指定人打招呼的函数
		func SayHi(name string) {
			fmt.Printf("你好%s，我是Meta39。很高兴认识你。\n", name)
		}

		2、由于这次改动巨大（修改了函数之前的调用规则），对之前使用该包作为依赖的用户影响巨大。
		因此我们需要发布一个主版本号递增的v2版本。在这种情况下，我们通常会修改当前包的引入路径，像下面的示例一样为引入路径添加版本后缀。

		// gohello/go.mod修改如下
		//v0.1.0
		//module github.com/Meta39/gohello

		//v2.0.0
		module github.com/Meta39/gohello/v2

		命令行操作把修改后的代码提交：
			git add .
			git commit -m "feat: SayHi现在支持给指定人打招呼啦"
			git push
		打好 tag 推送到远程仓库。
			git tag -a v2.0.0 -m "release version v2.0.0"
			git push [远程仓库分支别名] v2.0.0

		3、这样在不影响使用旧版本的用户的前提下，我们新的版本也发布出去了。想要使用v2版本的代码包的用户只需按修改后的引入路径下载即可。
			当前项目根目录（golang）命令行输入
			go get github.com/Meta39/gohello/v2@v2.0.0

		4、在代码中使用的过程与之前类似，只是需要注意引入路径要添加 v2 版本后缀。

			import "github.com/Meta39/gohello/v2" // 引入v2版本

			gohello.SayHi("张三") // v2版本的SayHi函数需要传入字符串参数
	*/
	fmt.Println("调用github.com/Meta39/gohello/v2 SayHi()方法")
	gohello2.SayHi("张三")

	/*
		废弃已发布版本
		如果某个发布的版本存在致命缺陷不再想让用户使用时，我们可以使用retract声明废弃的版本。
		例如我们在 github.com/Meta39/gohello/go.mod文件中按如下方式声明即可对外废弃v0.1.2版本：
			module github.com/Meta39/gohello
			go [版本号]
			retract v0.1.2 //废弃已发布版本
		用户使用go get下载v0.1.2版本时就会收到提示，催促其升级到其他版本。
	*/

	fmt.Println("============ 包 ============")
}
