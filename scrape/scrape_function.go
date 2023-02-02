package scrape

import (
	"fmt"
	"net/http"
	"restapi-bank-scraper/account/service"
	"restapi-bank-scraper/helper"
	"restapi-bank-scraper/model"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/labstack/echo/v4"
)

type ScrapeController struct {
	service service.AccountService
}

func NewScrape(service service.AccountService) *ScrapeController {
	return &ScrapeController{service}
}

func (s *ScrapeController) GetSaldo(c echo.Context) error {
	// get account by id
	ctx := c.Request().Context()
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	account, err := s.service.FindById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.JSONResponses(http.StatusNotFound, err.Error(), nil))
	}

	// browser edge
	u := launcher.New().Bin("/usr/bin/microsoft-edge").MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://ib.bri.co.id/ib-bri")

	// browser google chrome
	// browser := rod.New().MustConnect()
	// page := browser.MustPage("https://ib.bri.co.id/ib-bri").MustWaitLoad()
	fmt.Println("into website https://bri.co.id/ib-bri")

	// Get captcha image
	captcha, _ := page.MustElement("#simple_img > img").MustWaitVisible().Screenshot(proto.PageCaptureScreenshotFormatPng, 1500)
	text := helper.Captcha2Text(captcha)
	fmt.Println(text)

	// Login
	if len(text) > 4 {
		page.MustElement("#loginForm > div.validation > input[type=text]").MustInput(text[1:5])
	}
	err = page.MustElement("#loginForm > div.validation > input[type=text]").MustInput(text).WaitVisible()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.JSONResponses(http.StatusBadRequest, err.Error(), nil))
	}
	page.MustElement("#loginForm > input[type=text]:nth-child(5)").MustInput(account.Username).WaitVisible()
	page.MustElement("#loginForm > input[type=password]:nth-child(8)").MustInput(account.Password).WaitVisible()
	err = page.MustElement("#loginForm > button").MustClick().WaitInvisible()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, helper.JSONResponses(http.StatusNotAcceptable, err.Error(), nil))
	}
	fmt.Println("success login into bri website")

	page.MustElement("#myaccounts").MustClick().WaitVisible()

	// Get Saldo Tabungan
	fmt.Println("start scraping....")
	fr1 := page.MustElement("#iframemenu").MustFrame()
	fr1.MustElement("body > div.submenu.active > div:nth-child(2) > a").MustClick().MustWaitVisible()
	time.Sleep(3 * time.Second)
	page.MustScreenshot("total-saldo.png")

	// Get tabel saldo tabungan
	fr2 := page.MustElement("#content").MustFrame()
	noRek := fr2.MustElement("#Any_0 > td:nth-child(1)").MustText()
	jenisProduk := fr2.MustElement("#Any_0 > td:nth-child(2)").MustText()
	nama := fr2.MustElement("#Any_0 > td:nth-child(3)").MustText()
	mataUang := fr2.MustElement("#Any_0 > td:nth-child(4)").MustText()
	saldo := fr2.MustElement("#Any_0 > td:nth-child(5)").MustText()
	fmt.Println("success scraping")

	response := model.GetSaldo{
		NomorRekening: noRek,
		JenisProduk:   jenisProduk,
		Nama:          nama,
		MataUang:      mataUang,
		Saldo:         saldo,
	}
	c.JSON(http.StatusOK, helper.JSONResponses(http.StatusOK, "success get saldo", response))

	fmt.Printf("\nNomor Rekening : %s\n\nJenis Produk : %s\n\nNama : %s\n\nMata Uang : %s\n\nSaldo : %s\n\n", noRek, jenisProduk, nama, mataUang, saldo)

	time.Sleep(time.Minute)

	// logout
	page.MustElement("#main-page > div.headerwrap > div > div.uppernav.col-1-2 > span:nth-child(1) > a:nth-child(4)").MustClick()

	return err
}
