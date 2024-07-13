package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"scylla/entity"
	"scylla/pkg/config"
	"scylla/pkg/exception"
)

type DmsService interface {
	GetVehicle(ctx context.Context, dataFilter entity.GeneralQueryFilter) ([]entity.VehicleResponse, entity.Meta, error)
}

type DmsServiceImpl struct{}

func NewDmsServiceImpl() DmsService {
	return &DmsServiceImpl{}
}

func (service *DmsServiceImpl) GetVehicle(ctx context.Context, dataFilter entity.GeneralQueryFilter) ([]entity.VehicleResponse, entity.Meta, error) {
	config, err := config.LoadConfig(".")

	if err != nil {
		exception.NewInternalServerError(err.Error())
	}

	endpointURL := fmt.Sprintf("%s/master/v1/vehicles?q=&page=%d&limit=%d&is_active=%d", config.KongUrl, dataFilter.Page, dataFilter.Limit, dataFilter.IsActive)
	req, err := http.NewRequestWithContext(ctx, "GET", endpointURL, nil)
	if err != nil {
		return nil, entity.Meta{}, err
	}
	req.Header.Set("cust_id", "C220010001")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, entity.Meta{}, err
	}
	defer resp.Body.Close()

	var response struct {
		Data   []entity.VehicleResponse `json:"data"`
		Paging struct {
			TotalRecord int `json:"total_record"`
			PageCurrent int `json:"page_current"`
			PageLimit   int `json:"page_limit"`
			PageTotal   int `json:"page_total"`
		} `json:"paging"`
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, entity.Meta{}, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, entity.Meta{}, err
	}

	pagination := &entity.Meta{
		TotalData: response.Paging.TotalRecord,
		Page:      response.Paging.PageCurrent,
		Limit:     response.Paging.PageLimit,
		TotalPage: response.Paging.PageTotal,
	}

	return response.Data, *pagination, nil
}
