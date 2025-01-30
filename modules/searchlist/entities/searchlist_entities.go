package entities

type SearchlistUsecase interface {
	Search(req *CalculateCarbonReq) (*map[string]any, error)
}

type SearchlistRepository interface {
	SearchListDb(req *CalculateCarbonReq) (*CalculateCarbonRes, error)
	CheckOutRP() (res *[]Products)
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

type BeefRes struct {
	Beef map[string]int32 `json:"beef"`
}

type Products struct {
	Item string `json:"item"`
	Category string `json:"category"`
	UnitPrice int `json:"unit_price"`
	Quantity int `json:"quantity"`
}

type Output struct {
	Result map[string]int32 `json:"result"`
}
