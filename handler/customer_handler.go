package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"scylla/dto"
	"scylla/pkg/exception"
	"scylla/pkg/helper"
	"scylla/pkg/utils"
	"scylla/service"
	"time"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (handler *CustomerHandler) Route(app *fiber.App) {
	qParamId := ":customerId"
	customerRouter := app.Group("/api/v1/customers")
	customerRouter.Get("", handler.FindAll)
	customerRouter.Get("/"+qParamId, handler.FindById)
	customerRouter.Get("/export", handler.Export)
	customerRouter.Post("/import", handler.Import)
	customerRouter.Post("", handler.Create)
	customerRouter.Post("/batch", handler.CreateBatch)
	customerRouter.Patch("/"+qParamId, handler.Update)
	customerRouter.Delete("/batch", handler.DeleteBatch)
}

// Note            godoc
//
//	@Summary		Create customer
//	@Description	Create customer.
//	@Param			data	formData	dto.CreateCustomerRequest	true	"create customer"
//	@Produce		application/json
//	@Tags			customers
//	@Success		201	{object}	dto.JsonCreated{data=nil}   "Data"
//	@Failure		400	{object}	dto.JsonBadRequest{}			"Validation error"
//	@Failure		404	{object}	dto.JsonNotFound{}				"Data not found"
//	@Failure		500	{object}	dto.JsonInternalServerError{}	"Internal server error"
//	@Router			/customers [post]
func (handler *CustomerHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := dto.CreateCustomerRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	handler.customerService.Create(c, request)

	webResponse := dto.Response{
		Code:    fiber.StatusCreated,
		Status:  "Created",
		Message: "Created Successful",
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// Note            godoc
//
//	@Summary		Create customer batch
//	@Description	Create customer batch.
//	@Param			data	body	dto.CreateCustomerBatchRequest	true	"create customer batch"
//	@Produce		application/json
//	@Tags			customers
//	@Success		201	{object}	dto.JsonCreated{data=nil}       "Data"
//	@Failure		400	{object}	dto.JsonBadRequest{}			"Validation error"
//	@Failure		404	{object}	dto.JsonNotFound{}				"Data not found"
//	@Failure		500	{object}	dto.JsonInternalServerError{}	"Internal server error"
//	@Router			/customers/batch [post]
func (handler *CustomerHandler) CreateBatch(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := dto.CreateCustomerBatchRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	handler.customerService.CreateBatch(c, request)

	webResponse := dto.Response{
		Code:    fiber.StatusCreated,
		Status:  "Created",
		Message: "Created Batch Successful",
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// Note            godoc
//
//	@Summary		update customer
//	@Description	update customer.
//	@Param			data		body	dto.UpdateCustomerRequest	true	"update customer"
//	@Param			customerId	path	string						true	"customer_id"
//	@Produce		application/json
//	@Tags			customers
//	@Success		200	{object}	dto.JsonSuccess{data=nil}		"Data"
//	@Failure		400	{object}	dto.JsonBadRequest{}			"Validation error"
//	@Failure		404	{object}	dto.JsonNotFound{}				"Data not found"
//	@Failure		500	{object}	dto.JsonInternalServerError{}	"Internal server error"
//	@Router			/customers/{customerId} [patch]
func (handler *CustomerHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := dto.UpdateCustomerRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	var params dto.CustomerParams

	if err := ctx.ParamsParser(&params); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	request.ID = params.CustomerId

	handler.customerService.Update(c, request)

	webResponse := dto.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Message: "Update Successful",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// Note             godoc
//
//	@Summary		Delete batch customer
//	@Description	Delete batch customer.
//	@Param			data	body	dto.DeleteBatchCustomerRequest	true	"delete batch customer"
//	@Produce		application/json
//	@Tags			customers
//	@Success		200	{object}	dto.JsonSuccess{data=nil}		"Data"
//	@Failure		400	{object}	dto.JsonBadRequest{}			"Validation error"
//	@Failure		404	{object}	dto.JsonNotFound{}				"Data not found"
//	@Failure		500	{object}	dto.JsonInternalServerError{}	"Internal server error"
//	@Router			/customers/batch [delete]
func (handler *CustomerHandler) DeleteBatch(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := dto.DeleteBatchCustomerRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	handler.customerService.DeleteBatch(c, request)

	webResponse := dto.Response{
		Code:    fiber.StatusOK,
		Status:  "OK",
		Message: "Delete Batch Successful",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

// Note 		    godoc
//
//	@Summary		get customer by id.
//	@Param			customerId	path	string	true	"customer_id"
//	@Description	get customer by id.
//	@Produce		application/json
//	@Tags			customers
//	@Success		200	{object}	dto.JsonSuccess{data=dto.CustomerResponse{}}    "Data"
//	@Failure		400	{object}	dto.JsonBadRequest{}							"Validation error"
//	@Failure		404	{object}	dto.JsonNotFound{}								"Data not found"
//	@Failure		500	{object}	dto.JsonInternalServerError{}					"Internal server error"
//	@Router			/customers/{customerId} [get]
func (handler *CustomerHandler) FindById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var params dto.CustomerParams

	if err := ctx.ParamsParser(&params); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	data := handler.customerService.FindById(c, params)

	webResponse := dto.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// Note             godoc
//
//	@Summary		Get all customers.
//	@Description	Get all customers.
//	@Produce		application/json
//	@Param			all		    query	string	false	"true/false"
//	@Param			limit		query	string	false	"limit"
//	@Param			page		query	string	false	"page"
//	@Param			username	query	string	false	"username"
//	@Param			email		query	string	false	"email"
//	@Param			start_date	query	string	false	"start_date"
//	@Param			end_date	query	string	false	"end_date"
//	@Param			sort		query	string	false	"sort"
//	@Tags			customers
//	@Success		200	{object}	dto.Response{data=[]dto.CustomerResponse{}}	    "Data"
//	@Failure		400	{object}	dto.JsonBadRequest{}							"Validation error"
//	@Failure		404	{object}	dto.JsonNotFound{}								"Data not found"
//	@Failure		500	{object}	dto.JsonInternalServerError{}					"Internal server error"
//	@Router			/customers [get]
func (handler *CustomerHandler) FindAll(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var dataFilter dto.CustomerQueryFilter

	if err := ctx.QueryParser(&dataFilter); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	response, paging := handler.customerService.FindAll(c, dataFilter)

	webResponse := dto.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
		Meta:   &paging,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// Note 		    godoc
//
// @Summary		Export Excel customer.
// @Description	Export Excel customer.
// @Produce		application/json
// @Tags		customers
// @Param		all     	query		string	true	"true"
// @Param		start_date	query		string	false	"start_date"
// @Param		end_date	query		string	false	"end_date"
// @Param		username	query		string	false	"username"
// @Param		email		query		string	false	"email"
// @Success		200			{object}	dto.JsonSuccess{data=string}    "Data"
// @Failure		400			{object}	dto.JsonBadRequest{}			"Validation error"
// @Failure		404			{object}	dto.JsonNotFound{}				"Data not found"
// @Failure		500			{object}	dto.JsonInternalServerError{}	"Internal server error"
// @Router		/customers/export [get]
func (handler *CustomerHandler) Export(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var dataFilter dto.CustomerQueryFilter

	if err := ctx.QueryParser(&dataFilter); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	filePath, err := handler.customerService.Export(c, dataFilter)
	helper.ErrorPanic(err)
	defer os.Remove(filePath) // Remove the file after the function exits

	fileName := filepath.Base(filePath)
	// Set headers for the Excel file
	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	// Read the Excel file and write to the response body
	data, err := os.ReadFile(filePath)
	helper.ErrorPanic(err)

	// Write data to the response body
	return ctx.Status(fiber.StatusOK).Send(data)
}

// Note 		    godoc
//
// @Summary		Import Excel customer.
// @Description	Import Excel customer.
// @Produce		application/json
// @Accept		multipart/form-data
// @Tags		customers
// @Param		file	formData	file	true	  "Import Excel customer"
// @Success		200		{object}	dto.JsonSuccess{data=string}    "Data"
// @Failure		400		{object}	dto.JsonBadRequest{}			"Validation error"
// @Failure		404		{object}	dto.JsonNotFound{}				"Data not found"
// @Failure		500		{object}	dto.JsonInternalServerError{}	"Internal server error"
// @Router		/customers/import [post]
func (handler *CustomerHandler) Import(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := new(dto.UploadCustomerRequest)

	file, err := ctx.FormFile("file")
	if err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	fileExtension := filepath.Ext(file.Filename)
	if fileExtension != ".xlsx" && fileExtension != ".xls" {
		panic(exception.NewBadRequestHandler("Invalid file type. Only .xlsx and .xls are allowed"))
	}

	request.File = file

	error := handler.customerService.Import(c, *request)
	helper.ErrorPanic(error)

	webResponse := dto.Response{
		Code:    fiber.StatusOK,
		Status:  "Ok",
		Message: "Import Successful",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
