package float

import (
	"fmt"
	"math"
)

// Блок float закрепляет работу с float64:
// дробные числа, проценты, округление, форматирование,
// преобразование float64 в int и простые вычисления.

// Add складывает два дробных числа.
//
// TODO: верните сумму a и b.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract вычитает второе дробное число из первого.
//
// TODO: верните результат вычитания b из a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply умножает два дробных числа.
//
// TODO: верните произведение a и b.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide делит первое дробное число на второе.
//
// TODO: верните результат дробного деления a на b. Если b равен 0, верните 0.
func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// DiscountPrice считает цену после скидки.
//
// TODO: верните цену после применения percent процентов скидки. Отрицательный percent не меняет цену, а percent >= 100 даёт результат 0.
func DiscountPrice(price, percent float64) float64 {
	if percent >= 100 {
		return 0
	}
	if percent < 0 {
		return price
	}
	return price * (100 - percent) / 100
}

// AddTax считает цену с налогом.
//
// TODO: верните цену после добавления taxPercent процентов налога. Если taxPercent отрицательный, верните исходную цену.
func AddTax(price, taxPercent float64) float64 {
	if taxPercent < 0 {
		return price
	}
	return price * (100 + taxPercent) / 100
}

// CelsiusToFahrenheit переводит градусы Цельсия в Фаренгейты.
//
// TODO: верните температуру в градусах Фаренгейта. Контрольные значения: 0°C = 32°F, 100°C = 212°F.
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

// FahrenheitToCelsius переводит градусы Фаренгейта в Цельсии.
//
// TODO: верните температуру в градусах Цельсия. Контрольные значения: 32°F = 0°C, 212°F = 100°C.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Average считает среднее арифметическое двух дробных чисел.
//
// TODO: верните среднее арифметическое a и b как float64.
func Average(a, b float64) float64 {

	return (a + b) / 2
}

// Round2 округляет число до 2 знаков после точки.
//
// TODO: верните value, округлённое до двух знаков после десятичной точки. Например, 12.345 должно стать 12.35.
func Round2(value float64) float64 {
	return math.Round(value*100) / 100
}

// FormatPrice форматирует цену с двумя знаками после точки.
//
// TODO: верните строку ровно с двумя знаками после точки и обычным округлением: 12.3 -> "12.30", 12.345 -> "12.35".
func FormatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

// PercentOf считает процент от числа.
//
// TODO: верните percent процентов от total. Отрицательный percent должен давать отрицательный результат.
func PercentOf(total, percent float64) float64 {
	return total * percent / 100
}

// GrowthPercent считает рост в процентах между старым и новым значением.
//
// TODO: верните процентное изменение от oldValue к newValue: рост — положительный, снижение — отрицательное. При oldValue == 0 верните 0.
func GrowthPercent(oldValue, newValue float64) float64 {
	if oldValue == 0 {
		return 0
	}
	return (newValue - oldValue) / oldValue * 100
}

// IsPositive проверяет, что число строго больше нуля.
//
// TODO: верните true только для значения строго больше 0; для нуля и отрицательных значений верните false.
func IsPositive(value float64) bool {
	if value > 0 {
		return true
	}
	return false
}

// FloatToInt преобразует float64 в int.
//
// TODO: верните value с типом int. Дробная часть должна отбрасываться в сторону нуля, например -12.99 превращается в -12.
func FloatToInt(value float64) int {
	return int(value)
}
