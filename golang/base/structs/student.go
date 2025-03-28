package structs

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Student2 学生
type Student2 struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key，即：Gender
	name   string //私有不能被json包访问（即：age序列化和反序列化都无法进行）
}
