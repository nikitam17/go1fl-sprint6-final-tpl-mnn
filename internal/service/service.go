package service

func IsMorseCode(s string) bool {
	// Проверяем наличие букв русского алфавита и цифр в строке
	for _, char := range s {
		if (char >= 'А' && char <= 'Я') || (char >= 'а' && char <= 'я') || (char >= '0' && char <= '9') {
			return false // Если найден символ, отличный от точек и тире, это не Морзе
		}
	}
	return true
}
