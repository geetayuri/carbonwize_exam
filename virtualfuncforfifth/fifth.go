package virtualfuncforfifth

import (

)

func CalculateEmission(activity_type *string, input_value *int, emission_factor *float32) *map[string]any{
	resUse := map[string]any{
		"activity_type": *activity_type,
		"carbon_emission": float32(*input_value) * *emission_factor,
	}

	return &resUse
}