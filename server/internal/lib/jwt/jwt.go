package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"twit-hub111/internal/domain"
)

// NewToken creates new JWT token for given profile and app.
func NewToken(user domain.TokenUser) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(domain.NewApp().Secret))
	if err != nil {
		fmt.Println("jwt.NewToken", err)
		return "", err
	}

	return tokenString, nil
}
