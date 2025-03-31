module golang

go 1.24.1

require github.com/Meta39/gohello v0.1.0 // v0.1.0

require github.com/Meta39/gohello/v2 v2.0.0

require (
	github.com/Meta39/overtime v0.0.0
	golang.org/x/sync v0.12.0
)

replace github.com/Meta39/overtime => ../overtime
