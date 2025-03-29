package sing

import "fmt"

// Dog 狗
type Dog struct{}

// Sing 唱（Bird、Dog都是值接收者，Cat是指针接收者）
func (d Dog) Sing() {
	fmt.Println("汪汪汪")
}

// Jump 跳
func (d Dog) Jump() {
	fmt.Println("狗跳")
}
