package boolean

// Блок boolean закрепляет работу с bool:
// сравнения, логические операторы &&, ||, !,
// простые условия и функции, которые отвечают true или false.

// CanEnter проверяет, может ли человек войти.
//
// TODO: верните true, только если age не меньше 18 и hasTicket равен true.
func CanEnter(age int, hasTicket bool) bool {
	if age >= 18 && hasTicket == true {
		return true
	}
	return false
}

// IsAdult проверяет, является ли человек совершеннолетним.
//
// TODO: верните true при age >= 18; для меньшего, нулевого или отрицательного возраста верните false.
func IsAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

// CanBuyAlcohol проверяет, можно ли купить алкоголь.
//
// TODO: в рамках задания верните true при age >= 18, иначе false.
func CanBuyAlcohol(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

// CanRest проверяет, можно ли отдыхать.
//
// TODO: верните true, если isWeekend или isHoliday равен true. Если оба флага false, верните false.
func CanRest(isWeekend, isHoliday bool) bool {
	if isWeekend == true || isHoliday == true {
		return true
	}
	return false
}

// IsWorkingDay проверяет, является ли день рабочим.
//
// TODO: верните true для строк "monday", "tuesday", "wednesday", "thursday" и "friday". Для выходных и неизвестных значений верните false.
func IsWorkingDay(day string) bool {
	switch day {
	case "monday":
		return true

	case "tuesday":
		return true

	case "wednesday":
		return true

	case "thursday":
		return true

	case "friday":
		return true

	default:
		return false
	}
}

// HasAccess проверяет доступ пользователя.
//
// TODO: верните true, если isAdmin или isOwner равен true; без обоих прав верните false.
func HasAccess(isAdmin, isOwner bool) bool {
	if isAdmin == true || isOwner == true {
		return true
	}
	return false
}

// CanApplyDiscount проверяет, можно ли применить скидку.
//
// TODO: верните true для VIP-пользователя либо при total >= 5000. В остальных случаях верните false.
func CanApplyDiscount(isVIP bool, total int) bool {
	if isVIP == true || total >= 5000 {
		return true
	}
	return false
}

// ShouldNotify проверяет, нужно ли отправлять уведомление.
//
// TODO: верните true, только если emailVerified и notificationsEnabled одновременно равны true.
func ShouldNotify(emailVerified, notificationsEnabled bool) bool {
	if emailVerified == true && notificationsEnabled == true {
		return true
	}
	return false
}

// IsValidScore проверяет корректность оценки.
//
// TODO: верните true для score от 0 до 100 включительно; вне диапазона верните false.
func IsValidScore(score int) bool {
	if score >= 0 && score <= 100 {
		return true
	}
	return false
}

// IsInRange проверяет, что value находится в диапазоне [min, max].
//
// TODO: верните true, если value находится между min и max включительно; иначе false.
func IsInRange(value, min, max int) bool {
	if value >= min && value <= max {
		return true
	}
	return false
}

// IsLeapYear проверяет, является ли год високосным.
//
// TODO: верните true, если год кратен 400 либо кратен 4, но не кратен 100. Для остальных лет верните false.
func IsLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// CanWithdraw проверяет, можно ли снять деньги.
//
// TODO: верните true, если аккаунт не заблокирован, amount > 0 и balance >= amount. В остальных случаях верните false.
func CanWithdraw(balance, amount int, blocked bool) bool {
	if blocked == false && amount > 0 && balance >= amount {
		return true
	}
	return false
}

// LoginAllowed проверяет, разрешён ли вход.
//
// TODO: верните true, только если passwordOK и otpOK одновременно равны true.
func LoginAllowed(passwordOK, otpOK bool) bool {
	if passwordOK == true && otpOK == true {
		return true
	}
	return false
}

// IsEmpty проверяет, является ли строка пустой.
//
// TODO: верните true только для строки длиной 0. Пробел, перевод строки и любой другой символ означают непустую строку.
func IsEmpty(text string) bool {
	if text == "" {
		return true
	}
	return false
}

// Not возвращает противоположное bool-значение.
//
// TODO: верните логическое значение, противоположное flag.
func Not(flag bool) bool {
	if flag == true {
		return false
	}
	return true
}
