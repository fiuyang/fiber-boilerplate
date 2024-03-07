package exception

import (
	"fiber-boilerplate/data/response"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if notFoundError(ctx, err) {
		return nil
	} else if validationError(ctx, err) {
		return nil
	} else {
		internalServerError(ctx, err)
		return nil
	}
}

func validationError(ctx *fiber.Ctx, err interface{}) bool {

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		report := make(map[string]string)

		for _, e := range castedObject {
			fieldName := e.Field()
			switch e.Tag() {
			case "required":
				report[fieldName] = fmt.Sprintf("%s is required", fieldName)
			case "email":
				report[fieldName] = fmt.Sprintf("%s is not valid email", fieldName)
			case "gte":
				report[fieldName] = fmt.Sprintf("%s value must be greater than %s", fieldName, e.Param())
			case "lte":
				report[fieldName] = fmt.Sprintf("%s value must be lower than %s", fieldName, e.Param())
			case "unique":
				report[fieldName] = fmt.Sprintf("%s has already been taken %s", fieldName, e.Param())
			case "max":
				report[fieldName] = fmt.Sprintf("%s value must be lower than %s", fieldName, e.Param())
			case "min":
				report[fieldName] = fmt.Sprintf("%s value must be greater than %s", fieldName, e.Param())
			case "numeric":
				report[fieldName] = fmt.Sprintf("%s value must be numeric", fieldName)
			case "oneof":
				report[fieldName] = fmt.Sprintf("%s value must be %s", fieldName, e.Param())
			case "len":
				report[fieldName] = fmt.Sprintf("%s value must be exactly %s characters long", fieldName, e.Param())
			}
		}

		ctx.Status(fiber.StatusBadRequest).JSON(response.Error{
			Code:   fiber.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: report,
		})
		return true
	}
	return false
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*NotFoundErrorStruct)
	if ok {
		ctx.Status(fiber.StatusNotFound).JSON(response.Error{
			Code:   fiber.StatusNotFound,
			Status: "NOT FOUND",
			Errors: exception.Error(),
		})
		return true
	}
	return false
}

func internalServerError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*InternalServerErrorStruct)
	if ok {
		ctx.Status(fiber.StatusInternalServerError).JSON(response.Error{
			Code:   fiber.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Errors: exception.Error(),
		})
		return true
	}
	return false

}
