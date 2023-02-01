package helper

import "restapi-bank-scraper/model"

func Account(bankAccount model.Account) model.Account {
	return model.Account{
		Id:            bankAccount.Id,
		Bank:          bankAccount.Bank,
		Username:      bankAccount.Username,
		Password:      bankAccount.Password,
		AccountNumber: bankAccount.AccountNumber,
		CheckInterval: bankAccount.CheckInterval,
		ActiveTime:    bankAccount.ActiveTime,
		AutoLogout:    bankAccount.AutoLogout,
	}
}

// to collect single bank account
func AccountResponse(bankAccount model.Account) model.AccountResponse {
	return model.AccountResponse{
		Id:            bankAccount.Id,
		Bank:          bankAccount.Bank,
		Username:      bankAccount.Username,
		AccountNumber: bankAccount.AccountNumber,
		CheckInterval: bankAccount.CheckInterval,
		ActiveTime:    bankAccount.ActiveTime,
		AutoLogout:    bankAccount.AutoLogout,
	}
}

// to collect all bank accounts
func AccountResponses(bankAccounts []model.Account) []model.AccountResponse {
	var bankAccountResponses []model.AccountResponse
	for _, bankAccount := range bankAccounts {
		bankAccountResponses = append(bankAccountResponses, AccountResponse(bankAccount))
	}
	return bankAccountResponses
}
