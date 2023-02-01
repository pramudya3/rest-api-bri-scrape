package service

import (
	"context"
	"restapi-bank-scraper/account/repository"
	"restapi-bank-scraper/helper"
	"restapi-bank-scraper/model"

	"github.com/go-playground/validator/v10"
)

type AccountService interface {
	Save(ctx context.Context, request model.CreateAccount) (model.AccountResponse, error)
	FindById(ctx context.Context, id int) (model.Account, error)
	FindAll(ctx context.Context) ([]model.AccountResponse, error)
	Update(ctx context.Context, request model.UpdateAccount) (model.AccountResponse, error)
	Delete(ctx context.Context, id int) error
}

type AccountImpl struct {
	account repository.AccountRepository
}

func NewAccount(account repository.AccountRepository) AccountService {
	return &AccountImpl{account}
}

func (service *AccountImpl) Save(ctx context.Context, request model.CreateAccount) (model.AccountResponse, error) {
	validate := validator.New()
	err := validate.Struct(&request)

	account := model.Account{
		Bank:          request.Bank,
		Username:      request.Username,
		Password:      request.Password,
		AccountNumber: request.AccountNumber,
		CheckInterval: request.CheckInterval,
		ActiveTime:    request.ActiveTime,
		AutoLogout:    request.AutoLogout,
	}
	account = service.account.Save(ctx, account)
	return helper.AccountResponse(account), err
}

func (service *AccountImpl) FindById(ctx context.Context, id int) (model.Account, error) {
	account, err := service.account.FindById(ctx, id)
	return helper.Account(account), err
}

func (service *AccountImpl) FindAll(ctx context.Context) ([]model.AccountResponse, error) {
	accounts := service.account.FindAll(ctx)
	return helper.AccountResponses(accounts), nil
}

func (service *AccountImpl) Update(ctx context.Context, request model.UpdateAccount) (model.AccountResponse, error) {
	account, err := service.account.FindById(ctx, request.Id)

	if request.Bank != "" {
		account.Bank = request.Bank
	}
	if request.Username != "" {
		account.Username = request.Username
	}
	if request.Password != "" {
		account.Password = request.Password
	}
	if request.AccountNumber != "" {
		account.AccountNumber = request.AccountNumber
	}
	if request.CheckInterval != 0 {
		account.CheckInterval = request.CheckInterval
	}
	if request.ActiveTime != 0 {
		account.ActiveTime = request.ActiveTime
	}

	updated := service.account.Update(ctx, account)
	responseUpdate := model.AccountResponse{
		Id:            account.Id,
		Bank:          updated.Bank,
		Username:      updated.Username,
		AccountNumber: updated.AccountNumber,
		CheckInterval: updated.CheckInterval,
		ActiveTime:    updated.ActiveTime,
		AutoLogout:    updated.AutoLogout,
	}
	return responseUpdate, err
}

func (service *AccountImpl) Delete(ctx context.Context, id int) error {
	account, err := service.account.FindById(ctx, id)
	service.account.Delete(ctx, account)
	return err
}
