package models

type Driver struct {
	DriverID     int64   `json:"driver_id"`
	OrderID      int64   `json:"order_id"`
	DriverName   string  `json:"driver_name"`
	DriverMobile string  `json:"driver_mobile"`
	VehicleID    string  `json:"vehicle_id"`
	VehicleNo    string  `json:"vehicle_no"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Status       int     `json:"status"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
