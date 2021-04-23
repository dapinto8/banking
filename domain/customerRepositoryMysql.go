package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryMysql struct {
	client *sql.DB
}

func (db CustomerRepositoryMysql) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := db.client.Query(findAllSql)
	if err != nil {
		log.Println("Error querying customer table")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error scanning customers")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (db CustomerRepositoryMysql) ById(id string) (*Customer, error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	row := db.client.QueryRow(customerSql, id)
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		log.Println("Error scanning customer" + err.Error())
		return nil, err
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
