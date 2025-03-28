package structs

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

// User2 用户结构体
type User2 struct {
	Name    string
	Gender  string
	Address //匿名字段
}

// User3 用户结构体
type User3 struct {
	Name     string
	Gender   string
	Address2 //匿名字段和 Email.CreateTime 冲突
	Email    //匿名字段和 Address2.CreateTime 冲突
}
