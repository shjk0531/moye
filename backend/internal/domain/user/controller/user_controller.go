package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shjk0531/moye/backend/internal/domain/user/model"
	"github.com/shjk0531/moye/backend/internal/domain/user/service"
)

type UserController struct {
	service service.Service
}

func NewUserController(s service.Service) *UserController {
	return &UserController{service: s}
}

// RegisterRoutes는 /users 경로 하위의 라우트를 등록합니다.
func (ctrl *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.POST("/", ctrl.CreateUser)
	r.GET("/:id", ctrl.GetUser)
}

// 사용자 생성
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "사용자 생성 실패"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// 사용자 조회
func (ctrl *UserController) GetUser(c *gin.Context) {
	idStr := c.Param("id")


	user, err := ctrl.service.GetUser(uuid.MustParse(idStr))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "사용자를 찾을 수 없음"})
		return
	}

	c.JSON(http.StatusOK, user)
}
