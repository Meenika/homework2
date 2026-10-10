package pointers

// Блок pointers закрепляет работу с указателями:
// оператор &, оператор *, nil, изменение значения через указатель,
// возврат указателя и безопасные проверки перед разыменованием.

// ValueOrDefault возвращает значение по указателю или значение по умолчанию.
//
// TODO: если p не nil, верните значение по адресу, даже если оно равно 0. Если p nil, верните def.
func ValueOrDefault(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}

// Increment увеличивает значение по указателю на 1.
//
// TODO: если p не nil, увеличьте исходное значение на 1 и верните новый результат. Для nil верните 0.
func Increment(p *int) int {
	if p == nil {
		return 0
	}
	*p = *p + 1
	return *p
}

// SetValue записывает новое значение по указателю.
//
// TODO: если p не nil, запишите value по адресу и верните true. Для nil верните false.
func SetValue(p *int, value int) bool {
	if p == nil {
		return false
	}
	*p = value
	return true

}

// Swap меняет местами два значения по указателям.
//
// TODO: если оба указателя не nil, поменяйте значения местами и верните true. Иначе ничего не изменяйте и верните false.
func Swap(a, b *int) bool {
	if a != nil && b != nil {
		c := *a
		*a = *b
		*b = c
		return true
	}
	return false
}

// ResetToZero сбрасывает значение по указателю в 0.
//
// TODO: если p не nil, запишите по адресу 0 и верните true. Для nil верните false.
func ResetToZero(p *int) bool {
	if p == nil {
		return false
	}
	*p = 0
	return true
}

// AddToValue прибавляет delta к значению по указателю.
//
// TODO: если p не nil, прибавьте delta к исходному значению и верните результат. Для nil верните 0.
func AddToValue(p *int, delta int) int {
	if p == nil {
		return 0
	}
	*p = *p + delta

	return *p
}

// MaxPointer возвращает указатель на большее значение.
//
// TODO: верните указатель на большее значение; при равенстве верните a. Если один указатель nil, верните другой; если оба nil — nil.
func MaxPointer(a, b *int) *int {
	if a == nil && b == nil {
		return nil
	}
	if a != nil && b == nil {
		return a
	}
	if b != nil && a == nil {
		return b
	}
	if *a > *b {
		return a
	} else if *b > *a {
		return b
	} else {
		return a
	}
}

// IsNil проверяет, равен ли указатель nil.
//
// TODO: верните true только для nil-указателя.
func IsNil(p *int) bool {
	if p == nil {
		return true
	}
	return false
}

// CopyValue возвращает копию значения по указателю.
//
// TODO: верните копию значения по адресу, не изменяя исходную переменную. Для nil верните 0.
func CopyValue(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// DoubleInPlace умножает значение по указателю на 2.
//
// TODO: если p не nil, умножьте исходное значение на 2 и верните true. Для nil верните false.
func DoubleInPlace(p *int) bool {
	if p == nil {
		return false
	}
	*p = *p * 2
	return true
}

// NewInt создаёт новое int-значение и возвращает указатель на него.
//
// TODO: создайте новое int-значение, равное value, и верните ненулевой указатель на него.
func NewInt(value int) *int {
	newValue := value
	return &newValue
}

// DivideInto делит a на b и записывает результат в out.
//
// TODO: если out не nil и b не равен 0, запишите в out результат целочисленного деления a на b и верните true. При ошибке верните false и не меняйте out.
func DivideInto(out *int, a, b int) bool {
	if out != nil && b != 0 {
		*out = a / b
		return true
	}
	return false
}

// ApplyDiscountInPlace применяет скидку к цене по указателю.
//
// TODO: если price не nil и percent >= 0, примените целочисленную скидку и верните true; percent >= 100 должен дать цену 0. Для nil или отрицательного процента верните false и не меняйте цену.
func ApplyDiscountInPlace(price *int, percent int) bool {
	if price == nil || percent < 0 {
		return false
	}

	if percent >= 100 {
		*price = 0
		return true
	}

	*price = *price * (100 - percent) / 100
	return true

}

// ChoosePointer выбирает первый доступный указатель.
//
// TODO: верните primary, если он не nil; иначе верните fallback. Значение по адресу, включая 0, не влияет на выбор.
func ChoosePointer(primary, fallback *int) *int {
	if primary != nil {
		return primary
	}
	return fallback

}

// PointToLarger возвращает указатель на большее значение.
//
// TODO: верните указатель на большее значение; при равенстве верните a. Если один указатель nil, верните другой; если оба nil — nil.
func PointToLarger(a, b *int) *int {
	if a == nil && b == nil {
		return nil
	}
	if a != nil && b == nil {
		return a
	}
	if b != nil && a == nil {
		return b
	}
	if *a > *b {
		return a
	} else if *b > *a {
		return b
	} else {
		return a
	}
}
