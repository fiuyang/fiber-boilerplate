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

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// Note            godoc
//
// @Summary		Create customer
// @Description	Create customer.
// @Param			data	formData	entity.CreateCustomerRequest	true	"create customer"
// @Produce		application/json
// @Tags			customers
// @Success		201	{object}	entity.JsonCreated{data=nil}"Data"
// @Failure		400	{object}	entity.JsonBadRequest{}				"Validation error"
// @Failure		404	{object}	entity.JsonNotFound{}				"Data not found"
// @Failure		500	{object}	entity.JsonInternalServerError{}	"Internal server error"
// @Router			/customers [post]
func (handler *CustomerHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := entity.CreateCustomerRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	handler.customerService.Create(c, request)

	webResponse := entity.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Created Successful",
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

// Note            godoc
//
// @Summary		Create customer batch
// @Description	Create customer batch.
// @Param			data	body	entity.CreateCustomerBatchRequest	true	"create customer batch"
// @Produce		application/json
// @Tags			customers
// @Success		201	{object}	entity.JsonCreated{data=nil}"Data"
// @Failure		400	{object}	entity.JsonBadRequest{}				"Validation error"
// @Failure		404	{object}	entity.JsonNotFound{}				"Data not found"
// @Failure		500	{object}	entity.JsonInternalServerError{}	"Internal server error"
// @Router			/customers/batch [post]
func (handler *CustomerHandler) CreateBatch(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := entity.CreateCustomerBatchRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	handler.customerService.CreateBatch(c, request)

	webResponse := entity.Response{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: "Created Batch Successful",
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

// Note            godoc
//
// @Summary		update customer
// @Description	update customer.
// @Param			data		body	entity.UpdateCustomerRequest	true	"update customer"
// @Param			customerId	path	string							true	"customer_id"
// @Produce		application/json
// @Tags			customers
// @Success		200	{object}	entity.JsonSuccess{data=nil}		"Data"
// @Failure		400	{object}	entity.JsonBadRequest{}				"Validation error"
// @Failure		404	{object}	entity.JsonNotFound{}				"Data not found"
// @Failure		500	{object}	entity.JsonInternalServerError{}	"Internal server error"
// @Router			/customers/{customerId} [patch]
func (handler *CustomerHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := entity.UpdateCustomerRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	var params entity.CustomerParams

	if err := ctx.ParamsParser(&params); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	request.ID = params.CustomerId

	handler.customerService.Update(c, request)

	webResponse := entity.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Update Successful",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

// Note             godoc
//
//	@Summary		Delete batch customer
//	@Description	Delete batch customer.
//	@Param			data	body	entity.DeleteBatchCustomerRequest	true	"delete batch customer"
//	@Produce		application/json
//	@Tags			customers
//	@Success		200	{object}	entity.JsonSuccess{data=nil}		"Data"
//	@Failure		400	{object}	entity.JsonBadRequest{}				"Validation error"
//	@Failure		404	{object}	entity.JsonNotFound{}				"Data not found"
//	@Failure		500	{object}	entity.JsonInternalServerError{}	"Internal server error"
//	@Router			/customers/batch [delete]
func (handler *CustomerHandler) DeleteBatch(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	request := entity.DeleteBatchCustomerRequest{}
	err := ctx.BodyParser(&request)
	helper.ErrorPanic(err)

	handler.customerService.DeleteBatch(c, request)

	webResponse := entity.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Delete Batch Successful",
		Data:    nil,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

// Note 		    godoc
//
// @Summary		get customer by id.
// @Param			customerId	path	string	true	"customer_id"
// @Description	get customer by id.
// @Produce		application/json
// @Tags			customers
// @Success		200	{object}	entity.JsonSuccess{data=entity.CustomerResponse{}}	"Data"
// @Failure		400	{object}	entity.JsonBadRequest{}								"Validation error"
// @Failure		404	{object}	entity.JsonNotFound{}								"Data not found"
// @Failure		500	{object}	entity.JsonInternalServerError{}					"Internal server error"
// @Router			/customers/{customerId} [get]
func (handler *CustomerHandler) FindById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var params entity.CustomerParams

	if err := ctx.ParamsParser(&params); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	data := handler.customerService.FindById(c, params)

	webResponse := entity.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

// Note             godoc
//
// @Summary		Get all customers.
// @Description	Get all customers.
// @Produce		application/json
// @Param			limit		query	string	false	"limit"
// @Param			page		query	string	false	"page"
// @Param			username	query	string	false	"username"
// @Param			email		query	string	false	"email"
// @Param			start_date	query	string	false	"start_date"
// @Param			end_date	query	string	false	"end_date"
// @Param			sort		query	string	false	"sort"
// @Tags			customers
// @Success		200	{object}	entity.Response{data=[]entity.CustomerResponse{}}	"Data"
// @Failure		400	{object}	entity.JsonBadRequest{}								"Validation error"
// @Failure		404	{object}	entity.JsonNotFound{}								"Data not found"
// @Failure		500	{object}	entity.JsonInternalServerError{}					"Internal server error"
// @Router			/customers [get]
func (handler *CustomerHandler) FindAllPaging(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var dataFilter entity.CustomerQueryFilter

	if err := ctx.QueryParser(&dataFilter); err != nil {
		panic(exception.NewBadRequestHandler(err.Error()))
	}

	response, paging := handler.customerService.FindAllPaging(c, dataFilter)

	webResponse := entity.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
		Meta:   &paging,
	}
	utils.ResponseInterceptor(c, &webResponse)
	return ctx.Status(http.StatusOK).JSON(webResponse)
}
