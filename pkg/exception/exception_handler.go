package exception

import (
	"fmt"
	"scylla/dto"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ExceptionHandlers(ctx *fiber.Ctx, err error) error {
	if notFoundError(ctx, err) {
		return nil
	} else if validationError(ctx, err) {
		return nil
	} else if excelValidation(ctx, err) {
		return nil
	} else if badRequestError(ctx, err) {
		return nil
	} else if unauthorizedError(ctx, err) {
		return nil
	} else {
		internalServerError(ctx, err)
		return nil
	}
}

func validationError(ctx *fiber.Ctx, err interface{}) bool {

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		report := make(map[string]string)
		var fieldName string

		for _, e := range castedObject {
			if len(e.Namespace()) > 0 && unicode.IsUpper(rune(e.Namespace()[0])) {
				dotIndex := strings.Index(e.Namespace(), ".")
				if dotIndex != -1 {
					fieldName = e.Namespace()[dotIndex+1:]
				}
			} else {
				fieldName = e.Field()
			}
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
				report[fieldName] = fmt.Sprintf("%s has already been taken", fieldName)
			case "max":
				report[fieldName] = fmt.Sprintf("%s value must be lower than %s", fieldName, e.Param())
			case "min":
				report[fieldName] = fmt.Sprintf("%s value must be greater than %s", fieldName, e.Param())
			case "numeric":
				report[fieldName] = fmt.Sprintf("%s value must be numeric", fieldName)
			case "number":
				report[fieldName] = fmt.Sprintf("%s value must be number", fieldName)
			case "oneof":
				report[fieldName] = fmt.Sprintf("%s value must be %s", fieldName, e.Param())
			case "len":
				report[fieldName] = fmt.Sprintf("%s value must be exactly %s characters long", fieldName, e.Param())
			case "alphanum":
				report[fieldName] = fmt.Sprintf("%s value must be char and numeric %s", fieldName, e.Param())
			case "sliceString":
				report[fieldName] = fmt.Sprintf("%s value ​​in the array cannot be empty is string", fieldName)
			case "dive":
				report[fieldName] = fmt.Sprintf("%s value ​​in the array cannot be empty", fieldName)
			case "datetime":
				report[fieldName] = fmt.Sprintf("%s value must be date (yyyy-mm-dd)", fieldName)
			case "required_if":
				report[fieldName] = fmt.Sprintf("%s must be filled in if %s", fieldName, e.Param())
			case "sliceInt":
				report[fieldName] = fmt.Sprintf("%s value ​​in the array cannot be empty is int", fieldName)
			case "equal":
				report[fieldName] = fmt.Sprintf("%s and %s do not match do not match", fieldName, e.Param())
			case "image":
				report[fieldName] = fmt.Sprintf("%s file must be of type jpg, jpeg, png", fieldName)
			case "base64Image":
				report[fieldName] = fmt.Sprintf("%s value must be base64 encoded image", fieldName)
			}
		}

		ctx.Status(fiber.StatusBadRequest).JSON(dto.Error{
			Code:    fiber.StatusBadRequest,
			Status:  "BAD REQUEST",
			Errors:  report,
			TraceID: ctx.Locals("requestid").(string),
		})
		return true
	}
	return false
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*NotFoundErrorStruct)
	if ok {
		ctx.Status(fiber.StatusNotFound).JSON(dto.Error{
			Code:    fiber.StatusNotFound,
			Status:  "NOT FOUND",
			Errors:  exception.Error(),
			TraceID: ctx.Locals("requestid").(string),
		})
		return true
	}
	return false
}

func badRequestError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*BadRequestErrorStruct)
	if ok {
		ctx.Status(fiber.StatusBadRequest).JSON(dto.Error{
			Code:    fiber.StatusBadRequest,
			Status:  "BAD REQUESTsss",
			Errors:  exception.Error(),
			TraceID: ctx.Locals("requestid").(string),
		})
		return true
	}
	return false
}

func excelValidation(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*ExcelValidation)
	if ok {
		ctx.Status(fiber.StatusBadRequest).JSON(dto.Error{
			Code:    fiber.StatusBadRequest,
			Status:  "BAD REQUEST",
			Errors:  exception.Errors,
			TraceID: ctx.Locals("requestid").(string),
		})
		return true
	}
	return false
}

func unauthorizedError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*UnauthorizedErrorStruct)
	if ok {
		ctx.Status(fiber.StatusUnauthorized).JSON(dto.Error{
			Code:    fiber.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Errors:  exception.Error(),
			TraceID: ctx.Locals("requestid").(string),
		})
		return true
	}
	return false
}

func internalServerError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(*InternalServerErrorStruct)
	if ok {
		ctx.Status(fiber.StatusInternalServerError).JSON(dto.Error{
			Code:    fiber.StatusInternalServerError,
			Status:  "INTERNAL SERVER ERROR",
			Errors:  exception.Error(),
			TraceID: ctx.Locals("requestid").(string),
		})
		return true
	}
	return false

}
