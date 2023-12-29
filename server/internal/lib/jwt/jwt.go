package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"twit-hub111/internal/domain"
)

// NewToken creates new JWT token for given profile and app.
func NewToken(user domain.TokenUser) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["email"] = user.Email

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(domain.NewApp().Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
