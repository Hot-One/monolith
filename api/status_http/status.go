package statushttp

import "net/http"

type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var (
	OkStatus = Status{
		Code:        http.StatusOK,
		Status:      "OK",
		Description: "The request has succeeded",
	}

	CreatedStatus = Status{
		Code:        http.StatusCreated,
		Status:      "CREATED",
		Description: "The request has been fulfilled and has resulted in one or more new resources being created",
	}

	NoContentStatus = Status{
		Code:        http.StatusNoContent,
		Status:      "NO_CONTENT",
		Description: "There is no content to send for this request, but the headers may be useful",
	}

	BadEnvironmentStatus = Status{
		Code:        http.StatusBadRequest,
		Status:      "BAD_ENVIRONMENT",
		Description: "The service has an invalid environment value",
	}

	BadRequestStatus = Status{
		Code:        http.StatusBadRequest,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax",
	}

	InvalidArgumentStatus = Status{
		Code:        http.StatusBadRequest,
		Status:      "INVALID_ARGUMENT",
		Description: "Invalid argument value passed",
	}

	UnauthorizedStatus = Status{
		Code:        http.StatusUnauthorized,
		Status:      "UNAUTHORIZED",
		Description: "...",
	}

	ForbiddenStatus = Status{
		Code:        http.StatusForbidden,
		Status:      "FORBIDDEN",
		Description: "...",
	}

	TooManyRequestsStatus = Status{
		Code:        http.StatusTooManyRequests,
		Status:      "TOO_MANY_REQUESTS",
		Description: "The user has sent too many requests in a given amount of time",
	}

	InternalServerErrorStatus = Status{
		Code:        http.StatusInternalServerError,
		Status:      "INTERNAL_SERVER_ERROR",
		Description: "The server encountered an unexpected condition that prevented it from fulfilling the request",
	}

	NotFoundStatus = Status{
		Code:        http.StatusNotFound,
		Status:      "NOT_FOUND",
		Description: "The user not found",
	}
)
