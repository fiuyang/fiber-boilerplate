package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"scylla/dto"
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

func (handler *DmsHandler) Route(app *fiber.App) {
	app.Get("api/v1/vehicles", handler.GetVehicle)
}

// Note             godoc
//
//	@Summary		Get All vehicles.
//	@Description	Get All vehicles.
//	@Produce		application/json
//	@Param			limit		query	string	false	"limit"
//	@Param			page		query	string	false	"page"
//	@Param			is_active	query	string	false	"is_active"
//	@Tags			vehicle
//	@Success		200	{object}	dto.Response{data=[]dto.VehicleResponse}	"Data"
//	@Router			/vehicles [get]
func (handler *DmsHandler) GetVehicle(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var dataFilter dto.GeneralQueryFilter

	if err := ctx.QueryParser(&dataFilter); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	response, paging, err := handler.dmsService.GetVehicle(c, dataFilter)
	helper.ErrorPanic(err)

	webResponse := dto.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Meta:   &paging,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
