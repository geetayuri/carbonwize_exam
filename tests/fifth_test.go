package tests

import (
	"testing"

	"github.com/nguitarpb/carbowize-test/virtualfuncforfifth"
	"github.com/stretchr/testify/assert"
)

func TestCalculateEmission(t *testing.T) {
	activity := "driving"
	input := 100
	factor := float32(2.5)

	expectedEmission := float32(input) * factor

	result := virtualfuncforfifth.CalculateEmission(&activity, &input, &factor)

	assert.NotNil(t, result)

	assert.Equal(t, activity, (*result)["activity_type"].(string))

	assert.Equal(t, expectedEmission, (*result)["carbon_emission"])
}

