package dto

type VehicleResponse struct {
	VehicleID   int64   `json:"vehicle_id"`
	VehicleNo   string  `json:"vehicle_no"`
	VehicleDesc string  `json:"vehicle_desc"`
	VehicleType string  `json:"vehicle_type_name"`
	Length      float64 `json:"length"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Volume      float64 `json:"volume"`
	DriverID    int64   `json:"driver_id"`
	HelperID    int64   `json:"helper_id"`
	DriverName  string  `json:"driver_name"`
	HelperName  string  `json:"helper_name"`
}
