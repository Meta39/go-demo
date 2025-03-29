package interfaces

/*
Payer 支付接口
*/
type Payer interface {
	Pay(amount int64)
}

// Checkout 结账函数
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}
