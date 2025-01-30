package repositories

import (
	// "fmt"
	"context"
	"time"
	"errors"

	"github.com/nguitarpb/carbowize-test/modules/searchlist/entities"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SearchlistRepo struct {
	Db *pgxpool.Pool
}

func NewSearchlistRepository(db *pgxpool.Pool) entities.SearchlistRepository {
	return &SearchlistRepo{
		Db: db,
	}
}

func (r *SearchlistRepo) SearchListDb(req *entities.CalculateCarbonReq) (*entities.CalculateCarbonRes, error) {
	query := `
		SELECT 
		activity_group, 
		unit, 
		emission_factor
	FROM 
		"emission_factors"
	WHERE 
		activity_group = $1
		AND type = $2
		AND name = $3;
	`

	res := entities.CalculateCarbonRes{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	row := r.Db.QueryRow(ctx, query, req.Activity, req.VehicleType, req.FuelType)
	err := row.Scan(&res.Activity, &res.Unit, &res.EmissionFactor)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("error, Emission Factor not found")
		}
		return nil, err
	}
	return &res, nil
}

func (r *SearchlistRepo) CheckOutRP() (*[]entities.Products) {
	res := []entities.Products{
		{
			Item : "Nike Air Zoom",
			Category: "sport",
			UnitPrice: 3000,
			Quantity: 1,
		},
		{
			Item : "Protein Bar Herbalife Deluxe",
			Category: "sport",
			UnitPrice: 50,
			Quantity: 10,
		},
		{
			Item : "Starbucks Coffee",
			Category: "beverage",
			UnitPrice: 500,
			Quantity: 3,
		},
		{
			Item : "iPhone 13",
			Category: "electronic",
			UnitPrice: 40000,
			Quantity: 1,
		},
		{
			Item : "Cleaning alcohol",
			Category: "household",
			UnitPrice: 5000,
			Quantity: 1,
		},
	}

	return &res
}