package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

// Блок conversion закрепляет:
// явные преобразования типов, named types,
// strconv, парсинг строк, форматирование значений
// и обработку ошибок при преобразовании.

// Rub — пользовательский тип для рублей.
type Rub int

// Kop — пользовательский тип для копеек.
type Kop int

// UserID — пользовательский тип для идентификатора пользователя.
type UserID int64

// IntToInt64 явно преобразует int в int64.
//
// TODO: верните то же числовое значение с типом int64.
func IntToInt64(n int) int64 {
	return int64(n)
}

// Int64ToInt явно преобразует int64 в int.
//
// TODO: верните то же числовое значение с типом int. Переполнение в этой учебной задаче не обрабатывается.
func Int64ToInt(n int64) int {
	return int(n)
}

// RubToKop переводит рубли в копейки.
//
// TODO: верните количество копеек типа Kop, считая один рубль равным 100 копейкам. Знак исходного значения сохраняется.
func RubToKop(rub Rub) Kop {
	return Kop(rub * 100)
}

// KopToRub переводит копейки в рубли через целочисленное деление.
//
// TODO: верните целое количество рублей типа Rub. Неполные рубли должны отбрасываться по правилам целочисленного деления.
func KopToRub(kop Kop) Rub {
	return Rub(kop / 100)
}

// UserIDToString преобразует UserID в string.
//
// TODO: верните десятичную строку со значением id, включая нулевые и отрицательные значения.
func UserIDToString(id UserID) string {
	return strconv.Itoa(int(id))
}

// ParseInt преобразует строку в int.
//
// TODO: преобразуйте text в int. Для корректной строки верните число и nil, для пустой или некорректной — ошибку.
func ParseInt(text string) (int, error) {
	return strconv.Atoi(text)
}

// ParseAndDouble парсит строку в int и умножает результат на 2.
//
// TODO: преобразуйте text в int и верните удвоенное значение. Если строку нельзя распарсить, верните ошибку.
func ParseAndDouble(text string) (int, error) {
	num, err := ParseInt(text)
	if err != nil {
		return 0, err
	}
	return num * 2, nil
}

// IntToString преобразует int в string.
//
// TODO: верните десятичное строковое представление n.
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// FloatToString форматирует float64 с двумя знаками после точки.
//
// TODO: верните строку ровно с двумя знаками после точки и округлением, например 12.345 -> "12.35".
func FloatToString(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

// ParseBoolText преобразует строку в bool.
//
// TODO: преобразуйте стандартные строки "true", "false", "1" и "0" в bool. Для неподдерживаемого значения верните ошибку.
func ParseBoolText(text string) (bool, error) {

	switch text {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "1":
		return true, nil
	case "0":
		return false, nil
	default:
		return false, errors.New("Not supported value")
	}

}

// BoolToText преобразует bool в string.
//
// TODO: верните строку "true" для true и "false" для false.
func BoolToText(value bool) string {
	if value == true {
		return "true"
	}
	return "false"
}

// SumIntAndInt64 складывает int и int64.
//
// TODO: верните сумму a и b с типом int64, сохранив знаки обоих значений.
func SumIntAndInt64(a int, b int64) int64 {
	return int64(a) + b
}

// PriceRubStringToKop преобразует строку с рублями в копейки.
//
// TODO: распарсите text как целое количество рублей и верните цену типа Kop. Для отрицательной или некорректной цены верните ошибку.
func PriceRubStringToKop(text string) (Kop, error) {
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	if num < 0 {
		return 0, errors.New("Negative price")
	}

	return Kop(num) * 100, nil
}

// SafeParsePositive безопасно парсит положительное число.
//
// TODO: верните распарсенное положительное целое число. Для нуля, отрицательного числа, пустой или некорректной строки верните 0.
func SafeParsePositive(text string) int {
	num, err := strconv.Atoi(text)
	if num <= 0 || err != nil {
		return 0
	}
	return num

}

// FormatUser форматирует пользователя в строку.
//
// TODO: верните строку формата "user:<id>:<name>". Например, id=1001 и name="Maria" дают "user:1001:Maria"; пустое имя допустимо.
func FormatUser(id UserID, name string) string {
	return fmt.Sprintf("user:%d:%s", id, name)
}
