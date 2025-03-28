package structs

import "fmt"

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承，因此Dog拥有了Animal.Name的属性和Move()方法
}

func (d *Dog) Wang() {
	fmt.Printf("%s会狗叫\n", d.Name)
}
