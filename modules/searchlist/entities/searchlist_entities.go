package entities

type SearchlistUsecase interface {
	Search(req *CalculateCarbonReq) (*map[string]any, error)
}

type SearchlistRepository interface {
	SearchListDb(req *CalculateCarbonReq) (*CalculateCarbonRes, error)
}

type CalculateCarbonReq struct{
	Activity string `json:"activity_type"`
	Distance int `json:"distance_km"`
	VehicleType string `json:"vehicle_type"`
	FuelType string `json:"fuel_type"`
}

type CalculateCarbonRes struct{
	Activity string `db:"activity_group"`
	Unit string `db:"unit"`
	EmissionFactor float32 `db:"emission_factor"`
}
