package text

import (
	"strings"
)

// Блок text закрепляет работу со строками:
// string, byte, rune, immutable string, Unicode,
// пакет strings и простые операции над текстом.

// ByteLen возвращает длину строки в байтах.
//
// TODO: верните размер s в байтах. Для Unicode-строк количество байт может быть больше количества символов.
func ByteLen(s string) int {
	return len(s)
}

// RuneLen возвращает количество Unicode-символов в строке.
//
// TODO: верните количество Unicode-символов в s. Кириллица и emoji должны считаться как отдельные символы.
func RuneLen(s string) int {
	return len([]rune(s))
}

// FirstRune возвращает первый Unicode-символ строки.
//
// TODO: верните первый Unicode-символ s как строку. Для пустой строки верните "".
func FirstRune(s string) string {
	if s == "" {
		return ""
	}
	sRune := []rune(s)
	return string(sRune[0])

}

// LastRune возвращает последний Unicode-символ строки.
//
// TODO: верните последний Unicode-символ s как строку. Для пустой строки верните "".
func LastRune(s string) string {
	if s == "" {
		return ""
	}
	lenS := []rune(s)

	return string(lenS[len(lenS)-1])
}

// Trim убирает пробелы по краям строки.
//
// TODO: верните s без пробельных символов по краям. Внутреннее содержимое строки не изменяйте.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ToLower переводит строку в нижний регистр.
//
// TODO: верните s в нижнем регистре с корректной обработкой латиницы и кириллицы.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper переводит строку в верхний регистр.
//
// TODO: верните s в верхнем регистре с корректной обработкой латиницы и кириллицы.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// NormalizeEmail нормализует email.
//
// TODO: удалите пробельные символы по краям email и приведите всю строку к нижнему регистру.
func NormalizeEmail(email string) string {
	email = strings.TrimSpace(email)
	return strings.ToLower(email)
}

// ContainsWord проверяет, содержит ли text подстроку word.
//
// TODO: верните true, если word встречается внутри text. Поиск чувствителен к регистру; пустая строка word считается найденной.
func ContainsWord(text, word string) bool {
	return strings.Contains(text, word)
}

// ReplaceFirstRune заменяет первый Unicode-символ строки.
//
// TODO: верните новую строку, заменив первый Unicode-символ s на r. Для пустой строки верните "".
func ReplaceFirstRune(s string, r rune) string {
	if s == "" {
		return ""
	}
	sRune := []rune(s)
	sRune[0] = r
	return string(sRune)
}

// ReverseRunes разворачивает строку по Unicode-символам.
//
// TODO: верните s в обратном порядке по Unicode-символам. Кириллица и emoji не должны повреждаться.
func ReverseRunes(s string) string {
	runes := []rune(s)

	if len(runes) <= 1 {
		return s
	}

	return string(runes[len(runes)-1]) + ReverseRunes(string(runes[:len(runes)-1]))
}

// Initials возвращает инициалы имени и фамилии.
//
// TODO: очистите имя и фамилию по краям, возьмите первые Unicode-символы непустых частей, переведите их в верхний регистр и соедините без разделителя.
func Initials(firstName, lastName string) string {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)
	if firstName == "" && lastName != "" {
		lastName = strings.ToLower(lastName)
		LetterBig := []rune(lastName)
		return strings.ToUpper(string(LetterBig[0]))
	} else if lastName == "" && firstName != "" {
		firstName = strings.ToLower(firstName)
		LetterBig := []rune(firstName)
		return strings.ToUpper(string(LetterBig[0]))
	} else {
		firstName = strings.ToLower(firstName)
		lastName = strings.ToLower(lastName)
		LetterBig1 := []rune(firstName)
		LetterBig2 := []rune(lastName)

		return strings.ToUpper(string(LetterBig1[0])) + strings.ToUpper(string(LetterBig2[0]))
	}
}

// RepeatWord повторяет слово count раз.
//
// TODO: верните word, повторённое count раз без разделителя. При count <= 0 верните "".
func RepeatWord(word string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(word, count)
}

// JoinWithComma объединяет строки через запятую.
//
// TODO: соедините элементы values через запятую без дополнительных пробелов. Для пустого или nil-среза верните "".
func JoinWithComma(values []string) string {
	if values == nil {
		return ""
	}
	return strings.Join(values, ",")
}

// IsPalindrome проверяет, является ли строка палиндромом.
//
// TODO: верните true, если s читается одинаково в обоих направлениях после удаления пробелов и приведения к одному регистру. Поддержите Unicode.
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
