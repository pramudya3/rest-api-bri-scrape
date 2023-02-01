package controller

import (
	"net/http"
	"restapi-bank-scraper/account/service"
	"restapi-bank-scraper/helper"
	"restapi-bank-scraper/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountController struct {
	service service.AccountService
}

func NewAccount(service service.AccountService) *AccountController {
	return &AccountController{service}
}

func (controller *AccountController) Save(c echo.Context) error {
	var new model.CreateAccount
	err := c.Bind(&new)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.JSONResponses(http.StatusBadRequest, err.Error(), nil))
	}
	ctx := c.Request().Context()
	account, err := controller.service.Save(ctx, new)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.JSONResponses(http.StatusInternalServerError, err.Error(), nil))
	}
	return c.JSON(http.StatusCreated, helper.JSONResponses(http.StatusCreated, "success create account", account))
}

func (controller *AccountController) FindById(c echo.Context) error {
	ctx := c.Request().Context()
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	account, err := controller.service.FindById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.JSONResponses(http.StatusNotFound, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.JSONResponses(http.StatusOK, "success find account by id", account))
}

func (controller *AccountController) FindAll(c echo.Context) error {
	ctx := c.Request().Context()
	accounts, err := controller.service.FindAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.JSONResponses(http.StatusInternalServerError, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.JSONResponses(http.StatusOK, "success find all accounts", accounts))
}

func (controller *AccountController) Update(c echo.Context) error {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	update := model.UpdateAccount{}
	err := c.Bind(&update)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.JSONResponses(http.StatusBadRequest, err.Error(), nil))
	}

	update.Id = id

	ctx := c.Request().Context()
	account, err := controller.service.Update(ctx, update)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.JSONResponses(http.StatusNotFound, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.JSONResponses(http.StatusOK, "success update account", account))
}

func (controller *AccountController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := controller.service.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.JSONResponses(http.StatusNotFound, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.JSONResponses(http.StatusOK, "success delete account", nil))
}
