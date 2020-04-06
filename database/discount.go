package database

import "database/sql"

type Discount interface {
	FindCurrentDiscount() int
}

type DiscountRepository struct {
	db *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

func (dc *DiscountRepository) FindCurrentDiscount() (discount int) {

	sql := "SELECT value FROM discount ORDER BY RAND() LIMIT 1"
	row := dc.db.QueryRow(sql)
	row.Scan(&discount)

	return discount
}
