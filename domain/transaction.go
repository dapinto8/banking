package domain

import (
	"github.com/dapinto8/banking/dto"
	"github.com/dapinto8/banking/errs"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

type TransactionRepository interface {
	Withdraw(Transaction) (*Transaction, *errs.AppError)
	Deposit(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t *Transaction) ToResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
