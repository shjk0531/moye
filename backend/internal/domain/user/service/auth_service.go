package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/repository"
	"github.com/shjk0531/moye/backend/internal/global/config"
	"github.com/shjk0531/moye/backend/internal/global/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (string, error)
	Register(user *model.User) error
	ValidateToken(token string) (uuid.UUID, error)
}

type authService struct {
	repo       repository.Repository
	jwtService jwt.Service
}

func NewAuthService(repo repository.Repository) AuthService {
	jwtConfig := jwt.Config{
		SecretKey:     config.Config.JWT.SecretKey,
		TokenDuration: config.Config.JWT.TokenDuration,
	}
	jwtService := jwt.NewService(jwtConfig)
	
	return &authService{
		repo:       repo,
		jwtService: jwtService,
	}
}

// Login은 이메일과 비밀번호로 사용자를 인증하고 JWT 토큰을 반환
func (s *authService) Login(email, password string) (string, error) {
	// 이메일로 사용자 검색
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("사용자를 찾을 수 없습니다")
	}

	// 비밀번호 검증
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("비밀번호가 일치하지 않습니다")
	}

	// 모델의 ID는 uint 타입이므로 변환해야 합니다.
	// 실제로는 사용자 모델에 UUID를 저장하는 것이 좋습니다.
	userID, err := uuid.Parse(user.ID.String())
	if err != nil {
		return "", errors.New("사용자 ID 변환 오류")
	}

	// JWT 토큰 생성
	token, err := s.jwtService.GenerateToken(userID)
	if err != nil {
		return "", errors.New("토큰 생성 실패")
	}

	return token, nil
}

// Register는 새 사용자를 등록
func (s *authService) Register(user *model.User) error {
	// 비밀번호 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("비밀번호 해싱 실패")
	}
	user.Password = string(hashedPassword)

	// 사용자 저장
	return s.repo.Create(user)
}

// ValidateToken은 토큰의 유효성을 검사하고 사용자 ID를 반환
func (s *authService) ValidateToken(token string) (uuid.UUID, error) {
	return s.jwtService.GetUserID(token)
} 