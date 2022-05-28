package models

import "database/sql"

type Transaction struct {
	ID           int
	UUID         int
	Name         string
	Amount       float32
	Date         int
	Type         string
	CategoryUuid sql.NullInt32
}
