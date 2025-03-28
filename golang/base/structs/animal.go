package structs

import "fmt"

// Animal 动物
type Animal struct {
	Name string
}

func (a *Animal) Move() {
	fmt.Printf("%s会行走\n", a.Name)
}
