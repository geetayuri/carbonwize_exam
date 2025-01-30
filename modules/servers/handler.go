package servers

import (
	_searchlistHttp "github.com/nguitarpb/carbowize-test/modules/searchlist/controllers"
	_searchlistRepository "github.com/nguitarpb/carbowize-test/modules/searchlist/repositories"
	_searchlistUsecase "github.com/nguitarpb/carbowize-test/modules/searchlist/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	searchlistGroup := s.App.Group("/api")
	carbonGroup := searchlistGroup.Group("/carbon")
	footprintGroup := carbonGroup.Group("/footprint")
	
	searchlistRepository := _searchlistRepository.NewSearchlistRepository(s.PostgresDB)
	searchlistUsecase := _searchlistUsecase.NewSearchlistUsecase(searchlistRepository)
	_searchlistHttp.NewSearchlistController(footprintGroup, searchlistUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
