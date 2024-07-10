package user

import (
	"application/internal/infra"
	"application/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router      gin.IRouter
	UserService *UserService
}

func NewUserRouter(server *infra.Server, userService *UserService) *UserRouter {
	userRouter := &UserRouter{
		router:      server.Router.Group("/users"),
		UserService: userService,
	}
	userRouter.router.GET("/", userRouter.Get)
	userRouter.router.POST("/", userRouter.Post)

	return userRouter
}

func (r *UserRouter) Get(c *gin.Context) {
	users, err := r.UserService.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (r *UserRouter) Post(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = r.UserService.Add(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
