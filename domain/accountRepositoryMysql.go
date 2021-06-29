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

func (db AccountRepositoryMysql) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"
	var account Account
	err := db.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &account, nil
}

func (db AccountRepositoryMysql) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := db.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := db.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryMysql(dbClient *sqlx.DB) AccountRepositoryMysql {
	return AccountRepositoryMysql{dbClient}
}
