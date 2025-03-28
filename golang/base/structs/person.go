package structs

type Person struct {
	Name   string
	Age    int8
	Dreams []string
}

/*
SetDreams 共享底层数组、受外部修改影响、不安全（共享变量）
如果需要data2赋值完p11里面的Dreams内容，后续修改data2，同时也要修改Dreams，推荐使用SetDreams
*/
func (p *Person) SetDreams(dreams []string) {
	p.Dreams = dreams
}

/*
SetDreams2 分配新数组、不受外部修改影响、安全（独立各自的变量）
如果需要data3赋值完p12里面的Dreams内容不受影响，推荐使用SetDreams2
*/
func (p *Person) SetDreams2(dreams []string) {
	p.Dreams = make([]string, len(dreams))
	copy(p.Dreams, dreams)
}
