package basesbytes

import (
	"fmt"
	"strconv"
)

// Блок basesbytes закрепляет системы счисления,
// byte как alias для uint8, форматирование чисел
// и простые операции с ASCII-байтами.
//
// В этом файле важно не просто получить правильные числа,
// а потренироваться записывать и выводить их в разных системах счисления.

// Decimal42 возвращает число 42 в десятичной записи.
//
// TODO: верните число 42, записав его обычным десятичным литералом.
func Decimal42() int {
	return 42
}

// Hex2A возвращает число 42 в шестнадцатеричной записи.
//
// TODO: верните число 42, записав его шестнадцатеричным литералом.
func Hex2A() int {
	return 42
}

// Binary101010 возвращает число 42 в двоичной записи.
//
// TODO: верните число 42, записав его двоичным литералом.
func Binary101010() int {
	return 42
}

// Octal52 возвращает число 42 в восьмеричной записи.
//
// TODO: верните число 42, записав его восьмеричным литералом.
func Octal52() int {
	return 42
}

// SameNumber проверяет, что 42 можно записать разными способами.
//
// TODO: верните true, только если Decimal42, Hex2A, Binary101010 и Octal52 возвращают одно и то же значение 42.
func SameNumber() bool {
	Decimal := Decimal42()
	Hex := Hex2A()
	Binary := Binary101010()
	Octal := Octal52()
	if Decimal == Hex && Binary == Octal && Decimal == Binary {
		return true
	}
	return false
}

// MaxByte возвращает максимальное значение byte.
//
// TODO: верните максимальное значение типа byte — 255.
func MaxByte() byte {
	return 255
}

// FormatDecimal форматирует число в десятичной системе.
//
// TODO: верните десятичную запись n без префикса, например 42 должно стать строкой "42".
func FormatDecimal(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// FormatBinary форматирует число в двоичной системе.
//
// TODO: верните двоичную запись n без префикса, например 42 должно стать строкой "101010".
func FormatBinary(n uint64) string {
	return strconv.FormatUint(n, 2)
}

// FormatOctal форматирует число в восьмеричной системе.
//
// TODO: верните восьмеричную запись n без префикса, например 42 должно стать строкой "52".
func FormatOctal(n uint64) string {
	return strconv.FormatUint(n, 8)
}

// FormatHex форматирует число в шестнадцатеричной системе.
//
// TODO: верните шестнадцатеричную запись n без префикса и с буквами в нижнем регистре, например 42 должно стать строкой "2a".
func FormatHex(n uint64) string {
	return strconv.FormatUint(n, 16)
}

// FormatAllBases выводит одно число в разных системах счисления.
//
// TODO: верните строку формата "dec=<decimal> bin=<binary> oct=<octal> hex=<hex>". Для n=42 результат: "dec=42 bin=101010 oct=52 hex=2a".
func FormatAllBases(n uint64) string {
	return fmt.Sprintf("dec=%s bin=%s oct=%s hex=%s", FormatDecimal(n), FormatBinary(n), FormatOctal(n), FormatHex(n))
}

// IsASCII проверяет, входит ли byte в ASCII-диапазон.
//
// TODO: верните true для значений от 0 до 127 включительно; для остальных значений byte верните false.
func IsASCII(b byte) bool {
	return b <= 127
}

// ToUpperASCII переводит маленькую ASCII-букву в большую.
//
// TODO: если b содержит строчную латинскую ASCII-букву a-z, верните соответствующую A-Z. Любое другое значение верните без изменений.
func ToUpperASCII(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}

// ToLowerASCII переводит большую ASCII-букву в маленькую.
//
// TODO: если b содержит заглавную латинскую ASCII-букву A-Z, верните соответствующую a-z. Любое другое значение верните без изменений.
func ToLowerASCII(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

// PackTwoBytes объединяет два byte в одно uint16.
//
// TODO: верните uint16, в котором high занимает старшие 8 бит, а low — младшие. Например, high=1 и low=2 должны дать 258.
func PackTwoBytes(high, low byte) uint16 {
	return (uint16(high) << 8) | uint16(low)
}
