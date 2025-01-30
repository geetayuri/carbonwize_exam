package usecases

import (
	"github.com/nguitarpb/carbowize-test/modules/searchlist/entities"

	"fmt"
)

type searchlistUse struct {
	SearchlistRepo entities.SearchlistRepository
}

// Constructor
func NewSearchlistUsecase(searchlistRepo entities.SearchlistRepository) entities.SearchlistUsecase {
	return &searchlistUse{
		SearchlistRepo: searchlistRepo,
	}
}

func (u *searchlistUse) Search(req *entities.CalculateCarbonReq) (*map[string]any, error) {
	res, err := u.SearchlistRepo.SearchListDb(req)
	if err != nil {
		return nil, err
	}

	unit := fmt.Sprintf("carbon_emission_%v", res.Unit)

	resUse := map[string]any{
		"activity_type": res.Activity,
		unit: float32(req.Distance) * res.EmissionFactor,
	}

	return &resUse, nil
}
