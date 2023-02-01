package repository

import (
	"context"
	"database/sql"
	"errors"
	"restapi-bank-scraper/helper"
	"restapi-bank-scraper/model"
)

type AccountRepository interface {
	Save(ctx context.Context, account model.Account) model.Account
	FindById(ctx context.Context, id int) (model.Account, error)
	FindAll(ctx context.Context) []model.Account
	Update(ctx context.Context, update model.Account) model.Account
	Delete(ctx context.Context, account model.Account)
}

type AccountImpl struct {
	db *sql.DB
}

func NewAccount(db *sql.DB) AccountRepository {
	return &AccountImpl{db}
}

func (repository *AccountImpl) Save(ctx context.Context, account model.Account) model.Account {
	query := "insert into accounts (bank, username, password, account_number, check_interval, time_active, auto_logout) values(?,?,?,?,?,?,?)"
	result, err := repository.db.ExecContext(ctx, query, account.Bank, account.Username, account.Password, account.AccountNumber, account.CheckInterval, account.ActiveTime, account.AutoLogout)
	helper.PanicIfError(err)
	id, _ := result.LastInsertId()
	account.Id = int(id)
	return account
}

func (repository *AccountImpl) FindById(ctx context.Context, id int) (model.Account, error) {
	query := "select id, bank, username, password, account_number, check_interval, time_active, auto_logout from accounts where id = ?"
	rows, err := repository.db.QueryContext(ctx, query, id)

	helper.PanicIfError(err)
	defer rows.Close()

	ba := model.Account{}
	if rows.Next() {
		err := rows.Scan(&ba.Id, &ba.Bank, &ba.Username, &ba.Password, &ba.AccountNumber, &ba.CheckInterval, &ba.ActiveTime, &ba.AutoLogout)
		helper.PanicIfError(err)
		return ba, nil
	} else {
		return ba, errors.New("id not found")
	}
}

func (repository *AccountImpl) FindAll(ctx context.Context) []model.Account {
	query := "select id, bank, username, account_number, check_interval, time_active, auto_logout from accounts"

	rows, err := repository.db.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		account := model.Account{}
		err := rows.Scan(&account.Id, &account.Bank, &account.Username, &account.AccountNumber, &account.CheckInterval, &account.ActiveTime, &account.AutoLogout)
		helper.PanicIfError(err)
		accounts = append(accounts, account)
	}
	return accounts
}

func (repository *AccountImpl) Update(ctx context.Context, update model.Account) model.Account {
	query := "update accounts set bank = ?, username = ?, password = ?, account_number = ?, check_interval = ?, time_active = ?, auto_logout = ? where id = ?"
	result, err := repository.db.ExecContext(ctx, query, update.Bank, update.Username, update.Password, update.AccountNumber, update.CheckInterval, update.ActiveTime, update.AutoLogout, update.Id)
	helper.PanicIfError(err)

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return model.Account{}
	}
	return update
}

func (repository *AccountImpl) Delete(ctx context.Context, account model.Account) {
	query := "delete from accounts where id = ?"
	_, err := repository.db.ExecContext(ctx, query, account.Id)
	helper.PanicIfError(err)
}
