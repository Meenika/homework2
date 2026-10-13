package order

import "fmt"

// Блок order — объединяющая задача.
// Здесь нужно закрепить:
// constants/iota, switch, int conversion, bool,
// вычисление цены в копейках и сборку итоговой строки.

// Статусы заказа.
//
// TODO: задайте StatusNew=0, StatusPaid=1 и StatusCanceled=2 как последовательные значения в указанном порядке.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// OrderSummary собирает краткое описание заказа.
//
// TODO: верните строку "status=<status_text> payment=<payment_text> price_kop=<price>". Статусы: new, paid, canceled, неизвестный — unknown; paid=true даёт paid, иначе not_paid; отрицательную цену считайте 0, остальные рубли переведите в копейки.
func OrderSummary(status int, priceRub int, paid bool) string {
	payment := "not_paid"
	var statusString string

	switch status {
	case StatusNew:
		statusString = "new"
	case StatusPaid:
		statusString = "paid"
	case StatusCanceled:
		statusString = "canceled"
	default:
		statusString = "unknown"
	}

	if paid == true {
		payment = "paid"
	}

	if priceRub > 0 {
		priceRub = priceRub * 100
	} else {
		priceRub = 0
	}

	return fmt.Sprintf("status=%s payment=%s price_kop=%d", statusString, payment, priceRub)
}
