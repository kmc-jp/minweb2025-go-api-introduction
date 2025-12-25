package service

import (
	"errors"
	"time"

	"part3/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 署名に使用する秘密鍵（本番環境では環境変数から読み込むべきです）
var jwtSecretKey = []byte("your_secret_key")

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
}

type authService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db: db}
}

func (s *authService) Register(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return s.db.Create(&user).Error
}

func (s *authService) Login(username, password string) (string, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                               // Subject (ユーザーID)
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 有効期限 (24時間)
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
