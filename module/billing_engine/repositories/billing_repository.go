package repositories

import (
	"github.com/qlixes/amartha/module/billing_engine/models"
	"gorm.io/gorm"
)

type BillingRepository struct {
	DB *gorm.DB
}

func (r *BillingRepository) Store(user *models.User) error {
	err = r.DB.Create(&user)

	return err
}
