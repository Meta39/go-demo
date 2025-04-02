module golang

go 1.24.1

require github.com/Meta39/gohello v0.1.0 // v0.1.0

require github.com/Meta39/gohello/v2 v2.0.0

require (
	github.com/Meta39/overtime v0.0.0
	github.com/go-sql-driver/mysql v1.9.1
	golang.org/x/sync v0.12.0
)

//在 Go 项目中拉取 MySQL 依赖时，如果同时引入了 filippo.io/edwards25519 包，通常是因为 MySQL 的某些功能（如身份验证插件）依赖了 Ed25519 椭圆曲线算法。
require filippo.io/edwards25519 v1.1.0 // indirect

replace github.com/Meta39/overtime => ../overtime
