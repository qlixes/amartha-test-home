package services

import "github.com/qlixes/amartha/module/billing_engine/repositories"

type Billing interface {
	UserStore() error
}

type BillingService struct {
	Repository *repositories.BillingRepository
}

func (repo *BillingService) UserStore() error {

}
