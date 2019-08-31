package services

import (
	"github.com/freeddser/dmgin/repository"
	"github.com/freeddser/dmgin/response"
)

func GetDriverByID(ID int64) (response.DriverResponse, error) {
	if ID == 0 {
		return response.DriverResponse{}, nil
	}
	repo := repository.GetDriverRepository()
	data, err := repo.GetDriverByID(ID)
	if err != nil {
		return response.DriverResponse{}, err
	}
	if data.DriverID == 0 {
		return response.DriverResponse{}, err
	}

	//driver := make([]response.DriverResponse, 0)
	//
	//d
	return response.DriverResponse{
		DriverID:     data.DriverID,
		DriverName:   data.DriverName,
		DriverMobile: data.DriverMobile,
		VehicleID:    data.VehicleID,
		VehicleNo:    data.VehicleNo,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}, nil
}
