package domain

import (
	"database/sql"

	"github.com/dapinto8/banking/errs"
	"github.com/dapinto8/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryMysql struct {
	client *sqlx.DB
}

func (db CustomerRepositoryMysql) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = db.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = db.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error querying customers" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (db CustomerRepositoryMysql) ById(id string) (*Customer, *errs.AppError) {

	var customer Customer
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	err := db.client.Get(&customer, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error scanning customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &customer, nil
}

func NewCustomerRepositoryMysql(dbClient *sqlx.DB) CustomerRepositoryMysql {

	return CustomerRepositoryMysql{dbClient}
}
