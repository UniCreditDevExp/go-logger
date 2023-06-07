package db

import "database/sql"

type PostgresRepo struct {
	db *sql.DB
}

func (*PostgresRepo) SaveFilters() {

}

func (*PostgresRepo) LoadFilters() {

}
