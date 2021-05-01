package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/dapinto8/banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryMysql struct {
	client *sql.DB
}

func (db CustomerRepositoryMysql) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
	var findAllSql string
	if status == "" {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = db.client.Query(findAllSql)
	} else {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = db.client.Query(findAllSql, status)
	}

	if err != nil {
		return nil, errs.NewUnexpectedError("Error querying customer table")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error scanning customers")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (db CustomerRepositoryMysql) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	row := db.client.QueryRow(customerSql, id)
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error scanning customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryMysql() CustomerRepositoryMysql {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryMysql{client}
}
