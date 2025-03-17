package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var 

func NewBillingRepository(db *gorm.DB) *BillingRepository {
	return &BillingRepository{
		DB: db,
	}
}

func (r *BillingRepository) Register(c *gin.Context) {

}

func (r *BillingRepository) UserLoan(c *gin.Context) {

}
