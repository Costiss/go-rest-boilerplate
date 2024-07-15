package user

import (
	"application/internal/infra"
	"application/internal/util/httputil"
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

//	@Summary		List all users
//	@Description	List all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}		UserDTO
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/users [get]
func (r *UserRouter) Get(c *gin.Context) {
	users, err := r.UserService.Get()
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, UserDTO{}.fromModelArr(users))
}

//	@Summary		Create a user
//	@Description	Create a user
//	@Tags			users
//	@Param			request	body	CreateUserDTO	true	"User"
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	UserDTO
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/users [post]
func (r *UserRouter) Post(c *gin.Context) {
	var user CreateUserDTO
	err := c.BindJSON(&user)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	created, err := r.UserService.Add(&user)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, UserDTO{}.fromModel(created))
}
