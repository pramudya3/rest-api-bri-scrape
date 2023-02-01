package main

import (
	"log"
	"restapi-bank-scraper/account/controller"
	"restapi-bank-scraper/account/repository"
	"restapi-bank-scraper/account/service"
	"restapi-bank-scraper/app"
	"restapi-bank-scraper/scrape"

	"github.com/labstack/echo/v4"
)

// var (
// 	check_interval int
// 	active_time    int
// 	isAutoLogout   = true
// )

func API() {
	db := app.Connect()

	// bank account
	bar := repository.NewAccount(db)
	bas := service.NewAccount(bar)
	bac := controller.NewAccount(bas)

	// scrape
	scrape := scrape.NewScrape(bar, bas, *bac)

	e := echo.New()
	app.BankAccountRoute(e, *bac)
	app.ScrapeRoute(e, *scrape)
	log.Fatal(e.Start(":1234"))
}

func main() {
	API()
}
