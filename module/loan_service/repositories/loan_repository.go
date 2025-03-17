package repositories

import "gorm.io/gorm"

type LoanRepository struct {
	DB *gorm.DB
}
