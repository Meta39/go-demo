package model

import (
	"database/sql"
	"mysql/config"
	"time"
)

/*
User 用户
MySQL 日期时间类型 vs Go 类型
MySQL 		类型	Go 类型 (database/sql)		说明
DATE		time.Time						仅日期（YYYY-MM-DD）
TIME		time.Time						时间（HH:MM:SS）
DATETIME	time.Time						日期时间（YYYY-MM-DD HH:MM:SS）
TIMESTAMP	time.Time						时间戳（自动转换时区）
*/
type User struct {
	//主键
	Id int `json:"id"`
	//创建时间
	CreatedAt time.Time `json:"createdAt"`
	//创建者
	CreatedBy string `json:"createdBy"`
	//修改时间
	UpdateAt time.Time `json:"updateAt"`
	//修改者
	UpdateBy string `json:"updateBy"`
	//用户名
	Username string `json:"username"`
	//密码
	Password string `json:"password"`
	//姓名
	Name string `json:"name"`
}

var baseSql = "id, created_at, created_by, update_at, update_by, username, password, name"
var noIdSql = "created_at, created_by, update_at, update_by, username, password, name"

// Create 创建用户（包含事务控制）
func (u *User) Create(tx *sql.Tx) (int64, error) {
	result, err := tx.Exec(
		"INSERT INTO users ("+noIdSql+") VALUES (?, ?, ?, ?, ?, ?, ?)",
		u.CreatedAt, u.CreatedBy, u.UpdateAt, u.UpdateBy, u.Username, u.Password, u.Name,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// FindByID 获取用户
func FindByID(id int) (*User, error) {
	u := &User{}
	//sqlStr := fmt.Sprintf("select id, name, age from user where name='%v'", id)//这种会有SQL注入问题。而应该使用预编译?代替这种方式
	err := config.DB.QueryRow(
		"SELECT "+baseSql+" FROM users WHERE id = ?", id,
	).Scan(&u.Id, &u.CreatedAt, &u.CreatedBy, &u.UpdateAt, &u.UpdateBy, &u.Username, &u.Password, &u.Name)
	return u, err
}

// Update 更新用户（包含事务控制）
func (u *User) Update(tx *sql.Tx) error {
	_, err := tx.Exec(
		"UPDATE users SET created_at = ?, created_by = ?, update_at = ?, update_by = ?, username = ?, password = ?, name = ? WHERE id = ?",
		u.CreatedAt, u.CreatedBy, u.UpdateAt, u.UpdateBy, u.Username, u.Password, u.Name, u.Id,
	)
	return err
}

// Delete 删除用户（包含事务控制）
func Delete(tx *sql.Tx, id int) error {
	_, err := tx.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
