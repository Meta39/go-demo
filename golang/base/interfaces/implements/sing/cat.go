package sing

import "fmt"

// Cat 猫
type Cat struct{}

// Sing 唱（Bird、Dog都是值接收者，Cat是指针接收者）
func (c *Cat) Sing() {
	fmt.Println("喵喵喵")
}

func (c *Cat) Jump() {
	fmt.Println("猫跳")
}
