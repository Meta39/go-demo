package pay

import "fmt"

// ZhiFuBao 支付宝
type ZhiFuBao struct {
}

// Pay 实现支付接口
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}
