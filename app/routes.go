package app

import (
	"restapi-bank-scraper/account/controller"
	"restapi-bank-scraper/scrape"

	"github.com/labstack/echo/v4"
)

func BankAccountRoute(e *echo.Echo, controller controller.AccountController) {
	group := e.Group("/api")
	group.POST("/account", controller.Save)
	group.GET("/account", controller.FindAll)
	group.GET("/account/:id", controller.FindById)
	group.PUT("/account/:id", controller.Update)
	group.DELETE("/account/:id", controller.Delete)
}

func ScrapeRoute(e *echo.Echo, controller scrape.ScrapeController) {
	e.POST("/scrape/:id", controller.GetSaldo)
}
