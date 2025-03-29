package sing

import "fmt"

type SuperCat struct {
	Cat // 嵌入 Cat，继承它的Sing、Jump方法
}

// Sing 重写Cat的Sing方法
func (s SuperCat) Sing() {
	fmt.Println("喵喵喵？只有低等猫才会喵喵喵。人类！我是超级猫！我会说人话！你会说猫话吗？")
}
