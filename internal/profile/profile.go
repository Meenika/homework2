package profile

import (
	"fmt"
	"strings"
)

// Блок profile — объединяющая задача.
// Здесь нужно закрепить:
// string normalization, int, bool, if,
// форматирование строки и простую бизнес-логику.

// BuildUserCard собирает карточку пользователя.
//
// TODO: верните строку "name=<name> age=<age> group=<group> status=<status>". Имя очистите по краям, приведите к нижнему регистру и сделайте первый символ заглавным; пустое имя замените на "Unknown". Возраст от 18 — adult, иначе minor; active=true — active, иначе inactive.
func BuildUserCard(name string, age int, active bool) string {
	group := "minor"
	status := "inactive"

	name = strings.TrimSpace(name)
	if name == "" {
		name = "Unknown"
	} else {
		runes := []rune(strings.ToLower(name))
		runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
		name = string(runes)
	}

	if age >= 18 {
		group = "adult"
	}

	if active == true {
		status = "active"
	}

	return fmt.Sprintf("name=%s age=%d group=%s status=%s", name, age, group, status)
}
