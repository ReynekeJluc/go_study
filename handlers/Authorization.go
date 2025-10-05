package handlers

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	database "github.com/ReynekeJluc/go_study.git/db"
	models "github.com/ReynekeJluc/go_study.git/struct/models"
	jwt "github.com/golang-jwt/jwt/v5"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var (
	TokenSecret = os.Getenv("TOKEN_PEPPER")
	JWTSecret   = os.Getenv("JWT_SECRET")
)


// --- LOGIN ---
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "некорректный json"})
		return
	}

	user, err := Login(req.Login, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	accessToken, err := GenerateAccessToken(user.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "ошибка генерации токена"})
		return
	}

	refreshToken, err := GenerateRefreshToken(user.UserId, 24 * time.Hour)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "ошибка генерации refresh токена"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}


// --- ЧИСТЫЙ SQL ---
func Login(login, password string) (*models.User, error) {
	var user models.User
	row := database.DB.QueryRow(`
		SELECT user_id, login, password
		FROM users
		WHERE login = ?
	`, login)

	err := row.Scan(&user.UserId, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("неверный пароль")
	}

	return &user, nil
}


// --- ACCESS TOKEN ---
func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}


// --- REFRESH TOKEN ---
func GenerateRefreshToken(userID uint, expiresIn time.Duration) (string, error) {
	_, err := database.DB.Exec(`UPDATE tokens SET revoked = 1 WHERE user_id = ? AND revoked = 0`, userID)
	if err != nil {
		return "", err
	}

	rawToken, err := generateRandomToken(32)
	if err != nil {
		return "", err
	}
	tokenHash := hashToken(rawToken, TokenSecret)

	_, err = database.DB.Exec(`
		INSERT INTO tokens (user_id, token_hash, created_at, expires_at, revoked)
		VALUES (?, ?, ?, ?, 0)
	`, userID, tokenHash, time.Now(), time.Now().Add(expiresIn))
	if err != nil {
		return "", err
	}

	return rawToken, nil
}


// --- VALIDATE REFRESH TOKEN ---
func ValidateRefreshToken(rawToken string) (*models.Token, error) {
	tokenHash := hashToken(rawToken, TokenSecret)

	var token models.Token
	row := database.DB.QueryRow(`
		SELECT token_id, user_id, token_hash, created_at, expires_at, revoked
		FROM tokens
		WHERE token_hash = ? AND expires_at > ? AND revoked = 0
	`, tokenHash, time.Now())

	err := row.Scan(&token.TokenId, &token.UserId, &token.TokenHash, &token.CreatedAt, &token.ExpiresAt, &token.Revoked)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("неправильный или истекший токен")
		}
		return nil, err
	}

	return &token, nil
}


// --- REFRESH ACCESS TOKEN ---
func RefreshAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserId       uint    `json:"user_id"`
		RefreshToken string  `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "некорректный json"})
		return
	}

	token, err := ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	if token.UserId != req.UserId {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "неверный refresh токен"})
		return
	}

	accessToken, err := GenerateAccessToken(req.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "ошибка генерации access токена"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
}


// --- LOGOUT ---
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserId int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "некорректный json"})
		return
	}

	_, err := database.DB.Exec(`UPDATE tokens SET revoked = 1 WHERE user_id = ? AND revoked = 0`, req.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "ошибка выхода"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "успешно вышли"})
}


// --- ВСПОМОГАТЕЛЬНЫЕ ---
func generateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}


func hashToken(token, secret string) string {
	m := hmac.New(sha256.New, []byte(secret))
	m.Write([]byte(token))
	return hex.EncodeToString(m.Sum(nil))
}
