package sing

import "fmt"

// Bird 我们有一个Bird 结构体类型如下。
type Bird struct{}

/*
Sing 因为Singer接口只包含一个Sing方法，所以只需要给Bird结构体添加一个Sing方法就可以满足Singer接口的要求。
这样就称为Bird实现了Singer接口。（Bird、Dog都是值接收者，Cat是指针接收者）
*/
func (b Bird) Sing() {
	fmt.Println("叽叽喳喳")
}
