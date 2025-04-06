package main

import (
	"fmt"
	"log"
	"mysql/config"
	"mysql/controller"
	"mysql/router"
	"net/http"
)

/*
Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。
使用database/sql包时必须注入（至少）一个数据库驱动。
以MySQL为例：
1.下载MySQL驱动：go get -u github.com/go-sql-driver/mysql
2.使用MySQL驱动：func Open(driverName, dataSourceName string) (*DB, error)
*/
func main() {
	// 初始化数据库
	err := config.InitDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	defer config.DB.Close()
	// 注册用户资源
	router.RegisterResource("users", &controller.UserHandler{DB: config.DB})

	// 后续添加其他资源（例如角色）
	// RegisterResource("roles", &controller.RoleHandler{})

	// 启动服务
	log.Println("服务器运行在 :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
