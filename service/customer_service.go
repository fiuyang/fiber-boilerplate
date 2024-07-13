package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"math"
	"scylla/entity"
	"scylla/model"
	"scylla/pkg/exception"
	"scylla/pkg/helper"
	"scylla/repository"
)

type CustomerService interface {
	Create(ctx context.Context, request entity.CreateCustomerRequest)
	CreateBatch(ctx context.Context, request entity.CreateCustomerBatchRequest)
	Update(ctx context.Context, request entity.UpdateCustomerRequest)
	DeleteBatch(ctx context.Context, request entity.DeleteBatchCustomerRequest)
	FindById(ctx context.Context, request entity.CustomerParams) (response entity.CustomerResponse)
	FindAll(ctx context.Context) (response []entity.CustomerResponse)
	FindAllPaging(ctx context.Context, dataFilter entity.CustomerQueryFilter) (response []entity.CustomerResponse, paging entity.Meta)
}

type CustomerServiceImpl struct {
	customerRepo repository.CustomerRepo
	validate     *validator.Validate
}

func NewCustomerServiceImpl(customerRepo repository.CustomerRepo, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		customerRepo: customerRepo,
		validate:     validate,
	}
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request entity.CreateCustomerRequest) {
	err := service.validate.Struct(request)
	helper.ErrorPanic(err)

	dataset := model.Customer{
		Username: request.Username,
		Email:    request.Email,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	service.customerRepo.Insert(ctx, dataset)
}

func (service *CustomerServiceImpl) CreateBatch(ctx context.Context, request entity.CreateCustomerBatchRequest) {
	err := service.validate.Struct(request)
	helper.ErrorPanic(err)

	var customers []model.Customer
	for _, req := range request.Customers {
		customer := model.Customer{
			Username: req.Username,
			Email:    req.Email,
			Phone:    req.Phone,
			Address:  req.Address,
		}
		customers = append(customers, customer)
	}

	batchSize := len(request.Customers)

	err = service.customerRepo.InsertBatch(ctx, customers, batchSize)
	if err != nil {
		panic(exception.NewInternalServerError(err.Error()))
	}
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request entity.UpdateCustomerRequest) {
	err := service.validate.Struct(request)
	helper.ErrorPanic(err)

	dataset, err := service.customerRepo.FindById(ctx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	dataset.Username = request.Username
	dataset.Email = request.Email
	dataset.Phone = request.Phone
	dataset.Address = request.Address

	service.customerRepo.Update(ctx, dataset)
}

func (service *CustomerServiceImpl) DeleteBatch(ctx context.Context, request entity.DeleteBatchCustomerRequest) {
	err := service.validate.Struct(request)
	helper.ErrorPanic(err)

	service.customerRepo.DeleteBatch(ctx, request.ID)
}

func (service *CustomerServiceImpl) FindById(ctx context.Context, request entity.CustomerParams) (response entity.CustomerResponse) {
	result, err := service.customerRepo.FindById(ctx, request.CustomerId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	helper.Automapper(result, &response)
	return response
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) (response []entity.CustomerResponse) {
	result, err := service.customerRepo.FindAll(ctx)

	if err != nil {
		panic(exception.NewInternalServerError(err.Error()))
	}

	for _, row := range result {
		var res entity.CustomerResponse
		helper.Automapper(row, &res)
		response = append(response, res)
	}
	return response
}

func (service *CustomerServiceImpl) FindAllPaging(ctx context.Context, dataFilter entity.CustomerQueryFilter) (response []entity.CustomerResponse, paging entity.Meta) {

	result := service.customerRepo.FindAllPaging(ctx, dataFilter)

	for _, value := range result {
		var res entity.CustomerResponse
		helper.Automapper(value, &res)

		response = append(response, res)
	}

	if dataFilter.Limit == 0 {
		dataFilter.Limit = 10
	}

	if dataFilter.Page == 0 {
		dataFilter.Page = 1
	}

	var total int
	if len(result) > 0 {
		total = len(result)
	} else {
		total = 0
	}

	paging.Page = dataFilter.Page
	paging.Limit = dataFilter.Limit
	paging.TotalData = total
	paging.TotalPage = int(math.Ceil(float64(total) / float64(dataFilter.Limit)))

	return response, paging
}
