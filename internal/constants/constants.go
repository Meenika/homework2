package constants

// Блок constants закрепляет:
// const, iota, enum-like значения, switch,
// статусы, роли, приоритеты, типы событий и простые текстовые маппинги.

// TODO: задайте константу appName со значением "go-homework-2".
const appName = "go-homework-2"

// TODO: задайте константу maxAttempts со значением 3.
const maxAttempts = 3

// Статусы оплаты.
//
// TODO: задайте последовательные статусы StatusNew=0, StatusPaid=1 и StatusCanceled=2 в одном const-блоке.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// Роли пользователя.
//
// TODO: задайте последовательные роли RoleGuest=0, RoleUser=1 и RoleAdmin=2 в одном const-блоке.
const (
	RoleGuest = iota
	RoleUser
	RoleAdmin
)

// Приоритеты задачи.
//
// TODO: задайте PriorityLow=1, PriorityMedium=2 и PriorityHigh=3 в порядке возрастания важности; 0 должен оставаться неизвестным значением.
const (
	PriorityLow = iota + 1
	PriorityMedium
	PriorityHigh
)

// Типы событий.
//
// TODO: задайте EventCreated, EventUpdated и EventDeleted как три различных последовательных значения в указанном порядке.
const (
	EventCreated = iota
	EventUpdated
	EventDeleted
)

// AppName возвращает имя приложения.
//
// TODO: верните значение константы appName — строку "go-homework-2".
func AppName() string {
	return appName
}

// MaxAttempts возвращает максимальное количество попыток.
//
// TODO: верните значение константы maxAttempts — число 3.
func MaxAttempts() int {
	return maxAttempts
}

// StatusText возвращает текстовое представление статуса.
//
// TODO: верните "new" для StatusNew, "paid" для StatusPaid, "canceled" для StatusCanceled и "unknown" для любого другого значения.
func StatusText(status int) string {
	switch status {
	case StatusNew:
		return "new"
	case StatusPaid:
		return "paid"
	case StatusCanceled:
		return "canceled"
	default:
		return "unknown"
	}
}

// IsFinalStatus проверяет, является ли статус финальным.
//
// TODO: верните true для StatusPaid и StatusCanceled. Для StatusNew и неизвестных значений верните false.
func IsFinalStatus(status int) bool {
	if status == StatusPaid || status == StatusCanceled {
		return true
	} else {
		return false
	}
}

// NextStatus возвращает следующий статус.
//
// TODO: для StatusNew верните StatusPaid; StatusPaid и StatusCanceled оставьте без изменений; для неизвестного значения верните StatusNew.
func NextStatus(status int) int {
	switch status {
	case StatusNew:
		return StatusPaid
	case StatusPaid:
		return StatusPaid
	case StatusCanceled:
		return StatusCanceled
	default:
		return StatusNew
	}
}

// RoleText возвращает текстовое представление роли.
//
// TODO: верните "guest" для RoleGuest, "user" для RoleUser, "admin" для RoleAdmin и "unknown" для любого другого значения.
func RoleText(role int) string {
	switch role {
	case RoleGuest:
		return "guest"
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}

// CanEdit проверяет, может ли пользователь редактировать данные.
//
// TODO: верните true только для RoleAdmin. Для RoleGuest, RoleUser и неизвестных ролей верните false.
func CanEdit(role int) bool {
	switch role {
	case RoleAdmin:
		return true
	default:
		return false
	}
}

// HTTPStatusText возвращает текст HTTP-статуса.
//
// TODO: верните "OK" для 200, "Created" для 201, "Bad Request" для 400, "Not Found" для 404 и "Unknown" для остальных кодов.
func HTTPStatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	default:
		return "Unknown"
	}
}

// DayType возвращает тип дня недели.
//
// TODO: верните "working" для дней 1-5, "weekend" для 6-7 и "unknown" для любого другого номера. В этой задаче 1 — понедельник, 7 — воскресенье.
func DayType(day int) string {
	if day >= 1 && day <= 5 {
		return "working"
	} else if day >= 6 && day <= 7 {
		return "weekend"
	} else {
		return "unknown"
	}
}

// PriorityText возвращает текстовое представление приоритета.
//
// TODO: верните "low" для PriorityLow, "medium" для PriorityMedium, "high" для PriorityHigh и "unknown" для любого другого значения.
func PriorityText(priority int) string {
	switch priority {
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	default:
		return "unknown"
	}
}

// IsKnownStatus проверяет, известен ли статус.
//
// TODO: верните true только для StatusNew, StatusPaid и StatusCanceled; для всех остальных значений верните false.
func IsKnownStatus(status int) bool {
	if status == StatusNew || status == StatusPaid || status == StatusCanceled {
		return true
	}
	return false
}

// PaymentStateText возвращает текст состояния оплаты.
//
// TODO: если canceled=true, верните "canceled" независимо от paid; иначе при paid=true верните "paid", а при обоих false — "pending".
func PaymentStateText(paid, canceled bool) string {
	if canceled {
		return "canceled"
	}
	if paid {
		return "paid"
	}
	return "pending"

}

// TrafficLightAction возвращает действие по цвету светофора.
//
// TODO: верните "stop" для "red", "wait" для "yellow", "go" для "green" и "unknown" для остальных строк.
func TrafficLightAction(color string) string {
	switch color {
	case "red":
		return "stop"
	case "yellow":
		return "wait"
	case "green":
		return "go"
	default:
		return "unknown"
	}

}

// GradeText возвращает текстовую оценку по score.
//
// TODO: верните "invalid" вне диапазона 0-100; "excellent" для 90-100; "good" для 75-89; "passed" для 50-74; "retry" для 0-49.
func GradeText(score int) string {
	if score >= 0 && score <= 49 {
		return "retry"
	} else if score >= 50 && score <= 74 {
		return "passed"
	} else if score >= 75 && score <= 89 {
		return "good"
	} else if score >= 90 && score <= 100 {
		return "excellent"
	} else {
		return "invalid"
	}
}

// EventTypeText возвращает текстовое представление типа события.
//
// TODO: верните "created" для EventCreated, "updated" для EventUpdated, "deleted" для EventDeleted и "unknown" для любого другого значения.
func EventTypeText(eventType int) string {
	switch eventType {
	case EventCreated:
		return "created"
	case EventUpdated:
		return "updated"
	case EventDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}
