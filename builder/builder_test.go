package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SeatSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
