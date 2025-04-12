package services

import (
	"Cars/internal/models"
	"errors"
	_ "fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm
	"log"
	_ "os"
	"time"
)

var (
	jwtSecret       = []byte("your-secret-key")
	refreshSecret   = []byte("your_refresh_secret_key")
	accessTokenTTL  = time.Hour * 1
	refreshTokenTTL = time.Hour * 24 * 7
)

type Claims struct {
	UserID uint
	jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Токен на 24 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return DB.Create(user).Error
}

func LoginUser(email, password string) (*models.User, error) {
	var user models.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("неверный пароль")
	}

	return &user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func CreateUser(user models.User) error {
	// Check if user with this email already exists
	var existingUser models.User
	result := DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	user.Password = hashedPassword

	result = DB.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return result.Error
	}

	return nil
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func AuthenticateUser(email, password string) (*TokenResponse, error) {
	var user models.User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("қате email немесе password")
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("қате email немесе password")
	}

	// Generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	refreshTokenString, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}

	// Save refresh token to database
	refreshTokenModel := models.RefreshToken{
		Token:     refreshTokenString,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(refreshTokenTTL),
	}
	DB.Create(&refreshTokenModel)

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func RefreshAccessToken(refreshToken string) (*TokenResponse, error) {
	// Parse refresh token
	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("жарамсыз refresh token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("жарамсыз token claims")
	}

	// Check if refresh token exists in database
	var refreshTokenModel models.RefreshToken
	err = DB.Where("token = ? AND expires_at > ?", refreshToken, time.Now()).First(&refreshTokenModel).Error
	if err != nil {
		return nil, errors.New("refresh token жарамсыз немесе мерзімі өткен")
	}

	// Generate new access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: claims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken,
	}, nil
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func GetUserFromToken(token *jwt.Token) (*models.User, error) {
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("жарамсыз token claims")
	}

	var user models.User
	err := DB.First(&user, claims.UserID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
