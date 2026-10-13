package calculator

import "errors"

// Блок calculator — объединяющая задача.
// Здесь нужно закрепить:
// int, string operation, switch, if, error,
// целочисленное деление и остаток от деления.

// Calculate выполняет арифметическую операцию над двумя int-числами.
//
// TODO: поддержите операции "+", "-", "*", "/" и "%" над a и b. Деление выполняется как целочисленное; нулевой делитель для "/" и "%", пустая или неизвестная операция должны возвращать ошибку.
func Calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Error")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, errors.New("Error")
		}
		return a % b, nil
	default:
		return 0, errors.New("Error")
	}
}
