package domain

import (
	"strconv"

	"github.com/dapinto8/banking/errs"
	"github.com/dapinto8/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryMysql struct {
	client *sqlx.DB
}

func (db AccountRepositoryMysql) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := db.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting last insert id for new account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexp ected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryMysql(dbClient *sqlx.DB) AccountRepositoryMysql {
	return AccountRepositoryMysql{dbClient}
}
