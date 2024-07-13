package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"scylla/entity"
	"scylla/pkg/exception"
	"scylla/pkg/helper"
	"scylla/pkg/utils"
	"scylla/service"
	"time"
)

type DmsHandler struct {
	dmsService service.DmsService
}

func NewDmsHandler(service service.DmsService) *DmsHandler {
	return &DmsHandler{
		dmsService: service,
	}
}

// Note             godoc
//	@Summary		Get All vehicles.
//	@Description	Get All vehicles.
//	@Produce		application/json
//	@Param			limit		query	string	false	"limit"
//	@Param			page		query	string	false	"page"
//	@Param			is_active	query	string	false	"is_active"
//	@Tags			vehicle
//	@Success		200	{object}	entity.Response{data=[]entity.VehicleResponse}	"Data"
//	@Router			/vehicles [get]
func (handler *DmsHandler) GetVehicle(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var dataFilter entity.GeneralQueryFilter

	if err := ctx.QueryParser(&dataFilter); err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	response, paging, err := handler.dmsService.GetVehicle(c, dataFilter)
	helper.ErrorPanic(err)

	webResponse := entity.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
		Meta:   &paging,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(http.StatusOK).JSON(webResponse)
}
