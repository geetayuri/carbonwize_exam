package controllers

import (
	// "fmt"

	"github.com/nguitarpb/carbowize-test/modules/searchlist/entities"
	"github.com/gofiber/fiber/v2"
)

type searchlistCon struct {
	SearchlistUse entities.SearchlistUsecase
}

func NewSearchlistController(r fiber.Router, searchlistUse entities.SearchlistUsecase) {
	controller := &searchlistCon{
		SearchlistUse: searchlistUse,
	}

	r.Post("/calculate", controller.Search)
}

func checkParseRequest(c *fiber.Ctx, err error) error {
	return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
		"status":      fiber.ErrBadRequest.Message,
		"status_code": fiber.ErrBadRequest.Code,
		"message":     err.Error(),
		"result":      nil,
	})
}

func parseFromDbErrorFormat(c *fiber.Ctx, err error) error {
	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
		"status":      fiber.ErrInternalServerError.Message,
		"status_code": fiber.ErrInternalServerError.Code,
		"message":     err.Error(),
		"result":      nil,
	})
}

func jsonOutput(c *fiber.Ctx, res any) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func (h *searchlistCon) Search(c *fiber.Ctx) error {
	req := new(entities.CalculateCarbonReq)
	if err := c.BodyParser(req); err != nil {
		return checkParseRequest(c, err)
	}

	res, err := h.SearchlistUse.Search(req)
	if err != nil {
		return parseFromDbErrorFormat(c, err)
	}

	return jsonOutput(c, res)
}
