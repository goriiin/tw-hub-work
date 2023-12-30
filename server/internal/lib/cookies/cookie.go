package cookies

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/patrickmn/go-cache"
	"log/slog"
	"net/http"
	"time"
	"twit-hub111/internal/domain"
)

type CacheService struct {
	appCache *cache.Cache
	log      *slog.Logger
}

func NewCacheService(appCache *cache.Cache, log *slog.Logger) *CacheService {
	return &CacheService{
		appCache: appCache,
		log:      log,
	}
}

func (c *CacheService) SetTokenCookie(w http.ResponseWriter, token string, duration time.Duration) {
	// Создаем cookie с именем "token" и значением, равным сгенерированному токену

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(duration),
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	// Устанавливаем cookie в браузере
	http.SetCookie(w, &cookie)
}

func (c *CacheService) DeleteExpiredToken(w http.ResponseWriter, r *http.Request) {
	// Получаем куки из запроса
	cookie, err := r.Cookie("token")
	if err != nil {
		return
	}

	// Проверяем, истек ли срок действия куки
	if cookie.Expires.Before(time.Now()) {
		// Удаляем куки
		cookie = &http.Cookie{
			Name:   "token",
			Value:  "",
			Path:   "/",
			MaxAge: 3600,
		}
		http.SetCookie(w, cookie)
	}
}

func (c *CacheService) DelCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my_token")
	if err != nil {
		return
	}
	cookie = &http.Cookie{
		Name:     "my_token",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

func (c *CacheService) GetUserIdFromToken(tokenString string) (int, error) {
	// Расшифровываем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Возвращаем секретный ключ
		return []byte(domain.NewApp().Secret), nil
	})
	if err != nil {
		return -1, err
	}

	// Извлекаем userid из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, fmt.Errorf("invalid claims format")
	}
	i, ok := claims["uid"].(float64)

	userId := int(i)
	if !ok {
		return -1, fmt.Errorf("invalid uid format")
	}

	return userId, nil
}

func (c *CacheService) IsCookieValid(w http.ResponseWriter, r *http.Request) (bool, error) {
	// Проверяем, что куки не истекло
	cookie, err := r.Cookie("my_token")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cookie)

	if cookie.Expires.Before(time.Now()) {
		return false, nil
	}

	return true, nil
}
