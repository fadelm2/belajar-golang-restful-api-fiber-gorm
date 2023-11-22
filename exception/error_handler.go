package exception

import (
	"log"
	"restful-api-gorm-fiber/data/response"
)

import "github.com/gofiber/fiber/v2"

type Error string

const (
	ErrPermissionNotAllowed = Error("Permission not allowed")
	ErrUnauthorized         = Error("You're Unauthorized")
	ErrNotFound             = Error("not found")
	ErrInternal             = Error("internal server Error")
	ErrBadRequest           = Error("bad request")
)

func (e Error) Error() string {
	return string(e)
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Println(err)
	_, ok := err.(ValidationError)
	if ok {
		log.Println(err)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
			Code:    fiber.StatusUnprocessableEntity,
			Status:  "Error Validation",
			Message: err.Error(),
		})

	}

	if err == ErrPermissionNotAllowed {
		log.Println(err)
		return ctx.Status(fiber.StatusForbidden).JSON(response.Response{
			Code:    fiber.StatusForbidden,
			Status:  "Error Permission",
			Message: err.Error(),
		})
	}

	if err == ErrInternal {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}

	if err == ErrBadRequest {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Response{
			Code:    fiber.StatusBadRequest,
			Status:  "Error Bad Request",
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(response.Response{
		Code:    fiber.StatusUnprocessableEntity,
		Status:  err.Error(),
		Message: err.Error(),
	})

}
