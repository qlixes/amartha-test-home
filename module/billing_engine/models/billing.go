package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Delinquent Status = "Delinquent"
	Overdue    Status = "Overdue"
)

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status Status `json:"status"`
}

type Payment struct {
	gorm.Model
	UserID         uint         `json:"user_id"`
	User           User         `                        gorm:"foreignKey:UserID"`
	LoanScheduleID uint         `json:"loan_schedule_id"`
	LoanSchedules  LoanSchedule `                        gorm:"foreignKey:LoanID"`
	Amount         float32
	CreatedAt      time.Time
}

type Loan struct {
	gorm.Model
	UserID       uint `json:"user_id"`
	Users        User `json:"users"   gorm:"foreignKey:UserID"`
	Amount       float32
	LoanInterest float32
	TotalAmount  float32
}

type LoanSchedule struct {
	gorm.Model
	UserID     *uint `json:"user_id"`
	Users      *User `               gorm:"foreignKey:UserID"`
	LoanID     *uint `json:"loan_id"`
	Loans      *Loan `               gorm:"foreignKey:LoanID"`
	ScheduleAt time.Time
	ExpiredAt  time.Time
	IsExpired  bool
}

type UserBook struct {
	gorm.Model
	UserID    *uint `json:"user_id"`
	Users     *User `               gorm:"foreignKey:UserID"`
	Debet     float32
	Credit    float32
	Balance   float32
	CreatedAt time.Time
}
