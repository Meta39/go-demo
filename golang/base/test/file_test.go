package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

/*
file文件操作
*/
func TestFile(t *testing.T) {
	name := "../../log.log"
	//读取
	readFile(name, 512)     //Read 读取指定字节的数据
	readFileFor(name, 1024) //Read for 循环读取全部数据，并设置每次读取的多少字节的数据
	bufioReadFile(name)     //bufio 读取文件
	readFileAll(name)       //os.ReadFile 读取整个文件
	//写入
	writeFile(name)       //Write和WriteString 写文件操作
	directWriteFile(name) //os.WriteFile直接向文件写入指定内容
	bufioNewWriter(name)  //bufio.NewWriter 写文件，先将数据先写入缓存，再将缓存中的内容写入文件
	//练习
	//拷贝文件
	copyFileName := "../../copyLog.log"
	_, err := copyFile(name, copyFileName)
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")
}

/*
打开和关闭文件【使用 defer 确保关闭（推荐）】

必须关闭文件的场景
1.写入文件时，如果不关闭文件，缓冲区可能不会完全刷新（Flush），导致数据丢失或文件内容不完整。
2.长期运行的程序，如果频繁打开文件但不关闭，会导致系统文件描述符耗尽，引发 too many open files 错误。
3.加锁或独占访问时，某些操作系统会保持文件锁，直到文件关闭，不关闭会导致其他进程无法访问。

自动关闭的情况（但仍建议显式关闭）
1.使用 defer 延迟关闭是常见做法，确保函数退出时执行。
2.某些高阶操作（如 ioutil.ReadFile 或 ioutil.WriteFile）在内部会自动关闭文件，但直接操作 os.File 时仍需手动关闭。

不关闭的风险
1.资源泄漏：文件描述符不会被释放。
2.数据不一致：写入时缓冲区未刷新。
3.并发问题：其他进程可能无法访问文件。
*/
func readFile(name string, size int) {
	// 打开log.log文件
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
	}
	defer file.Close() // 函数返回前执行

	// 使用Read方法读取数据，它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF。
	var tmp = make([]byte, size)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据，【读取有限，可以使用for循环读取全部数据】\n", n)
	fmt.Println(string(tmp[:n]))
	return
}

// 循环读取，使用for循环读取文件中的所有数据。
func readFileFor(name string, size int) {
	fmt.Println("for循环读取文件的全部数据.")
	// 打开log.log文件
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, size)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println("for循环读取文件的全部数据为：", string(content))
}

// bufio读取文件 bufio 是在 file 的基础上封装了一层API，支持更多的功能。
func bufioReadFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // MaxScanTokenSize = 64 * 1024
	for scanner.Scan() {
		fmt.Println("bufio读取全部数据：", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("read file failed, err:", err)
	}
}

// ReadFile函数能够读取完整的文件，只需要将文件名作为参数传入。
func readFileAll(name string) {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println("ReadFile读取整个文件数据：", string(content))
}

/*
写入文件时显式关闭
os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。

	func OpenFile(name string, flag int, perm FileMode) (*File, error) {
		...
	}

其中，name：要打开的文件名 flag：打开文件的模式。
模式				含义
os.O_WRONLY		只写
os.O_CREATE		创建文件
os.O_RDONLY		只读
os.O_RDWR		读写
os.O_TRUNC		清空
os.O_APPEND		追加

perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
*/
func writeFile(name string) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
	}
	defer file.Close()

	str := "Hello, Go!"
	//Write和WriteString
	file.Write([]byte(str)) //写入字节切片数据
	file.WriteString(str)   //直接写入字符串数据
	return                  // defer 会在此后执行 Close()
}

// 直接写入文件
func directWriteFile(name string) {
	str := "你好，Meta39."
	err := os.WriteFile(name, []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

// bufio写文件，先将数据先写入缓存，再将缓存中的内容写入文件
func bufioNewWriter(name string) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("嘿嘿，\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

/*
拷贝文件
srcName:源文件名称
dstName:目标文件名称
*/
func copyFile(srcName, dstName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcName, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}
