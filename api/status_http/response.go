package statushttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        any    `json:"data"`
}

func BadRequest(c *gin.Context, data any) {
	c.JSON(
		http.StatusBadRequest,
		Response{
			Status:      BadRequestStatus.Status,
			Description: BadRequestStatus.Description,
			Data:        data,
		},
	)
}

func InternalServerError(c *gin.Context, data any) {
	c.JSON(
		http.StatusInternalServerError,
		Response{
			Status:      InternalServerErrorStatus.Status,
			Description: InternalServerErrorStatus.Description,
			Data:        data,
		},
	)
}

func OK(c *gin.Context, data any) {
	c.JSON(
		http.StatusOK,
		Response{
			Status:      OkStatus.Status,
			Description: OkStatus.Description,
			Data:        data,
		},
	)
}

func Created(c *gin.Context, data any) {
	c.JSON(
		http.StatusCreated,
		Response{
			Status:      CreatedStatus.Status,
			Description: CreatedStatus.Description,
			Data:        data,
		},
	)
}

func NoContent(c *gin.Context) {
	c.JSON(
		http.StatusNoContent,
		Response{
			Status:      NoContentStatus.Status,
			Description: NoContentStatus.Description,
			Data:        nil,
		},
	)
}

func TooManyRequests(c *gin.Context, data any) {
	c.JSON(
		http.StatusTooManyRequests,
		Response{
			Status:      TooManyRequestsStatus.Status,
			Description: TooManyRequestsStatus.Description,
			Data:        data,
		},
	)
}

func Unauthorized(c *gin.Context, data any) {
	c.JSON(
		http.StatusUnauthorized,
		Response{
			Status:      UnauthorizedStatus.Status,
			Description: UnauthorizedStatus.Description,
			Data:        data,
		},
	)
}

func Forbidden(c *gin.Context, data any) {
	c.JSON(
		http.StatusForbidden,
		Response{
			Status:      ForbiddenStatus.Status,
			Description: ForbiddenStatus.Description,
			Data:        data,
		},
	)
}

func InvalidArgument(c *gin.Context, data any) {
	c.JSON(
		http.StatusBadRequest,
		Response{
			Status:      InvalidArgumentStatus.Status,
			Description: InvalidArgumentStatus.Description,
			Data:        data,
		},
	)
}

func BadEnvironment(c *gin.Context, data any) {
	c.JSON(
		http.StatusBadRequest,
		Response{
			Status:      BadEnvironmentStatus.Status,
			Description: BadEnvironmentStatus.Description,
			Data:        data,
		},
	)
}
