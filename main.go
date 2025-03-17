package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/qlixes/amartha/config"
	"github.com/qlixes/amartha/internal/infrastructure/database"
	"github.com/qlixes/amartha/module/billing_engine/controllers"
	"github.com/qlixes/amartha/module/billing_engine/models"
	"github.com/qlixes/amartha/module/billing_engine/repositories"
	"github.com/qlixes/amartha/module/billing_engine/services"
	"github.com/qlixes/amartha/module/loan_service/controllers"
	"github.com/qlixes/amartha/module/loan_service/models"
	"github.com/qlixes/amartha/module/loan_service/repositories"
	"github.com/qlixes/amartha/module/loan_service/services"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed load config, %s \n", err.Error())
	}

	log.Println("Loaded config.")

	cfg := config.NewConfig()

	db := database.NewMysql(cfg.Db)

	err = db.AutoMigrate(
		// billing engine
		&models.User{},
		&models.Payment{},
		&models.Loan{},
		&models.LoanSchedule{},
		&models.UserBook{},

		// loan service

	)

	if err != nil {
		log.Fatalf("Failed migrate schema, %s \n", err.Error())
	}

	billingRepo := &repositories.BillingRepository{DB: db}
	billingService := &services.BillingService{Repository: billingRepo}
	billing := &controllers.BillingController{Service: billingService}

	loanRepo := &repositories.LoanRepository{DB: db}
	loanService := &service.LoanService{Repository: loanRepo}
	loan := &controllers.LoanController{service: loanService}

	r := gin.Default()

	err = r.Run(cfg.Port)

	if err != nil {
		log.Fatalf("Failed run service, %s \n", err.Error())
	}
}
