package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/shjk0531/moye/backend/internal/domain/user/dto"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/repository"
	"github.com/shjk0531/moye/backend/internal/global/config"
	"github.com/shjk0531/moye/backend/internal/global/constants"
	"github.com/shjk0531/moye/backend/internal/global/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 사용자 정의 에러
var (
	ErrDuplicateEmail = errors.New("이메일 중복")
	ErrInvalidToken   = errors.New("유효하지 않은 토큰")
)

type AuthService interface {
	Login(email, password string) (*dto.LoginResponse, string, error)
	Register(user *model.User) error
	RefreshToken(refreshToken string) (*dto.TokenResponse, string, error)
	ValidateToken(token string) (uuid.UUID, error)
}

type authService struct {
	repo       repository.Repository
	jwtService jwt.Service
}

func NewAuthService(repo repository.Repository) AuthService {
	jwtConfig := jwt.Config{
		AccessTokenSecret:  config.Config.JWT.AccessTokenSecret,
		RefreshTokenSecret: config.Config.JWT.RefreshTokenSecret,
		AccessDuration:     config.Config.JWT.AccessDuration,
		RefreshDuration:    config.Config.JWT.RefreshDuration,
	}
	jwtService := jwt.NewService(jwtConfig)

	return &authService{
		repo:       repo,
		jwtService: jwtService,
	}
}

// Login은 이메일과 비밀번호로 사용자를 인증하고 JWT 토큰을 반환
func (s *authService) Login(email, password string) (*dto.LoginResponse, string, error) {
	// 이메일로 사용자 검색
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("사용자를 찾을 수 없습니다")
	}

	// 비밀번호 검증
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("비밀번호가 일치하지 않습니다")
	}

	// 사용자 ID 변환
	if user.ID == uuid.Nil {
		return nil, "", errors.New("유효하지 않은 사용자 ID")
	}

	// 사용자 역할 및 권한 조회
	userWithRoles, err := s.repo.FindUserWithRolesAndPermissions(user.ID)
	if err != nil {
		return nil, "", errors.New("사용자 권한 조회 실패")
	}

	// 역할 이름 목록 추출
	roles := make([]string, 0, len(userWithRoles.Roles))
	for _, role := range userWithRoles.Roles {
		roles = append(roles, string(role))
	}

	// 권한 목록을 []string으로 변환
	permStrings := make([]string, len(userWithRoles.Permissions))
	for i, perm := range userWithRoles.Permissions {
		permStrings[i] = string(perm)
	}

	// JWT 토큰 생성
	tokenDetails, err := s.jwtService.GenerateTokenPair(user.ID, roles, permStrings)
	if err != nil {
		return nil, "", errors.New("토큰 생성 실패")
	}

	// ID를 int로 변환 (필요한 경우 실제 타입에 맞게 조정)
	userInfo := dto.User{
		ID:       user.ID,
		Email:    user.Email,
		Nickname: user.Nickname,
		Profile:  user.Profile,
	}

	return &dto.LoginResponse{
		AccessToken: tokenDetails.AccessToken,
		User:        userInfo,
	}, tokenDetails.RefreshToken, nil
}

// Register는 새 사용자를 등록
func (s *authService) Register(user *model.User) error {
	// 비밀번호 해싱
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("비밀번호 해싱 실패")
	}
	user.Password = string(hashedPassword)

	// 기본 역할 설정: User
	user.Roles = pq.StringArray{string(constants.UserRoleUser)}

	// 기본 권한 설정
	permStrings := []string{
		string(constants.UserPermissionStudyCreate),
		string(constants.UserPermissionStudyUpdateOwn),
		string(constants.UserPermissionStudyDeleteOwn),
	}
	user.Permissions = pq.StringArray(permStrings)

	// 사용자 저장
	err = s.repo.Create(user)
	if err != nil {
		// PostgreSQL 이메일 중복 에러 감지
		if strings.Contains(err.Error(), "duplicate key value") &&
			strings.Contains(err.Error(), "idx_users_email") {
			return ErrDuplicateEmail
		}
		return err
	}

	return nil
}

// RefreshToken은 리프레시 토큰을 사용하여 새 액세스 토큰을 발급
func (s *authService) RefreshToken(refreshToken string) (*dto.TokenResponse, string, error) {
	// 리프레시 토큰 검증
	claims, err := s.jwtService.ValidateToken(refreshToken, jwt.RefreshToken)
	if err != nil {
		fmt.Println("리프레시 토큰 검증 오류:", err)
		return nil, "", ErrInvalidToken
	}

	// 사용자 정보 조회
	userWithRoles, err := s.repo.FindUserWithRolesAndPermissions(claims.UserID)
	if err != nil {
		fmt.Println("사용자 정보 조회 실패:", err)
		return nil, "", errors.New("사용자 정보 조회 실패")
	}

	// 역할 이름 목록 추출
	roles := make([]string, 0, len(userWithRoles.Roles))
	for _, role := range userWithRoles.Roles {
		roles = append(roles, string(role))
	}

	// 권한 목록을 []string으로 변환
	permStrings := make([]string, len(userWithRoles.Permissions))
	for i, perm := range userWithRoles.Permissions {
		permStrings[i] = string(perm)
	}

	// 새 토큰 쌍 생성
	tokenDetails, err := s.jwtService.GenerateTokenPair(claims.UserID, roles, permStrings)
	if err != nil {
		return nil, "", errors.New("토큰 생성 실패")
	}

	return &dto.TokenResponse{
		AccessToken: tokenDetails.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(time.Until(time.Unix(tokenDetails.AtExpires, 0)).Seconds()),
	}, tokenDetails.RefreshToken, nil
}

// ValidateToken은 토큰의 유효성을 검사하고 사용자 ID를 반환
func (s *authService) ValidateToken(token string) (uuid.UUID, error) {
	return s.jwtService.GetUserID(token, jwt.AccessToken)
}
