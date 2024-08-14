package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"scylla/dto"
	"scylla/pkg/config"
)

type DmsService interface {
	GetVehicle(ctx context.Context, dataFilter dto.GeneralQueryFilter) ([]dto.VehicleResponse, dto.Meta, error)
}

type DmsServiceImpl struct{}

func NewDmsServiceImpl() DmsService {
	return &DmsServiceImpl{}
}

func (service *DmsServiceImpl) GetVehicle(ctx context.Context, dataFilter dto.GeneralQueryFilter) ([]dto.VehicleResponse, dto.Meta, error) {
	config := config.Get()

	if dataFilter.Page == 0 {
		dataFilter.Page = 1
	}

	if dataFilter.Limit == 0 {
		dataFilter.Limit = 10
	}

	endpointURL := fmt.Sprintf("%s/master/v1/vehicles?q=&page=%d&limit=%d&is_active=%d", config.Kong.Url, dataFilter.Page, dataFilter.Limit, dataFilter.IsActive)
	req, err := http.NewRequestWithContext(ctx, "GET", endpointURL, nil)
	if err != nil {
		return nil, dto.Meta{}, err
	}
	req.Header.Set("cust_id", "C220010001")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, dto.Meta{}, err
	}
	defer resp.Body.Close()

	var response struct {
		Data   []dto.VehicleResponse `json:"data"`
		Paging struct {
			TotalRecord int `json:"total_record"`
			PageCurrent int `json:"page_current"`
			PageLimit   int `json:"page_limit"`
			PageTotal   int `json:"page_total"`
		} `json:"paging"`
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, dto.Meta{}, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, dto.Meta{}, err
	}

	pagination := &dto.Meta{
		TotalData: response.Paging.TotalRecord,
		Page:      response.Paging.PageCurrent,
		Limit:     response.Paging.PageLimit,
		TotalPage: response.Paging.PageTotal,
	}

	return response.Data, *pagination, nil
}
