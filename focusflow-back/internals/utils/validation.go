package utils

import (
	"errors"
	"unicode"
)

// Проверка пароля
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("Пароль должен быть ≥6 символов")
	}

	var hasUpper, hasDigit bool
	for _, r := range password {
		if unicode.IsUpper(r) {
			hasUpper = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	if !hasUpper || !hasDigit {
		return errors.New("Пароль должен содержать хотя бы одну цифру и одну заглавную букву")
	}

	return nil
}

// Проверка имени пользователя
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("Имя пользователя должно быть ≥3 символов")
	}
	return nil
}

// Проверка роли
func ValidateRole(role string) error {
	validRoles := map[string]bool{
		"user":       true,
		"admin":      true,
		"superadmin": true,
	}
	if role == "" {
		return nil // пустая роль допустима, позже заменим на "user"
	}
	if !validRoles[role] {
		return errors.New("Недопустимая роль. Доступные роли: user, admin, superadmin")
	}
	return nil
}
