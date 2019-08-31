package response

type DriverResponse struct {
	DriverID     int64  `json:"driver_id"`
	DriverName   string `json:"driver_name"`
	DriverMobile string `json:"driver_mobile"`
	VehicleID    string `json:"vehicle_id"`
	VehicleNo    string `json:"vehicle_no"`
	Status       int    `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
