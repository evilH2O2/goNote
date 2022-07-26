package customtypeexample

import "fmt"

func Customtypeexample() {
	type errCode int

	// 属于不同类型，混用需要类型转换
	// 同类型不需要
	var notFountCode errCode = errCode(404)

	var code500 int = 500
	// 不同类型，需要类型转换
	var downErrorCode errCode = errCode(code500)

	fmt.Printf("notFoundCode Type => %T\n", notFountCode)
	fmt.Printf("downErrorCode Type => %T\n", downErrorCode)

	// 可以混用的情况，但这只是类型别名，而不是新的类型
	type successCode = int
	var code200 int = 200

	var getOk successCode = code200
	fmt.Printf("successCode Type => %T\n", getOk)
}
